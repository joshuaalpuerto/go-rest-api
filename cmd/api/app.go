package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpmiddlewares "github.com/joshuaalpuerto/go-rest-api/internal/common/http/middlewares"
	companyhttp "github.com/joshuaalpuerto/go-rest-api/internal/company/interfaces/http"
	onboardinghttp "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/interfaces/http"
	userhttp "github.com/joshuaalpuerto/go-rest-api/internal/user/interfaces/http"
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

	appMiddlewares := []httpmiddlewares.MiddlewareFunc{
		httpmiddlewares.RequestLogger(),
		httpmiddlewares.CORS(),
	}
	companyHandler := companyhttp.NewCompanyHandler(app.repositories.companyRepository, *app.validator)
	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodGet, version), httpmiddlewares.Chain(companyHandler.GetAllCompanies, appMiddlewares...))
	mux.HandleFunc(fmt.Sprintf("%s /%s/companies/{id}", http.MethodGet, version), httpmiddlewares.Chain(companyHandler.GetCompanyByID, appMiddlewares...))

	mux.HandleFunc(fmt.Sprintf("%s /%s/companies", http.MethodPost, version), httpmiddlewares.Chain(
		companyHandler.CreateCompany,
		appMiddlewares...,
	))

	userHandler := userhttp.NewUserHandler(app.repositories.userRepository, *app.validator)
	mux.HandleFunc(fmt.Sprintf("%s /%s/users", http.MethodPost, version), httpmiddlewares.Chain(
		userHandler.CreateUser,
		appMiddlewares...,
	))

	onboardingHandler := onboardinghttp.NewOnboardingHandler(app.repositories.onboardingRepository, app.repositories.companyRepository, app.repositories.userRepository, *app.validator)
	mux.HandleFunc(fmt.Sprintf("%s /%s/onboarding", http.MethodPost, version), httpmiddlewares.Chain(
		onboardingHandler.RegisterUserCompany,
		appMiddlewares...,
	))

	return mux
}
