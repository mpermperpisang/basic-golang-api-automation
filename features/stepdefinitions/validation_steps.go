package stepdefinitions

import (
	"log"
)

// ValidateGetResponse : validate endpoint get response
func ValidateGetResponse() error {
	if getResponse.StatusCode != 200 {
		log.Panicln("GET - Error code tidak sesuai")
	}

	return nil
}

// ValidatePostResponse : validate endpoint post response
func ValidatePostResponse() error {
	if postResponse.StatusCode != 201 {
		log.Panicln("POST - Error code tidak sesuai")
	}

	return nil
}

// ValidatePutResponse : validate endpoint put response
func ValidatePutResponse() error {
	if putResponse.StatusCode != 200 {
		log.Panicln("PUT - Error code tidak sesuai")
	}

	return nil
}

// ValidateDeleteResponse : validate endpoint delete response
func ValidateDeleteResponse() error {
	if deleteResponse.StatusCode != 200 {
		log.Panicln("DELETE - Error code tidak sesuai")
	}

	return nil
}
