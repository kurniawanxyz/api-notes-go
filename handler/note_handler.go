package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"github.com/kurniawanxyz/crud-notes-go/helper"
	"github.com/kurniawanxyz/crud-notes-go/usecase"
)


type NoteHandler struct {
	NoteUseCase *usecase.NoteUseCase
}

func NewNoteHandler(usecase *usecase.NoteUseCase) *NoteHandler{
	return &NoteHandler{NoteUseCase: usecase}
}

func (h *NoteHandler) Index(c *gin.Context){
	user := helper.GetUserFromContext(c)
	folderID, _ := strconv.Atoi(c.Param("folder_id"))
	notes, err := h.NoteUseCase.Index(user.ID, folderID)

	if len(notes) == 0 {
		helper.HandleResponse(c, 404, gin.H{
			"message": "Note not found",
		})
		return
	}

	if err != nil {
		helper.HandleResponse(c, 500, err.Error())
		return
	}
	helper.HandleResponse(c, 200, notes)
}

func (h *NoteHandler) Store(c *gin.Context) {
	user := helper.GetUserFromContext(c)
	folderID, _ := strconv.Atoi(c.Param("folder_id"))

	var request domain.Note = domain.Note{
		UserID: user.ID,
		FolderID: folderID,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.HandleResponse(c, 400, err.Error())
		return
	}

	if err := helper.HandleValidation(c, &request); err != nil {
		helper.HandleResponse(c, 400, err)
		return
	}

	note, err := h.NoteUseCase.Store(&request)
	if err != nil {
		helper.HandleResponse(c, 500, err.Error())
		return
	}

	helper.HandleResponse(c, 201, note)
}

func (h *NoteHandler) Show(c *gin.Context) {
	user := helper.GetUserFromContext(c)
	folderID, _ := strconv.Atoi(c.Param("folder_id"))
	noteID, _ := strconv.Atoi(c.Param("id"))

	note, err := h.NoteUseCase.Show(user.ID, folderID, noteID)
	if err != nil {
		helper.HandleResponse(c, 500, err.Error())
		return
	}

	helper.HandleResponse(c, 200, note)
}

func (h *NoteHandler) Update(c *gin.Context) {
	user := helper.GetUserFromContext(c)
	folderID, _ := strconv.Atoi(c.Param("folder_id"))
	noteID, _ := strconv.Atoi(c.Param("id"))

	var request domain.Note = domain.Note{
		UserID: user.ID,
		FolderID: folderID,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.HandleResponse(c, 400, err.Error())
		return
	}

	if err := helper.HandleValidation(c, &request); err != nil {
		helper.HandleResponse(c, 400, err)
		return
	}

	note, err := h.NoteUseCase.Update(user.ID, folderID, noteID, &request)
	if err != nil {
		helper.HandleResponse(c, 500, err.Error())
		return
	}

	helper.HandleResponse(c, 200, note)
}

func (h *NoteHandler) Delete(c *gin.Context) {
	user := helper.GetUserFromContext(c)
	folderID, _ := strconv.Atoi(c.Param("folder_id"))
	noteID, _ := strconv.Atoi(c.Param("id"))

	err := h.NoteUseCase.Delete(user.ID, folderID, noteID)
	if err != nil {
		helper.HandleResponse(c, 500, err.Error())
		return
	}

	helper.HandleResponse(c, 200, gin.H{
		"message": "Note deleted",
	})
}