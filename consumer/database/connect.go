package consumer_database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


func ConnectToDataBase() error{

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("not found a database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("not pinnging the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	return nil
}