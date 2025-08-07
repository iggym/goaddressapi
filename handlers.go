package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// getAddressEndpoint returns a single address by its ID.
func getAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range addresses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Address{})
}

// getAddressesEndpoint returns all addresses.
func getAddressesEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(addresses)
}

// createAddressEndpoint creates a new address.
func createAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	var address Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// a simple way to generate a new ID
	address.ID = fmt.Sprintf("%d", len(addresses)+1)
	addresses = append(addresses, address)
	saveAddressesToCSV()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(address)
}

// updateAddressEndpoint updates an existing address.
func updateAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedAddress Address
	if err := json.NewDecoder(r.Body).Decode(&updatedAddress); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses[index] = updatedAddress
			addresses[index].ID = params["id"]
			saveAddressesToCSV()
			json.NewEncoder(w).Encode(addresses[index])
			return
		}
	}
	json.NewEncoder(w).Encode(&Address{})
}

// deleteAddressEndpoint deletes an address by its ID.
func deleteAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses = append(addresses[:index], addresses[index+1:]...)
			saveAddressesToCSV()
			break
		}
		json.NewEncoder(w).Encode(addresses)
	}
}

// 1MB
const maxMemory = 1 * 1024 * 1024

// importCSVFile handles importing addresses from a CSV file.
func importCSVFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("importing file")
	if err := os.RemoveAll(tmpdir); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := os.MkdirAll(tmpdir, 0700); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	for key, value := range r.MultipartForm.Value {
		fmt.Fprintf(w, "%s:%s ", key, value)
		log.Printf("%s:%s", key, value)
	}

	var path string
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			path = fmt.Sprintf("tmp/%s", fileHeader.Filename)
			buf, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := ioutil.WriteFile(path, buf, os.ModePerm); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(path)
		}
	}
	csvfile, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var address Address
	for _, record := range rawCSVdata {
		address.Firstname = record[0]
		address.Lastname = record[1]
		address.EmailAddress = record[2]
		address.PhoneNumber = record[3]
		addresses = append(addresses, address)
	}

	fmt.Println("Addresses from csv file uploaded ")
	w.WriteHeader(http.StatusOK)
}

// exportCSVFile handles exporting addresses to a CSV file.
func exportCSVFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("exporting addresses to file")
	deleteFile(csvfilepath)
	createFile(csvfilepath)
	//fmt.Println(addresses)
	var file, err = os.OpenFile(csvfilepath, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	for i := range addresses {
		line := fmt.Sprintf("%s,%s,%s,%s\r\n", addresses[i].Firstname, addresses[i].Lastname, addresses[i].EmailAddress, addresses[i].PhoneNumber)
		_, err = file.WriteString(line)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("%s,%s,%s,%s\n", addresses[i].Firstname, addresses[i].Lastname, addresses[i].EmailAddress, addresses[i].PhoneNumber)
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=addresses.csv")
	http.ServeFile(w, r, csvfilepath)
}

// handlers sets up the routes for the API.
func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/addresses", getAddressesEndpoint).Methods("GET")
	router.HandleFunc("/addresses", createAddressEndpoint).Methods("POST")
	router.HandleFunc("/addresses/{id}", getAddressEndpoint).Methods("GET")
	router.HandleFunc("/addresses/{id}", updateAddressEndpoint).Methods("PUT")
	router.HandleFunc("/addresses/{id}", deleteAddressEndpoint).Methods("DELETE")
	//upload and download
	router.HandleFunc("/import", importCSVFile)
	router.HandleFunc("/export", exportCSVFile)
	router.Handle("/", http.FileServer(http.Dir("static")))
	return router
}
