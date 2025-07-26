package main

import (
	"fmt"
	"os"

	"github.com/joshuaalpuerto/go-rest-api/config"
	companyusecase "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecase"
	"github.com/joshuaalpuerto/go-rest-api/internals/infra/db"
	infrarepositories "github.com/joshuaalpuerto/go-rest-api/internals/infra/repositories"
)

// our DI container
type application struct {
	conf         config.Conf
	repositories repositories
}

type repositories struct {
	companyRepository companyusecase.CompanyRepository
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

	companyRepository := infrarepositories.NewCompanyRepository(*db)

	repositories := repositories{
		companyRepository: companyRepository,
	}

	app := &application{
		conf:         conf,
		repositories: repositories,
	}

	if err := app.Start(app.Routes()); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
