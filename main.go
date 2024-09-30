package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kurniawanxyz/crud-notes-go/config"
	"github.com/kurniawanxyz/crud-notes-go/db"
	"github.com/kurniawanxyz/crud-notes-go/handler"
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

	gin := gin.Default()
	r := gin.Group("/api")

	// user
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	



	// user
	port := fmt.Sprintf(":%s", config.ENV.ServerPort)
	gin.Run(port)
}
