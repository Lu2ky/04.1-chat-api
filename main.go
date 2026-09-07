package main

import (
	"log"

	"04.1-chat-api/Chats"
	"04.1-chat-api/Database"
	"04.1-chat-api/Messages"
	"04.1-chat-api/Usuarios"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el .env:", err)
	}
	if err := Database.ConnectToDatabase(); err != nil {
		log.Fatalf("Error conectando con la db: %v", err)
	}

	defer func() {
		if Database.Connection != nil {
			Database.Connection.Close()
		}
	}()

	router := gin.Default()
	v0 := router.Group("/api/")
	registerRoutes(v0)
	router.Run("0.0.0.0:8080")
}
func registerRoutes(router gin.IRouter){
	//Endpoints disponibles
	router.GET("/usuarios/chats/:idUser", Usuarios.GetUserChats)
	router.POST("/usuarios/create", Usuarios.CreateUser)
	router.GET("/chats/:idChat", Chats.GetMessagesChats)
	router.POST("/chats/create", Chats.CreateChat)
	router.POST("/messages/create", Messages.CreateMessage)
}