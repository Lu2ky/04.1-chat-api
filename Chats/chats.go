package Chats

import (
	Database "04.1-chat-api/Database"
	"04.1-chat-api/Types"
	"github.com/gin-gonic/gin"
)

func GetMessagesChats(context *gin.Context){
	idChat := context.Param("idChat")
	query := `SELECT m.idMessage, m.creado_en, m.message, m.enviado_por, m.Chat_idChat FROM messages m WHERE m.Chat_idChat = ?`;
	rows ,err := Database.Connection.Query(query, idChat)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error al obtener mensajes del chat"})
		return;
	}
	defer rows.Close()
	var messages []Types.Message
	for rows.Next() {
		var message Types.Message
		if err := rows.Scan(&message.ID,&message.Creado_en,&message.Message,&message.Enviado_por,&message.Chat_idChat); err != nil {
			context.JSON(500, gin.H{"error": "error al ajustar los datos al tipo"})
			return;
		}
		messages = append(messages, message)
	}
	if err:= rows.Err(); err!=nil{
		context.JSON(500, gin.H{"error": "Error en la iteración"})
	}
	context.JSON(200,messages)
}
func CreateChat(context *gin.Context){
	var CreateChat Types.CreateChat
	if err := context.ShouldBindJSON(&CreateChat); err != nil{
		context.JSON(400, gin.H{"error": "Error al parsear el JSON"})
		return;
	}
	
	result ,err := Database.Connection.Exec(`INSERT INTO Chat (nombre) VALUE (?)`, CreateChat.Nombre)
	if err != nil {
		context.JSON(500,gin.H{"error":"error al insertar chat"})
	}
	lastID, err:= result.LastInsertId()
	_, err1 := Database.Connection.Exec(`INSERT INTO chats_usuario (Chat_idChat, Usuario_idUsuario) VALUE (?,?)`, lastID, CreateChat.IdUsuario);
	if err1 != nil {
		context.JSON(500,gin.H{"error":"error al insertar registro compartido usuario-chat"})
	}
	context.JSON(200, gin.H{"exito":"Registro insertado correctamente"})
}