package main

import (
	"os"
	"testing"

	"github.com/abhilashdk2016/golang-design-patterns/configuration"
	"github.com/abhilashdk2016/golang-design-patterns/models"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}
	testApp = application{
		App:       configuration.New(nil),
		catSevice: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{ID: 1, Breed: "Tomcat", Details: "Some Details"},
	}

	return breeds, nil
}
