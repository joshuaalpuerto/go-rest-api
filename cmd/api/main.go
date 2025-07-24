package main

import (
	"fmt"
	"os"

	"github.com/joshuaalpuerto/go-rest-api/config"
)

// Bootstrap of the application
func main() {
	conf := config.New()

	// TODO: setup DB connection here.

	app := &application{
		conf: conf,
	}

	mux := app.CreateRoutes()

	err := app.CreateServer(mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
