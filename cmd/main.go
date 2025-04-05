package main

import (
	"go-api/cmd/db"
	"go-api/cmd/repository"
	"go-api/cmd/usecase"
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	UserRepository := repository.NewUserRepository(dbConnection)
	//camada use case
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	//camada repository
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/users", UserController.GetUsers)
	server.Run(":8000") // Inicia o servidor na porta 8000
}
