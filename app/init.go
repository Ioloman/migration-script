package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Ioloman/migration-script/app/db/mongodb"
	"github.com/Ioloman/migration-script/app/db/mysql"
	"github.com/joho/godotenv"
)

func getLogFile(name string) *os.File {
	path := filepath.Join(os.Getenv("LOG_DIR"), name)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err.Error())
	}
	return file
}

func init() {
	// read .env
	log.Println("Loading environment variables")
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file is not located")
	}

	// setup logger
	log.Println("Setting up logger")
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	if os.Getenv("APP_ENV") == "prod" {
		file := getLogFile("general.log")
		log.SetOutput(file)
	}

	mysql.SetupDB()
	mongodb.SetupDB()
}
