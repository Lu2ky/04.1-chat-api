package Messages

import (
	Database "04.1-chat-api/Database"
	"04.1-chat-api/Types"
	"github.com/gin-gonic/gin"
)


func CreateMessage(context *gin.Context){
	var CreateMessage Types.CreateMessage
	if err := context.ShouldBindJSON(&CreateMessage); err != nil{
		context.JSON(400, gin.H{"error": "Error al parsear el JSON"})
		return;
	}
	
	_ ,err := Database.Connection.Exec(`INSERT INTO messages (message, enviado_por, Chat_idChat) VALUE (?,?,?)`, CreateMessage.Message, CreateMessage.Enviado_por, CreateMessage.Enviado_en)
	if err != nil {
		context.JSON(500,gin.H{"error":"error al insertar mensaje"})
	}
	context.JSON(200, gin.H{"exito":"Registro insertado correctamente"})
}