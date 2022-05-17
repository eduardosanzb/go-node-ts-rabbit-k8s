package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbSchema   = "DB_SCHEMA"
)

var (
	DB       *sql.DB
	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host     = os.Getenv(dbHost)
	schema   = os.Getenv(dbSchema)
)

type DBClient interface {
	Execute(statement string)
	Query(statement string) *sql.Rows
}

func init() {
	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, username, password, schema)

	fmt.Println("hola")
	fmt.Println(connInfo)

	var err error
	DB, err = sql.Open("postgres", connInfo)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// defer DB.Close()

	if err = DB.Ping(); err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("Database ready to accept connections")
}
