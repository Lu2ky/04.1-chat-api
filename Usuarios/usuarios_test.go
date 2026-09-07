package Usuarios

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

func TestGetUserChats(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creando sqlmock: %v", err)
    }
    defer db.Close()
    Database.Connection = db

    rows := sqlmock.NewRows([]string{"idchats_usuario", "Usuario_idUsuario", "Chat_idChat"}).
        AddRow(5, "u1", 2)
    mock.ExpectQuery("SELECT cu.idchats_usuario").WithArgs("u1").WillReturnRows(rows)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{{Key: "idUser", Value: "u1"}}

    GetUserChats(c)

    if w.Code != 200 {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    var chats []Types.UserChat
    if err := json.Unmarshal(w.Body.Bytes(), &chats); err != nil {
        t.Fatalf("error decoding response: %v", err)
    }
    if len(chats) != 1 {
        t.Fatalf("expected 1 chat, got %d", len(chats))
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations not met: %v", err)
    }
}

func TestCreateUser(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creando sqlmock: %v", err)
    }
    defer db.Close()
    Database.Connection = db

    mock.ExpectExec("INSERT INTO Usuario").WithArgs("u2").WillReturnResult(sqlmock.NewResult(1, 1))

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    jsonStr := `{"idUser":"u2"}`
    c.Request = httptest.NewRequest("POST", "/usuarios/create", strings.NewReader(jsonStr))
    c.Request.Header.Set("Content-Type", "application/json")

    CreateUser(c)

    if w.Code != 200 {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    var resp map[string]string
    if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
        t.Fatalf("error decoding response: %v", err)
    }
    if resp["message"] != "Usuario creado exitosamente" {
        t.Fatalf("unexpected message: %v", resp)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations not met: %v", err)
    }
}
