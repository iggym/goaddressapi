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
