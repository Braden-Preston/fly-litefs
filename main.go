package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open database connection
	db, err := sql.Open("sqlite3", "/litefs/a.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create schema definitions
	stmtCreateTables := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY DEFAULT (LOWER(hex (randomblob (7)))) NOT NULL,
		name TEXT NOT NULL DEFAULT '',
		email TEXT NOT NULL DEFAULT ''
	)
	`
	_, err = db.Exec(stmtCreateTables)
	if err != nil {
		log.Printf("%q: %s\n", err, stmtCreateTables)
		return
	}

	// Create Echo web app
	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		// Query user data
		rows, err := db.Query("SELECT id, name, email from users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var (
				id    string
				name  string
				email string
			)
			err := rows.Scan(&id, &name, &email)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name, email)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, "hello litefs!")
	})

	fmt.Println("starting server!")
	if err := e.Start(":4000"); err != nil {
		log.Fatal(err)
	}
}
