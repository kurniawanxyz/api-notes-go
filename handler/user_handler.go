package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"github.com/kurniawanxyz/crud-notes-go/helper"
	"github.com/kurniawanxyz/crud-notes-go/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	useCase *usecase.UserUseCase 
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var data domain.User
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	if err := helper.HandleValidation(c, &data); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err)
		return
	}

	result, err := h.useCase.Store(data)

	if err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.HandleResponse(c, http.StatusCreated, result)
}

func (h *UserHandler) Login(c *gin.Context){
	var request domain.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	if err := helper.HandleValidation(c, &request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.useCase.FindByEmail(request.Email)

	if err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, "Password is incorrect")
		return
	}

	token, err := helper.GenerateJWT(&user)

	if err != nil {
		helper.HandleResponse(c, http.StatusInternalServerError,gin.H{"jwt": err.Error()})
		return
	}
	helper.HandleResponse(c, http.StatusOK, gin.H{"token": token})
}