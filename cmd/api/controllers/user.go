package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/joshuaalpuerto/go-rest-api/cmd/api/response"
	userusecases "github.com/joshuaalpuerto/go-rest-api/internal/user/usecases"
)

type UserHandler struct {
	userService userusecases.UserUsecase
	validator   Validator
}

func NewUserController(userRepository userusecases.UserRepository, validator Validator) UserHandler {
	return UserHandler{
		userService: userusecases.NewUserUsecase(userRepository),
		validator:   validator,
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userusecases.NewUser
	response := response.Response{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := h.validator.Validate(&user); errs != nil {
		response.SendErrorResponse(w, errs.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, createdUser, http.StatusCreated)
}
