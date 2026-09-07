package Types

type User struct {
	ID       string    `json:"idUser"`
}
type Message struct {
	ID        int    `json:"idMessage"`
	Chat_idChat   int    `json:"chatId"`
	Message  string    `json:"Message"`
	Enviado_por    string    `json:"Enviado_por"`
	Creado_en string `json:"timestamp"`
}
type Chat struct {
	ID        int    `json:"idChat"`
	Nombre    string `json:"Nombre"`
}
type UserChat struct {
	Chat_idChat int    `json:"chatId"`
	Usuario_idUsuario string `json:"idUser"`
	ID int `json:"id"`
}

type CreateChat struct{
	Nombre string `json:"Nombre"`
	IdUsuario string `json:"idUsuario"`
}
type CreateMessage struct{
	Message string `json:"Content"`
	Enviado_por string `json:"Enviado_por"`
	Enviado_en int `json:"Enviado_en"`

}