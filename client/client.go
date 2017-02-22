package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	models "github.com/c3sr/openwhisk-go/models"
)

const (
	actionPath = "/api/v1/namespaces/_/actions/"
)

type Client struct {
	endpoint string
	username string
	password string
}

func NewClient(endpoint, username, password string) *Client {
	return &Client{endpoint: endpoint, username: username, password: password}
}

func NewClientFromEnv() (*Client, error) {
	endpoint := os.Getenv("OPENWHISK_ENDPOINT")
	if endpoint == "" {
		return nil, errors.New("Environment variable OPENWHISK_ENDPOINT was not set")
	}
	username := os.Getenv("OPENWHISK_USERNAME")
	if username == "" {
		return nil, errors.New("Environment variable OPENWHISK_USERNAME was not set")
	}
	password := os.Getenv("OPENWHISK_PASSWORD")
	if password == "" {
		return nil, errors.New("Environment variable OPENWHISK_PASSWORD was not set")
	}
	return NewClient(endpoint, username, password), nil
}

func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)

}

func (c *Client) NewBasicAuthRequest(method string, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	return req, nil
}

func (c *Client) GetActions() ([]models.Action, error) {
	req, err := c.NewBasicAuthRequest("GET", c.endpoint+actionPath, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := doReq(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	var actions models.Actions
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (c *Client) Invoke(i *models.ActionInvocation) ([]byte, error) {

	jsonBytes, err := json.Marshal(i.Params)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonBytes))

	req, err := c.NewBasicAuthRequest("POST", c.endpoint+actionPath+i.Name, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	// Set request parameters
	q := req.URL.Query()
	if i.Blocking {
		q.Add("blocking", "true")
	} else {
		q.Add("blocking", "false")
	}
	if i.Result {
		q.Add("result", "true")
	} else {
		q.Add("result", "false")
	}

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)

	bytes, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
