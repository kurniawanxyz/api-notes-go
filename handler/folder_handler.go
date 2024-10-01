package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"github.com/kurniawanxyz/crud-notes-go/helper"
	"github.com/kurniawanxyz/crud-notes-go/usecase"
)

type FolderHandler struct {
	useCase *usecase.FolderUseCase
}

func NewFolderHandler(useCase *usecase.FolderUseCase) *FolderHandler {
	return &FolderHandler{useCase: useCase}
}

func (h *FolderHandler) Index(c *gin.Context){
	user := helper.GetUserFromContext(c)
	folders, err := h.useCase.Index(user.ID)
	if err != nil {
		helper.HandleResponse(c,500, err.Error())
		return
	}
	
	helper.HandleResponse(c,200, folders)
}

func (h *FolderHandler) Show(c *gin.Context){
	user := helper.GetUserFromContext(c)
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		helper.HandleResponse(c, 400, "Invalid folder ID")
		return
	}
	folder, err := h.useCase.Show(intID, user.ID)
	if err != nil {
		helper.HandleResponse(c,500, err.Error())
		return
	}
	
	helper.HandleResponse(c,200, folder)
}

func (h *FolderHandler) Store(c *gin.Context) {
	var request domain.Folder
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user := helper.GetUserFromContext(c)
	request.UserID = user.ID

	if err := helper.HandleValidation(c, &request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err)
		return
	}

	result, err := h.useCase.Store(request)
	if err != nil {
		helper.HandleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HandleResponse(c, http.StatusCreated, result)
}

func (h *FolderHandler) Update(c *gin.Context) {
	var request domain.UpdateFolderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := helper.HandleValidation(c, &request); err != nil {
		helper.HandleResponse(c, http.StatusBadRequest, err)
		return
	}

	user := helper.GetUserFromContext(c)
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.useCase.Update(user.ID, id, &request)

	if err != nil {
		helper.HandleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HandleResponse(c, http.StatusOK, result)

}