package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %s - error: %v\n", filePath, err)
	}

	return file
}

func RetrieveDataFromFile(filePath string) (*Data, *[]string, error) {
	file := openFile(filePath)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file %s: %v\n", filePath, err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	records := csv.NewReader(strings.NewReader(firstLine))
	firstLineRecords, err := records.Read()
	if err != nil {
		log.Fatalf("Error reading first line records - error: %v", err)
	}

	var data Data
	var incorrectData []string

	for scanner.Scan() {
		line := scanner.Text()
		records = csv.NewReader(strings.NewReader(line))
		lineRecords, err := records.Read()
		if err != nil {
			log.Fatalf("Error reading line records - error: %v", err)
		}

		if len(lineRecords) != len(firstLineRecords) {
			incorrectData = append(incorrectData, line)
			continue
		}

		lineData := map[string]string{}

		for index, item := range lineRecords {
			lineData[firstLineRecords[index]] = item
		}

		data.Lines = append(data.Lines, lineData)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner operation failed: %s", err)
	}

	return &data, &incorrectData, nil
}
