package main

import (
	"multiheaded/webservice_template/pkg/app"
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
