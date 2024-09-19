package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionString := "devbook_api:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	result, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	fmt.Println(result, err)
}
