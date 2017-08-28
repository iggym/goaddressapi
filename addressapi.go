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

//Address data
type Address struct {
	ID           string `json:"id,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	EmailAddress string `json:"emailaddress,omitempty"`
	PhoneNumber  string `json:"phonenumber,omitempty"`
}

var addresses []Address

const csvfilepath = "addresses.csv"
const tmpdir = "tmp"

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
func getAddressesEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(addresses)
}
func createAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var address Address
	_ = json.NewDecoder(r.Body).Decode(&address)
	address.ID = params["id"]
	addresses = append(addresses, address)
	json.NewEncoder(w).Encode(addresses)
	w.WriteHeader(http.StatusCreated)
}
func updateAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var address Address
	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses = append(addresses[:index], addresses[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(addresses)
	}
	_ = json.NewDecoder(r.Body).Decode(&address)
	address.ID = params["id"]
	addresses = append(addresses, address)
	json.NewEncoder(w).Encode(addresses)

}
func deleteAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses = append(addresses[:index], addresses[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(addresses)
	}
}

// 1MB
const maxMemory = 1 * 1024 * 1024

func importCSVFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploading file")
	os.RemoveAll(tmpdir)
	os.MkdirAll(tmpdir, 0700)
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for key, value := range r.MultipartForm.Value {
		fmt.Fprintf(w, "%s:%s ", key, value)
		log.Printf("%s:%s", key, value)
	}

	var path string
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path = fmt.Sprintf("tmp/%s", fileHeader.Filename)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
			fmt.Println(path)
		}
	}
	csvfile, err := os.Open(path)
	checkImportError(err)
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	fmt.Println(addresses)

}
func checkImportError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func exportCSVFile(w http.ResponseWriter, r *http.Request) {

	deleteFile(csvfilepath)
	createFile()
	//fmt.Println(addresses)
	var file, err = os.OpenFile(csvfilepath, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	for i := range addresses {
		line := fmt.Sprintf("%s,%s,%s,%s\r\n", addresses[i].Firstname, addresses[i].Lastname, addresses[i].EmailAddress, addresses[i].PhoneNumber)
		_, err = file.WriteString(line)
		if isError(err) {
			return
		}
		fmt.Printf("%s,%s,%s,%s\n", addresses[i].Firstname, addresses[i].Lastname, addresses[i].EmailAddress, addresses[i].PhoneNumber)
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=addresses.csv")
	http.ServeFile(w, r, csvfilepath)
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func deleteFile(f string) {
	// delete file
	var err = os.Remove(csvfilepath)
	if isError(err) {
		return
	}
	fmt.Println("==> done deleting addresses file")
}
func createFile() {
	// detect if file exists
	var _, err = os.Stat(csvfilepath)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(csvfilepath)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> done creating file", csvfilepath)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/addresses", getAddressesEndpoint).Methods("GET")
	router.HandleFunc("/addresses/{id}", getAddressEndpoint).Methods("GET")
	router.HandleFunc("/addresses/{id}", createAddressEndpoint).Methods("POST")
	router.HandleFunc("/addresses/{id}", updateAddressEndpoint).Methods("PUT")
	router.HandleFunc("/addresses/{id}", deleteAddressEndpoint).Methods("DELETE")
	//upload and download
	router.HandleFunc("/import", importCSVFile)
	router.HandleFunc("/export", exportCSVFile)
	router.Handle("/", http.FileServer(http.Dir("static")))
	return router
}

func main() {

	addresses = append(addresses, Address{ID: "1", Firstname: "Bryan", Lastname: "Connerty", EmailAddress: "bryanc@test-email.com", PhoneNumber: "214-555-5551"})
	addresses = append(addresses, Address{ID: "2", Firstname: "Kate", Lastname: "Sacker", EmailAddress: "kates@test-email.com", PhoneNumber: "214-555-5552"})
	addresses = append(addresses, Address{ID: "3", Firstname: "Lonnie", Lastname: "Watley", EmailAddress: "lonniew@test-email.com", PhoneNumber: "214-555-5553"})
	addresses = append(addresses, Address{ID: "4", Firstname: "Ira", Lastname: "Schirmer", EmailAddress: "iras@test-email", PhoneNumber: "214-555-5554"})
	addresses = append(addresses, Address{ID: "5", Firstname: "Adam", Lastname: "DeGiulio", EmailAddress: "adamgs@test-email", PhoneNumber: "214-555-5554"})

	log.Fatal(http.ListenAndServe(":8001", handlers()))
}
