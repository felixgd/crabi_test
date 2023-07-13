package repositories

import (
	"crabi_test/utils"
	"io/ioutil"
	"net/http"
)

// Request represents the HTTP request handler.
type Request struct {
	client    *http.Client
	authToken string
}

// NewRequest creates a new instance of the Request with the given client and auth token.
func NewRequest(client *http.Client, authToken string) Request {
	return Request{
		client:    client,
		authToken: authToken,
	}
}

// Get sends an HTTP GET request to the specified URL with optional headers.
func (r Request) Get(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the authorization header.
	req.Header.Set("Authorization", r.authToken)

	// Set additional headers if provided.
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", utils.APIError{
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
