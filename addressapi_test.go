package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server       *httptest.Server
	reader       io.Reader
	addressesURL string
)

func init() {
	server = httptest.NewServer(handlers())

	addressesURL = fmt.Sprintf("%s/addresses", server.URL)
}

func TestGetAddress(t *testing.T) {
	reader = strings.NewReader("")
	addressesURLParam := addressesURL + "/101"
	request, err := http.NewRequest("GET", addressesURLParam, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestCreateAddress(t *testing.T) {
	addressJSON := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJSON)
	addressesURLParam := addressesURL + "/101"
	request, err := http.NewRequest("POST", addressesURLParam, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestUpdateAddress(t *testing.T) {
	addressJSON := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJSON)
	addressesURLParam := addressesURL + "/101"
	request, err := http.NewRequest("PUT", addressesURLParam, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestDeleteAddress(t *testing.T) {
	reader = strings.NewReader("")
	addressesURLParam := addressesURL + "/101"
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
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", addressesURL, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestExport(t *testing.T) {
	reader = strings.NewReader("")
	addressesURLPath := addressesURL + "/export"
	request, err := http.NewRequest("GET", addressesURLPath, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestImport(t *testing.T) {
	reader = strings.NewReader("")
	addressesURLPath := addressesURL + "/export"
	request, err := http.NewRequest("GET", addressesURLPath, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
