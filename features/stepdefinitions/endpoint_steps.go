package stepdefinitions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/basic-golang-api-automation/features/helper"
	"github.com/yalp/jsonpath"
)

var urlEndpoint string
var getEndpoint, postEndpoint, putEndpoint, deleteEndpoint *http.Request
var getResponse, postResponse, putResponse, deleteResponse *http.Response
var id, jsonResponse interface{}

// GivenEndpoint : define endpoint
func GivenEndpoint(endpoint string) error {
	urlEndpoint = os.Getenv("API_ENDPOINT") + endpoint

	return nil
}

// GetIDCRUD : get ID for PUT and DELETE
func GetIDCRUD() error {
	jsonPath, _ := jsonpath.Prepare("$._id")
	responseBody, _ := ioutil.ReadAll(postResponse.Body)

	json.Unmarshal(responseBody, &jsonResponse)

	id, _ = jsonPath(jsonResponse)

	return nil
}

// GetEndpoint : hit endpoint get
func GetEndpoint() error {
	var err error

	client := &http.Client{}
	getEndpoint, err = http.NewRequest(http.MethodGet, urlEndpoint, nil)
	helper.LogPanicln(err)
	getResponse, err = client.Do(getEndpoint)
	helper.LogPanicln(err)

	return nil
}

// PostEndpoint : hit endpoint post
func PostEndpoint() error {
	var err error

	body := []byte(
		`{
			"name": "mpermperpisang",
			"age": 29,
			"colour": "blackred"
		}`)

	client := &http.Client{}
	postEndpoint, err := http.NewRequest(http.MethodPost, urlEndpoint, bytes.NewBuffer(body))
	helper.LogPanicln(err)

	postEndpoint.Header.Set("Content-Type", "application/json")

	postResponse, err = client.Do(postEndpoint)
	helper.LogPanicln(err)

	return nil
}

// PutEndpoint : hit endpoint put
func PutEndpoint() error {
	var err error

	body := []byte(
		`{
			"name": "mpermperpisang",
			"age": 30,
			"colour": "blackred"
		}`)

	client := &http.Client{}
	putEndpoint, err := http.NewRequest(http.MethodPut, urlEndpoint+"/"+fmt.Sprintf("%s", id), bytes.NewBuffer(body))
	helper.LogPanicln(err)

	putEndpoint.Header.Set("Content-Type", "application/json")

	putResponse, err = client.Do(putEndpoint)
	helper.LogPanicln(err)

	return nil
}

// DeleteEndpoint : hit endpoint delete
func DeleteEndpoint() error {
	var err error

	client := &http.Client{}
	deleteEndpoint, err := http.NewRequest(http.MethodDelete, urlEndpoint+"/"+fmt.Sprintf("%s", id), nil)
	helper.LogPanicln(err)
	deleteResponse, err = client.Do(deleteEndpoint)
	helper.LogPanicln(err)

	return nil
}
