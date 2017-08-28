package main

import (
	"bytes"
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

func init() {
	server = httptest.NewServer(handlers())
	hostnamePath = fmt.Sprintf("%s/", server.URL)
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

	t.SkipNow()
	reader = strings.NewReader("")
	importURLPath := hostnamePath + "import"
	fmt.Println(importURLPath)
	request, err := http.NewRequest("POST", importURLPath, reader)

	request.ParseMultipartForm(0)
	//extraParams := map[string]string{
	//	"title":       "My test file",
	//	"author":      "iggy",
	//	"description": "import test file",
	//}
	//request, err := newfileUploadRequest(importURLPath, extraParams, "file", "tmp.csv")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
