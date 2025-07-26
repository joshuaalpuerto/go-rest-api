package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joshuaalpuerto/go-rest-api/cmd/api/middlewares"
	"github.com/joshuaalpuerto/go-rest-api/config"

	companycontroller "github.com/joshuaalpuerto/go-rest-api/internals/api/company/controller"
	companyusecase "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecase"
)

// our DI container
type application struct {
	conf         config.Conf
	repositories struct {
		companyRepository companyusecase.CompanyRepository
	}
	// later we will have storer here and repository
}

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

	companyHandler := companycontroller.NewCompanyHandler(app.repositories.companyRepository)

	mux.HandleFunc(fmt.Sprintf("/%s/companies", version), WrapHandlerWithMiddlewares(companyHandler.GetAllCompanies, middlewares.RequestLogger(), middlewares.CORS()))
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.GetCompanyByID)
	// mux.HandleFunc(fmt.Sprintf("/%s/companies", version), companyHandler.CreateCompany)
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.UpdateCompany)
	// mux.HandleFunc(fmt.Sprintf("/%s/companies/{id}", version), companyHandler.DeleteCompany)

	// TODO: add here other endpoint

	return mux
}

// HandlerFunc represents a handler function that returns data and error
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (any, error)

// WrapHandler wraps a handler function with standard middleware and response handling
func WrapHandlerWithMiddlewares(handler HandlerFunc, mw ...middlewares.MiddlewareFunc) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(w, r)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		SuccessResponse(w, data)
	}

	// check if there are middlewares if yes then chain them, otherwise just call handler
	if len(mw) > 0 {
		return middlewares.Chain(h, mw...)
	}

	return h
}
