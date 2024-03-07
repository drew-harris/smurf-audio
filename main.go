package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		panic("No database url loaded")
	}
	db, err := sqlx.Open("postgres", db_url)
	if err != nil {
		fmt.Println(err)
		panic("Could not open database")
	}
	fmt.Println(db.DB.Stats().InUse)

}
