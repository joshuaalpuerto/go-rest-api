package main

import (
	"fmt"
	"os"

	"github.com/joshuaalpuerto/go-rest-api/config"
	companyusecase "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecase"
	"github.com/joshuaalpuerto/go-rest-api/internals/infra/db"
	infrarepositories "github.com/joshuaalpuerto/go-rest-api/internals/infra/repositories"
)

// Bootstrap of the application
func main() {
	conf := config.New()
	db, err := db.NewDatabase(conf.DB)
	if err != nil {
		fmt.Println("Error creating database:", err)
		os.Exit(1)
	}

	defer db.Close()

	companyRepository := infrarepositories.NewCompanyRepository(db)

	app := &application{
		conf: conf,
		repositories: struct {
			companyRepository companyusecase.CompanyRepository
		}{
			companyRepository: companyRepository,
		},
	}

	if err := app.Start(app.Routes()); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
