package onboardingusecases

import (
	"context"

	companyusecases "github.com/joshuaalpuerto/go-rest-api/internal/company/usecases"
	onboardingdomain "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/domain"
	userusecases "github.com/joshuaalpuerto/go-rest-api/internal/user/usecases"
)

type OnboardingRepository interface {
	Create(ctx context.Context, newUserCompany onboardingdomain.UserCompany) (*onboardingdomain.UserCompanyDB, error)
}

type OnboardingUsecase struct {
	onboardingRepository OnboardingRepository
	companyService       companyusecases.CompanyUsecase
	userService          userusecases.UserUsecase
}

func NewOnboardingUsecase(onboardingRepository OnboardingRepository, companyService companyusecases.CompanyUsecase, userService userusecases.UserUsecase) OnboardingUsecase {
	return OnboardingUsecase{
		onboardingRepository: onboardingRepository,
		companyService:       companyService,
		userService:          userService,
	}
}
