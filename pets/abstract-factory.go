package pets

import (
	"errors"
	"fmt"
	"log"

	"github.com/abhilashdk2016/golang-design-patterns/configuration"
	"github.com/abhilashdk2016/golang-design-patterns/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (df *DogAbstractFactory) newPetWithBreed(breed string) AnimalInterface {
	app := configuration.GetInstance()
	b, _ := app.Models.DogBreed.GetBreedByName(breed)
	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *b,
		},
	}
}

type CatAbstractFactory struct{}

func (cf *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *CatAbstractFactory) newPetWithBreed(breed string) AnimalInterface {
	app := configuration.GetInstance()
	b, err := app.CatSevice.Remote.GetBreedByName(breed)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &CatFromFactory{
		Pet: &models.Cat{
			Breed: *b,
		},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("unknown species")
	}
}

func NewPetWithBreedFromAbstractfactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("invalid species")
	}
}
