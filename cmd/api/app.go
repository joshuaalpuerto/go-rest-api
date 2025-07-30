package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joshuaalpuerto/go-rest-api/cmd/api/middlewares"

	"github.com/joshuaalpuerto/go-rest-api/cmd/api/controllers"
)

func (app *application) Start(mux http.Handler) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", app.conf.Server.Host, app.conf.Server.Port),
		Handler: mux,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), app.conf.Server.TimeoutIdle)
		defer cancel()

		fmt.Println("shutting down server")
		shutdown <- srv.Shutdown(ctx)
	}()

	fmt.Println("server is starting on port", app.conf.Server.Port)

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()
	version := app.conf.Version

	appMiddlewares := []middlewares.MiddlewareFunc{
		middlewares.RequestLogger(),
		middlewares.CORS(),
	}
	companyController := controllers.NewCompanyController(app.repositories.companyRepository, app.validator)
	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodGet, version), middlewares.Chain(companyController.GetAllCompanies, appMiddlewares...))
	mux.HandleFunc(fmt.Sprintf("%s /%s/companies/{id}", http.MethodGet, version), middlewares.Chain(companyController.GetCompanyByID, appMiddlewares...))

	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodPost, version), middlewares.Chain(
		companyController.CreateCompany,
		appMiddlewares...,
	))

	userController := controllers.NewUserController(app.repositories.userRepository, app.validator)
	mux.HandleFunc(fmt.Sprintf("%s /%s/users", http.MethodPost, version), middlewares.Chain(
		userController.CreateUser,
		appMiddlewares...,
	))

	return mux
}
