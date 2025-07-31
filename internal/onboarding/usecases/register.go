package onboardingusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/company/domain"
	onboardingdomain "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/domain"
	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/user/domain"
)

func (c *OnboardingUsecase) RegisterUserCompany(nwu onboardingdomain.NewUserCompany) (*onboardingdomain.UserCompany, error) {
	userId := "c8d9c08f-4f87-4c5c-8862-2f4abac75f1f"

	newCompany := companydomain.NewCompany{
		Name: nwu.CompanyName,
	}
	newCompany, companyErr := newCompany.ToDomainEntity(userId)
	if companyErr != nil {
		return nil, fmt.Errorf("failed to convert company to domain entity: %w", companyErr)
	}

	newUser := userdomain.NewUser{
		Name:     nwu.UserName,
		Email:    nwu.UserEmail,
		Password: nwu.UserPassword,
	}

	newUser, userErr := newUser.ToDomainEntity(userId)
	if userErr != nil {
		return nil, fmt.Errorf("failed to convert user to domain entity: %w", userErr)
	}

	createdCompany, companyRepoErr := c.companyRepository.Create(context.Background(), newCompany)
	if companyRepoErr != nil {
		return nil, fmt.Errorf("failed to create company: %w", companyRepoErr)
	}

	createdUser, userRepoErr := c.userRepository.Create(context.Background(), newUser)
	if userRepoErr != nil {
		return nil, fmt.Errorf("failed to create user: %w", userRepoErr)
	}

	company, companyErr := createdCompany.ToDomain()
	if companyErr != nil {
		return nil, fmt.Errorf("failed to convert company result to domain: %w", companyErr)
	}

	user, userErr := createdUser.ToDomain()
	if userErr != nil {
		return nil, fmt.Errorf("failed to convert user result to domain: %w", userErr)
	}

	newUserCompany := onboardingdomain.UserCompany{
		CompanyID: company.ID,
		UserID:    user.ID,
		CreatedBy: company.CreatedBy,
		UpdatedBy: company.UpdatedBy,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}

	createdUserCompany, userCompanyRepoErr := c.onboardingRepository.Create(context.Background(), newUserCompany)
	if userCompanyRepoErr != nil {
		return nil, fmt.Errorf("failed to create user company: %w", userCompanyRepoErr)
	}

	userCompany, userCompanyErr := createdUserCompany.ToDomain()
	if userCompanyErr != nil {
		return nil, fmt.Errorf("failed to convert user company result to domain: %w", userCompanyErr)
	}

	return &userCompany, nil
}
