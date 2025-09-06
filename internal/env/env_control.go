package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseCreds struct {
	Db_user string
	Db_pass string
	Db_url  string
	Db_name string
}

func GetDatabaseCreds() DatabaseCreds {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())

	}
	return DatabaseCreds{
		Db_user: os.Getenv("DB_USER"),
		Db_pass: os.Getenv("DB_PASS"),
		Db_url:  os.Getenv("DB_URL"),
		Db_name: os.Getenv("DB_NAME"),
	}

}
