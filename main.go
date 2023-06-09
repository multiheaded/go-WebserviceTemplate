package main

import (
	"github.com/multiheaded/go-WebserviceTemplate/pkg/app"
)

// Entry point for the application
func main() {
	// create a new app instance
	wsapp, err := app.NewInstance()

	if err != nil {
		panic(err)
	}

	wsapp.Run()
}
