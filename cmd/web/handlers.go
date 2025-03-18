package main

import (
	"fmt"
	"net/http"

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
