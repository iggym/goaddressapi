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
	server   *httptest.Server
	reader   io.Reader
	addressesUrl string
)

func init() {
	server = httptest.NewServer(Handlers())

	addressesUrl = fmt.Sprintf("%s/addresses", server.URL)
}

func TestGetAddress(t *testing.T) {
	addressJson := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJson)
	addressesUrlParam := addressesUrl + "/101"
	request, err := http.NewRequest("POST", addressesUrlParam , reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestCreateAddress(t *testing.T) {
	addressJson := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJson)
	addressesUrlParam := addressesUrl + "/101"
	request, err := http.NewRequest("POST", addressesUrlParam , reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestUpdateAddress(t *testing.T) {
	addressJson := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJson)
	addressesUrlParam := addressesUrl + "/101"
	request, err := http.NewRequest("POST", addressesUrlParam , reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestDeleteAddress(t *testing.T) {
	addressJson := `{"Firstname":"John","lastname":"Doe","emailaddress":"john@test.com","phonenumber":"214-555-5553"}`

	reader = strings.NewReader(addressJson)
	addressesUrlParam := addressesUrl + "/101"
	request, err := http.NewRequest("POST", addressesUrlParam , reader)

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

	request, err := http.NewRequest("GET", addressesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}