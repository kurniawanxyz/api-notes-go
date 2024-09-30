package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"github.com/kurniawanxyz/crud-notes-go/helper"
	"github.com/kurniawanxyz/crud-notes-go/usecase"
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
	
	if err := helper.HandleValidation(c, data); err != nil {
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