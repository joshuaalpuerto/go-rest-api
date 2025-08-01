package onboardingusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/company/domain"
	onboardingdomain "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/domain"
	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/user/domain"
)

func (c *OnboardingUsecase) RegisterUserCompany(nwu onboardingdomain.NewUserCompany) (*onboardingdomain.UserCompany, error) {

	createdCompany, companyRepoErr := c.companyService.CreateCompany(companydomain.NewCompany{
		Name: nwu.CompanyName,
	})
	if companyRepoErr != nil {
		return nil, fmt.Errorf("failed to create company: %w", companyRepoErr)
	}

	createdUser, userRepoErr := c.userService.CreateUser(userdomain.NewUser{
		Name:     nwu.UserName,
		Email:    nwu.UserEmail,
		Password: nwu.UserPassword,
	})
	if userRepoErr != nil {
		return nil, fmt.Errorf("failed to create user: %w", userRepoErr)
	}

	// createdCompany and createdUser are already domain entities
	company := createdCompany
	user := createdUser

	newUserCompany := onboardingdomain.UserCompany{
		CompanyID: company.ID,
		UserID:    user.ID,
		CreatedBy: company.CreatedBy,
		UpdatedBy: company.UpdatedBy,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}

	createdUserCompany, createdUserCompanyErr := c.onboardingRepository.Create(context.Background(), newUserCompany)
	if createdUserCompanyErr != nil {
		return nil, fmt.Errorf("failed to create user company: %w", createdUserCompanyErr)
	}

	userCompany, userCompanyErr := createdUserCompany.ToDomain()
	if userCompanyErr != nil {
		return nil, fmt.Errorf("failed to convert user company result to domain: %w", userCompanyErr)
	}

	return &userCompany, nil
}
