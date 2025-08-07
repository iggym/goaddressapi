package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	server       *httptest.Server
	reader       io.Reader
	addressesURL string
	hostnamePath string
)

func resetData() {
	addresses = []Address{
		{ID: "1", Firstname: "Bryan", Lastname: "Connerty", EmailAddress: "bryanc@test-email.com", PhoneNumber: "214-555-5551"},
		{ID: "2", Firstname: "Kate", Lastname: "Sacker", EmailAddress: "kates@test-email.com", PhoneNumber: "214-555-5552"},
		{ID: "3", Firstname: "Lonnie", Lastname: "Watley", EmailAddress: "lonniew@test-email.com", PhoneNumber: "214-555-5553"},
		{ID: "4", Firstname: "Ira", Lastname: "Schirmer", EmailAddress: "iras@test-email", PhoneNumber: "214-555-5554"},
		{ID: "5", Firstname: "Adam", Lastname: "DeGiulio", EmailAddress: "adamgs@test-email", PhoneNumber: "214-555-5554"},
	}
}

func init() {
	server = httptest.NewServer(handlers())
	hostnamePath = fmt.Sprintf("%s/", server.URL)
	addressesURL = fmt.Sprintf("%s/addresses", server.URL)
	resetData()
}

func TestGetAddress(t *testing.T) {
	resetData()
	reader = strings.NewReader("")
	// assuming address with ID 1 exists from the initial data
	addressesURLParam := addressesURL + "/1"
	request, err := http.NewRequest("GET", addressesURLParam, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	var address Address
	json.NewDecoder(res.Body).Decode(&address)

	if address.Firstname != "Bryan" {
		t.Errorf("Expected firstname to be 'Bryan', but got '%s'", address.Firstname)
	}
}
func TestCreateAddress(t *testing.T) {
	resetData()
	addressJSON := `{"firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJSON)
	request, err := http.NewRequest("POST", addressesURL, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	var createdAddress Address
	json.NewDecoder(res.Body).Decode(&createdAddress)

	if createdAddress.Firstname != "John" {
		t.Errorf("Expected firstname to be 'John', but got '%s'", createdAddress.Firstname)
	}
}
func TestUpdateAddress(t *testing.T) {
	resetData()
	addressJSON := `{"firstname":"Updated","lastname":"Name","emailaddress":"updated@test.com","phonenumber":"214-555-0000"}`

	reader = strings.NewReader(addressJSON)
	// assuming address with ID 1 exists from the initial data
	addressesURLParam := addressesURL + "/1"
	request, err := http.NewRequest("PUT", addressesURLParam, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	var updatedAddress Address
	json.NewDecoder(res.Body).Decode(&updatedAddress)

	if updatedAddress.Firstname != "Updated" {
		t.Errorf("Expected firstname to be 'Updated', but got '%s'", updatedAddress.Firstname)
	}
}
func TestDeleteAddress(t *testing.T) {
	resetData()
	reader = strings.NewReader("")
	addressesURLParam := addressesURL + "/1"
	request, err := http.NewRequest("DELETE", addressesURLParam, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestListAddresses(t *testing.T) {
	resetData()
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", addressesURL, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	var addresses []Address
	json.NewDecoder(res.Body).Decode(&addresses)

	// The initial data has 5 addresses
	if len(addresses) != 5 {
		t.Errorf("Expected at least 5 addresses, but got %d", len(addresses))
	}
}

func TestExport(t *testing.T) {
	resetData()
	reader = strings.NewReader("")
	exportURLPath := hostnamePath + "export"
	request, err := http.NewRequest("GET", exportURLPath, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestImport(t *testing.T) {
	resetData()
	reader = strings.NewReader("")
	importURLPath := hostnamePath + "import"

	// Create a temporary file with some CSV data
	tmpfile, err := ioutil.TempFile("import", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	text := []byte("test,import,test@import.com,123-456-7890")
	if _, err := tmpfile.Write(text); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	extraParams := map[string]string{}
	request, err := newfileUploadRequest(importURLPath, extraParams, "data", tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
