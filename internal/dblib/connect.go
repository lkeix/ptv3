package dblib

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Connect connect Postgres SQL DB server
func Connect() *sql.DB {
	godotenv.Load("./.env")
	host := "127.0.0.1"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB")
	db, err := sql.Open("postgres", "host="+host+" port=5432 user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
