package main

import (
	"fmt"
	"os"

	"github.com/joshuaalpuerto/go-rest-api/config"
	companyusecases "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecases"

	"github.com/joshuaalpuerto/go-rest-api/internals/infra/db"
	infrarepositories "github.com/joshuaalpuerto/go-rest-api/internals/infra/repositories"
	infravalidator "github.com/joshuaalpuerto/go-rest-api/internals/infra/validator"
)

// our DI container
type application struct {
	conf         config.Conf
	repositories repositories
	validator    *infravalidator.Validator
}

type repositories struct {
	companyRepository companyusecases.CompanyRepository
}

// Bootstrap of the application
func main() {
	conf := config.New()
	db, err := db.NewDatabase(conf.DB)
	if err != nil {
		fmt.Println("Error creating database:", err)
		os.Exit(1)
	}

	defer db.Close()

	validator := infravalidator.NewValidator()
	companyRepository := infrarepositories.NewCompanyRepository(*db)

	repositories := repositories{
		companyRepository: companyRepository,
	}

	app := &application{
		conf:         conf,
		repositories: repositories,
		validator:    validator,
	}

	if err := app.Start(app.Routes()); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
