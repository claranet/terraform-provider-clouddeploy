package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Username string
	Password string
	Endpoint string
}

type errorObject struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*errorObject, error) {
	var result map[string]errorObject
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", err)
	}
	s, ok := result["error"]
	if !ok {
		return nil, fmt.Errorf("JSON response does not have error field")
	}
	return &s, nil
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("Error calling the API endpoint: %v", err)
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		return resp, fmt.Errorf("Failed call API endpoint. HTTP response code: %v", resp.StatusCode)
	}
	return resp, nil
}

func (c *Client) do(method, path string, payload interface{}, headers map[string]string) (*http.Response, error) {
	url := c.Endpoint + path

	var body bytes.Buffer
	if payload != nil {
		data, err := json.Marshal(payload)
		if err == nil {
			body = *bytes.NewBuffer(data)
		}
	}

	req, _ := http.NewRequest(method, url, &body)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := netClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *Client) delete(path string, headers map[string]string) (*http.Response, error) {
	return c.do("DELETE", path, nil, headers)
}

func (c *Client) patch(path string, payload interface{}, headers map[string]string) (res *http.Response, err error) {
	return c.do("PATCH", path, payload, headers)
}

func (c *Client) post(path string, payload interface{}) (res *http.Response, err error) {
	return c.do("POST", path, payload, nil)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

// NewClient Return a Cloud Deploy client
func NewClient(endpoint string, username string, password string) *Client {
	return &Client{Endpoint: endpoint, Username: username, Password: password}
}
