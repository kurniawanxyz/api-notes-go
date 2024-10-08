package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/config"
	"github.com/kurniawanxyz/crud-notes-go/db"
	"github.com/kurniawanxyz/crud-notes-go/handler"
	"github.com/kurniawanxyz/crud-notes-go/helper"
	"github.com/kurniawanxyz/crud-notes-go/repository"
	"github.com/kurniawanxyz/crud-notes-go/usecase"
)

func main() {
	// load env
	config.LoadConfig()

	// init db
	db.Init()


	// user
	userRepo := repository.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// folder
	folderRepo := repository.NewFolderRepository(db.DB)
	folderUseCase := usecase.NewFolderUseCase(folderRepo)
	folderHandler := handler.NewFolderHandler(folderUseCase)

	// note
	noteRepo := repository.NewNoteRepository(db.DB)
	noteUseCase := usecase.NewNoteUseCase(noteRepo)
	noteHandler := handler.NewNoteHandler(noteUseCase)



	gin := gin.Default()
	r := gin.Group("/api")

	// user
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// folder
	folder := r.Group("/folder").Use(helper.JWTAuthMiddleware())
	folder.GET("/", folderHandler.Index)
	folder.GET("/:id", folderHandler.Show)
	folder.POST("/store", folderHandler.Store)
	folder.PUT("/:id", folderHandler.Update)
	folder.DELETE("/:id", folderHandler.Delete)

	// note
	note := r.Group("/note").Use(helper.JWTAuthMiddleware())
	note.GET("/:folder_id", noteHandler.Index)
	note.GET("/:folder_id/:id", noteHandler.Show)
	note.POST("/:folder_id/store", noteHandler.Store)
	note.PUT("/:folder_id/:id", noteHandler.Update)
	note.DELETE("/:folder_id/:id", noteHandler.Delete)

	// user
	port := fmt.Sprintf(":%s", config.ENV.ServerPort)
	gin.Run(port)
}
