package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

var args struct {
	Csv string `arg:"required"`
}

func main() {
	arg.MustParse(&args)
	filePath := args.Csv

	_, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Error getting file %s : %v\n", filePath, err)
	}
	log.Printf("Parsing file: %s\n", filePath)

	data, incorrectData, err := RetrieveDataFromFile(filePath)
	if err != nil {
		log.Fatalf("Error getting data from file %s : %v\n", filePath, err)
	}

	if len(*incorrectData) > 0 {
		log.Printf("Incorrect data found: %d elements\n", len(*incorrectData))
	}

	log.Printf("%d lines Collected.\n", len(data.Lines))

	insertedLines, err := StoreDataInSqliteDB(data)
	if err != nil {
		log.Fatalf("Error storing data in the database")
	}

	log.Printf("%d lines has been successfully inserted into database.\n", insertedLines)
}
