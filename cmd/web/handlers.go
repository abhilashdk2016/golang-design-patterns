package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/abhilashdk2016/golang-design-patterns/models"
	"github.com/abhilashdk2016/golang-design-patterns/pets"
	"github.com/abhilashdk2016/toolkit/v2"
	"github.com/go-chi/chi/v5"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CeateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed dog.").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CeateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("cat breed").
		SetWeight(15).
		SetDescription("A nice cat.").
		SetGeographicOrigin("London").
		SetColor("Black and White").
		SetAge(3).
		Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreedsWithAdapter(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools

	carBreeds, err := app.App.CatSevice.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, carBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	species := chi.URLParam(r, "species")
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)
	pet, err := pets.NewPetWithBreedFromAbstractfactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, pet)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	dom, _ := app.App.Models.Dog.GetDogOfMonthById(1)
	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-12-02")
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedID:          breed.ID,
			Color:            "Black & Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: 0,
			Description:      "Sam is a very good boy",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}
