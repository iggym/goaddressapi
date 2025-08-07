package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var addresses []Address

func loadInitialData() {
	file, err := os.Open("import/addresses.csv")
	if err != nil {
		log.Fatal("Error opening initial data file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading initial data file", err)
	}

	for i, record := range rawCSVdata {
		address := Address{
			ID:           strconv.Itoa(i + 1),
			Firstname:    record[0],
			Lastname:     record[1],
			EmailAddress: record[2],
			PhoneNumber:  record[3],
		}
		addresses = append(addresses, address)
	}
}

func main() {
	loadInitialData()
	fmt.Println("Server starting (serving port: 8001)")
	log.Fatal(http.ListenAndServe(":8001", handlers()))
}
