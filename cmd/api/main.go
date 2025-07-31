package main

import (
	"fmt"
	"os"

	"github.com/joshuaalpuerto/go-rest-api/config"
	validator "github.com/joshuaalpuerto/go-rest-api/internal/common/validator"
	companyusecases "github.com/joshuaalpuerto/go-rest-api/internal/company/usecases"
	onboardingusecases "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/usecases"
	userusecases "github.com/joshuaalpuerto/go-rest-api/internal/user/usecases"

	"github.com/joshuaalpuerto/go-rest-api/internal/infra/db"
	infrarepositories "github.com/joshuaalpuerto/go-rest-api/internal/infra/repositories"
)

// our DI container
type application struct {
	conf         config.Conf
	repositories repositories
	validator    *validator.Validator
}

type repositories struct {
	companyRepository    companyusecases.CompanyRepository
	userRepository       userusecases.UserRepository
	onboardingRepository onboardingusecases.OnboardingRepository
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

	validator := validator.NewValidator()

	companyRepository := infrarepositories.NewCompanyRepository(*db)
	userRepository := infrarepositories.NewUserRepository(*db)
	onboardingRepository := infrarepositories.NewOnboardingRepository(*db)
	repositories := repositories{
		companyRepository:    companyRepository,
		userRepository:       userRepository,
		onboardingRepository: onboardingRepository,
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
