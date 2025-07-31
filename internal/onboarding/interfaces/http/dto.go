package onboardinghttp

type NewUserCompany struct {
	CompanyName  string `json:"companyName" validate:"required"`
	UserName     string `json:"userName" validate:"required"`
	UserEmail    string `json:"userEmail" validate:"required,email"`
	UserPassword string `json:"userPassword" validate:"required"`
}
