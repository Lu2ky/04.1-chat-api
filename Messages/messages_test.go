package Messages

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	Database "04.1-chat-api/Database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestCreateMessage(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creando sqlmock: %v", err)
    }
    defer db.Close()
    Database.Connection = db

    mock.ExpectExec("INSERT INTO messages").WithArgs("hola", "u1", 2).WillReturnResult(sqlmock.NewResult(1, 1))

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    jsonStr := `{"Content":"hola","Enviado_por":"u1","Enviado_en":2}`
    c.Request = httptest.NewRequest("POST", "/messages/create", strings.NewReader(jsonStr))
    c.Request.Header.Set("Content-Type", "application/json")

    CreateMessage(c)

    if w.Code != 200 {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    var resp map[string]string
    if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
        t.Fatalf("error decoding response: %v", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations not met: %v", err)
    }
}
