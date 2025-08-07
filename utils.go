package main

import (
	"fmt"
	"os"
)

// deleteFile deletes the specified file.
func deleteFile(f string) {
	// delete file
	var err = os.Remove(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("==> done deleting addresses file")
}

// createFile creates a file if it does not already exist.
func createFile(f string) {
	// detect if file exists
	var _, err = os.Stat(f)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(f)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer file.Close()
	}
	fmt.Println("==> done creating file", f)
}

// saveAddressesToCSV saves the current list of addresses to the CSV file.
func saveAddressesToCSV() {
	deleteFile(csvfilepath)
	createFile(csvfilepath)
	var file, err = os.OpenFile(csvfilepath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file for saving", err)
		return
	}
	defer file.Close()

	for i := range addresses {
		line := fmt.Sprintf("%s,%s,%s,%s\r\n", addresses[i].Firstname, addresses[i].Lastname, addresses[i].EmailAddress, addresses[i].PhoneNumber)
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file for saving", err)
			return
		}
	}
}
