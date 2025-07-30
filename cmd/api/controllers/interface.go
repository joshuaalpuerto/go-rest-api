package controllers

type Validator interface {
	Validate(model any) error
}
