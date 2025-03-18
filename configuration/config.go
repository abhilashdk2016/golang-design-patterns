package configuration

import (
	"database/sql"
	"sync"

	"github.com/abhilashdk2016/golang-design-patterns/models"
)

type Application struct {
	Models *models.Models
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{Models: models.New(db)}
	})
	return instance
}
