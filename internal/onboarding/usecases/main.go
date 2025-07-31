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
	companyRepository    companyusecases.CompanyRepository
	userRepository       userusecases.UserRepository
}

func NewOnboardingUsecase(onboardingRepository OnboardingRepository, companyRepository companyusecases.CompanyRepository, userRepository userusecases.UserRepository) OnboardingUsecase {
	return OnboardingUsecase{
		onboardingRepository: onboardingRepository,
		companyRepository:    companyRepository,
		userRepository:       userRepository,
	}
}
