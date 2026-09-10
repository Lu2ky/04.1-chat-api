package Usuarios

import (
	"database/sql"
	"log"

	Database "04.1-chat-api/Database"
	"04.1-chat-api/Types"
	"github.com/gin-gonic/gin"
)

func GetUserChats(context *gin.Context){
	idUser := context.Param("idUser");
	query := `SELECT cu.idchats_usuario, cu.Usuario_idUsuario, cu.Chat_idChat, c.nombre FROM chats_usuario cu INNER JOIN Usuario u ON cu.Usuario_idUsuario = u.idUsuario INNER JOIN Chat c ON c.idChat = cu.Chat_idChat WHERE cu.Usuario_idUsuario = ?`
	rows, err := Database.Connection.Query(query, idUser)
	if err != nil {
		log.Println("Error al obtener los chats del usuario:", err)
		context.JSON(500, gin.H{"error": "Error al obtener los chats del usuario"})
		return
	}
	defer rows.Close()
	var chats []Types.UserChat
	for rows.Next() {
		var chat Types.UserChat
		if err := rows.Scan(&chat.ID, &chat.Usuario_idUsuario, &chat.Chat_idChat); err != nil {
			log.Println("Error escaneando fila de chats_usuarios:", err)
			context.JSON(500, gin.H{"error": "Error al escanear los chats del usuario"})
			return
		}		
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterando filas:", err)
		context.JSON(500, gin.H{"error": "Error al iterar los chats del usuario"})
		return
	}

	context.JSON(200, chats)
}
func CreateUser(context *gin.Context){
	var user Types.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": "Error al parsear el JSON"})
		return
	}
	_, err := Database.Connection.Exec("INSERT INTO Usuario (idUsuario, nombre) VALUES (?,?)", user.ID,user.Name)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error al crear el usuario"})
		return
	}
	context.JSON(200, gin.H{"message": "Usuario creado exitosamente"})
}	
func GetUser(context *gin.Context){
	idUser := context.Param("id")
	var user Types.User;
	query := `SELECT u.idUsuario, u.nombre FROM Usuario u WHERE u.idUsuario = ?`
	row := Database.Connection.QueryRow(query,idUser);
	err := row.Scan(&user.ID,&user.Name)
	if err != nil{
		if err == sql.ErrNoRows{
			context.JSON(400, gin.H{"Error":"usuario no encontrado"})
			return;
		}
		context.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	context.JSON(200,user)
}