package lexoffice

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	lexofficeBaseUrlV1   string = "https://api.lexoffice.io/v1"
	lexofficeDebugOutput bool   = false
)

type Client struct {
	apiKey     string
	baseURL    string
	debug      bool
	HTTPClient *http.Client
}

// NewClient creates new lexoffice.io client with given API key
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: lexofficeBaseUrlV1,
		debug:   lexofficeDebugOutput,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

// Content-type and body should be already added to req
func (c *Client) sendRequest(req *http.Request, result interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	if c.debug {
		fmt.Printf("\n%s\n\n", prettyPrintRequest(req))
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if c.debug {
		fmt.Printf("\n%s\n\n", prettyPrintResponse(res))
	}

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return err
	}

	return nil
}
