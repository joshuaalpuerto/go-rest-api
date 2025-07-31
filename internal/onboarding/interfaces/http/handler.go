package onboardinghttp

import (
	"encoding/json"
	"net/http"

	response "github.com/joshuaalpuerto/go-rest-api/internal/common/http/response"
	commonvalidator "github.com/joshuaalpuerto/go-rest-api/internal/common/validator"
	companyusecases "github.com/joshuaalpuerto/go-rest-api/internal/company/usecases"
	onboardingdomain "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/domain"
	onboardingusecases "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/usecases"
	userusecases "github.com/joshuaalpuerto/go-rest-api/internal/user/usecases"
)

type OnboardingHandler struct {
	onboardingService onboardingusecases.OnboardingUsecase
	validator         commonvalidator.Validator
}

func NewOnboardingHandler(onboardingRepository onboardingusecases.OnboardingRepository, companyRepository companyusecases.CompanyRepository, userRepository userusecases.UserRepository, validator commonvalidator.Validator) OnboardingHandler {
	return OnboardingHandler{
		onboardingService: onboardingusecases.NewOnboardingUsecase(onboardingRepository, companyRepository, userRepository),
		validator:         validator,
	}
}

func (h OnboardingHandler) RegisterUserCompany(w http.ResponseWriter, r *http.Request) {
	var newUserCompany NewUserCompany
	response := response.Response{}

	if err := json.NewDecoder(r.Body).Decode(&newUserCompany); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := h.validator.Validate(&newUserCompany); errs != nil {
		response.SendErrorResponse(w, errs.Error(), http.StatusBadRequest)
		return
	}

	createdCompany, err := h.onboardingService.RegisterUserCompany(onboardingdomain.NewUserCompany{
		CompanyName:  newUserCompany.CompanyName,
		UserName:     newUserCompany.UserName,
		UserEmail:    newUserCompany.UserEmail,
		UserPassword: newUserCompany.UserPassword,
	})
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, createdCompany, http.StatusCreated)
}
