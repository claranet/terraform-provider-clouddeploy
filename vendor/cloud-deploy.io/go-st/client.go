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
		//var eo *errorObject
		//var getErr error
		//if eo, getErr = c.getErrorFromResponse(resp); getErr != nil {
		//	return resp, fmt.Errorf("Response did not contain formatted error: %s. HTTP response code: %v. Raw response: %+v", getErr, resp.StatusCode, resp)
		//}
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

// NewClient Return a Ghost client
func NewClient(endpoint string, username string, password string) *Client {
	return &Client{Endpoint: endpoint, Username: username, Password: password}
}

// Returns all apps
func (c *Client) GetApps() (apps Apps, err error) {
	res, err := c.get("/apps")
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&apps)
	}
	return
}

// Create a new app
func (c *Client) CreateApp(app App) (metadata EveItemMetadata, err error) {
	res, err := c.post("/apps", app)
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&metadata)
	}
	return
}

// Return the requested app
func (c *Client) GetApp(id string) (app App, err error) {
	res, err := c.get("/apps/" + id)
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&app)
	}
	return
}

// Update an existing app
func (c *Client) UpdateApp(app *App, id string, etag string) (metadata EveItemMetadata, err error) {
	res, err := c.patch("/apps/"+id, app, map[string]string{"If-Match": etag})
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&metadata)
	}
	return
}

// Delete an existing app
func (c *Client) DeleteApp(id string, etag string) (err error) {
	_, err = c.delete("/apps/"+id, map[string]string{"If-Match": etag})
	return
}
