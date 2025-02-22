package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"fmt"
	"log"
	"strings"
)

func StoreDataInSqliteDB(data *Data, primaryKey *string) (int, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database: %v", err)
		}
	}()

	var columns []string
	for key := range data.Lines[0] {
		if len(*primaryKey) > 0 && key == *primaryKey {
			columns = append(columns, fmt.Sprintf("%s TEXT PRIMARY KEY", key))
		} else {
			columns = append(columns, fmt.Sprintf("%s TEXT", key))
		}
	}

	stmt, err := db.Prepare(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS data_table (%s)`, strings.Join(columns, ", ")))
	if err != nil {
		log.Fatalf("Error preaparing table creation: %v", err)
		return 0, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error executing table creation statement: %v", err)
		return 0, err
	}

	err = stmt.Close()
	if err != nil {
		log.Fatalf("Error closing create table statement- erro: %v\n", err)
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error starting database transaction: %v\n", err)
		return 0, err
	}

	insertedLines := 0
	for _, line := range data.Lines {
		var (
			columns []string
			values  []interface{}
			places  []string
		)
		for col, val := range line {
			columns = append(columns, col)
			values = append(values, val)
			places = append(places, "?")
		}
		query := fmt.Sprintf(
			"INSERT INTO data_table(%s) VALUES(%s);",
			strings.Join(columns, ","),
			strings.Join(places, ","),
		)
		stmt, err := db.Prepare(query)
		if err != nil {
			fmt.Printf("Error preparing the insert statement: %v\n", err)
			continue
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			// Handle errors here
			fmt.Printf("Error executing the insert statement: %v\n", err)
		}

		err = stmt.Close()
		if err != nil {
			// handle error
			fmt.Println("Error closing the statement: ", err)
		}

		insertedLines++
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error commiting line on db")
		return 0, err
	}

	return insertedLines, nil
}
