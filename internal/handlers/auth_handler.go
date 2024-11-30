package handlers

import (
	"net/http"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/responses"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/services"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: *authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dtos.RegisterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrUsernameRequired.Error())
				return
			case "Email":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrEmailRequired.Error())
				return
			case "Password":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrPasswordIsRequired.Error())
				return
			}
		}
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = h.AuthService.Register(user)
	if err != nil {
		if err == errors.ErrEmailExist {
			responses.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		} else if err == errors.ErrUsernameExist {
			responses.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, "User registered successfully!", nil)

}
