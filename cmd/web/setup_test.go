package main

import (
	"os"
	"testing"

	"github.com/abhilashdk2016/golang-design-patterns/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
