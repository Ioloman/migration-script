package mysql

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupDB() {
	uri := os.Getenv("DB_URI")
	log.Printf("Connecting to %s", uri)
	DB = sqlx.MustConnect("mysql", uri)
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(30)
	log.Println("Connected to mysql")
}
