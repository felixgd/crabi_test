package pld

import (
	"bytes"
	"crabi_test/domain"
	"crabi_test/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Repositories interface {
	GetUserInPLD(*domain.PLDPayload) (*domain.PLD, *errors.APIError)
}

// Request represents the HTTP request handler.
type PLD struct{}

const endpoint = "/check-blacklist"

// Get sends an HTTP POST request to the specified URL with optional headers.
func (r *PLD) GetUserInPLD(payload *domain.PLDPayload) (*domain.PLD, *errors.APIError) {
	// Convert struct to JSON bytes
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Err:     err,
		}
	}

	// Create an io.Reader from the JSON bytes
	reader := bytes.NewBuffer(jsonBytes)

	log.Println(reader)
	req, err := http.NewRequest("POST", "http://localhost:3000/check-blacklist", reader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		return nil, &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Err:     err,
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return nil, &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Err:     err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		if err != nil {
			return nil, &errors.APIError{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
				Err:     err,
			}
		}

		return nil, &errors.APIError{
			Code:    resp.StatusCode,
			Message: resp.Status,
			Err:     fmt.Errorf(resp.Status),
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Err:     err,
		}
	}

	pld := domain.PLD{}

	if err := json.Unmarshal(body, &pld); err != nil {
		return nil, &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Err:     err,
		}
	}

	return &pld, nil
}
