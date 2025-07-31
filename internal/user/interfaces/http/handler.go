package userhttp

import (
	"encoding/json"
	"net/http"

	response "github.com/joshuaalpuerto/go-rest-api/internal/common/http/response"
	validator "github.com/joshuaalpuerto/go-rest-api/internal/common/validator"

	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/user/domain"
	userusecases "github.com/joshuaalpuerto/go-rest-api/internal/user/usecases"
)

type UserHandler struct {
	userService userusecases.UserUsecase
	validator   validator.Validator
}

func NewUserHandler(userRepository userusecases.UserRepository, validator validator.Validator) UserHandler {
	return UserHandler{
		userService: userusecases.NewUserUsecase(userRepository),
		validator:   validator,
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user NewUser
	response := response.Response{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := h.validator.Validate(&user); errs != nil {
		response.SendErrorResponse(w, errs.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(userdomain.NewUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, createdUser, http.StatusCreated)
}
