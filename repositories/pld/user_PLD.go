package pld

import (
	"crabi_test/utils/errors"
	"io/ioutil"
	"net/http"
)

type Repositories interface {
	GetUserInPLD(string, map[string]string) (string, error)
}

// Request represents the HTTP request handler.
type PLD struct{}

// Get sends an HTTP GET request to the specified URL with optional headers.
func (r *PLD) GetUserInPLD(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set additional headers if provided.
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.APIError{
			Code:    resp.StatusCode,
			Message: resp.Status,
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
