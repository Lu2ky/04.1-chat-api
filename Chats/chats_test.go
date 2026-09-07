package Chats

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	Database "04.1-chat-api/Database"
	"04.1-chat-api/Types"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestGetMessagesChats(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creando sqlmock: %v", err)
    }
    defer db.Close()
    Database.Connection = db

    rows := sqlmock.NewRows([]string{"idMessage", "creado_en", "message", "enviado_por", "Chat_idChat"}).
        AddRow(1, "2023-01-01", "hola", "user1", 1)
    mock.ExpectQuery("SELECT m.idMessage").WithArgs("1").WillReturnRows(rows)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{{Key: "idChat", Value: "1"}}

    GetMessagesChats(c)

    if w.Code != 200 {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    var msgs []Types.Message
    if err := json.Unmarshal(w.Body.Bytes(), &msgs); err != nil {
        t.Fatalf("error decoding response: %v", err)
    }
    if len(msgs) != 1 {
        t.Fatalf("expected 1 message, got %d", len(msgs))
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations not met: %v", err)
    }
}

func TestCreateChat(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creando sqlmock: %v", err)
    }
    defer db.Close()
    Database.Connection = db

    mock.ExpectExec("INSERT INTO Chat").WithArgs("TestChat").WillReturnResult(sqlmock.NewResult(10, 1))
    mock.ExpectExec("INSERT INTO chats_usuario").WithArgs(int64(10), "userX").WillReturnResult(sqlmock.NewResult(1, 1))

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    jsonStr := `{"Nombre":"TestChat","IdUsuario":"userX"}`
    c.Request = httptest.NewRequest("POST", "/chats/create", strings.NewReader(jsonStr))
    c.Request.Header.Set("Content-Type", "application/json")

    CreateChat(c)

    if w.Code != 200 {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations not met: %v", err)
    }
}
