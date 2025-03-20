package configuration

import (
	"database/sql"
	"sync"

	"github.com/abhilashdk2016/golang-design-patterns/adapters"
	"github.com/abhilashdk2016/golang-design-patterns/models"
)

type Application struct {
	Models    *models.Models
	CatSevice *adapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var catService *adapters.RemoteService

func New(pool *sql.DB, cs *adapters.RemoteService) *Application {
	db = pool
	catService = cs
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{Models: models.New(db), CatSevice: catService}
	})
	return instance
}
