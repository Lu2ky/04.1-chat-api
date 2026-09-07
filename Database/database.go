package Database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func ConnectToDatabase() error {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	addr := os.Getenv("DB_ADDR")
	port := os.Getenv("DB_ADDR_PORT")
	dbname := os.Getenv("DB_NAME")

	if user == "" || addr == "" || port == "" || dbname == "" {
		return fmt.Errorf("falta alguna variable de entorno DB_USER/DB_ADDR/DB_ADDR_PORT/DB_NAME")
	}
	cfgDB := mysql.NewConfig()
	cfgDB.User = user
	cfgDB.Passwd = pass
	cfgDB.Net = "tcp"
	cfgDB.Addr = addr + ":" + port
	cfgDB.DBName = dbname
	var err error
	Connection, err = sql.Open("mysql", cfgDB.FormatDSN())
	if err != nil {
		log.Println("sql.Open error:", err)
		return err
	}
	if err = Connection.Ping(); err != nil {
		log.Println("Error ping DB:", err)
		return err
	}

	return nil
}