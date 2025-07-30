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

	companyController := controllers.NewCompanyHandler(app.repositories.companyRepository, app.validator)
	appMiddlewares := []middlewares.MiddlewareFunc{
		middlewares.RequestLogger(),
		middlewares.CORS(),
	}

	// TODO: move global level middleware. ( remove cors and request logger here.)
	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodGet, version), WrapHandlerWithMiddlewares(companyController.GetAllCompanies, appMiddlewares...))
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.GetCompanyByID)

	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodPost, version), WrapHandlerWithMiddlewares(
		companyController.CreateCompany,
		appMiddlewares...,
	))
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.UpdateCompany)
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.DeleteCompany)

	// TODO: add here other endpoint

	return mux
}

// WrapHandler wraps a handler function with standard middleware and response handling
func WrapHandlerWithMiddlewares(handler http.HandlerFunc, mw ...middlewares.MiddlewareFunc) http.HandlerFunc {

	// check if there are middlewares if yes then chain them, otherwise just call handler
	if len(mw) > 0 {
		return middlewares.Chain(handler, mw...)
	}

	return handler
}
