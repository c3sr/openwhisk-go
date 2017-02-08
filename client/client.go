package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	model "github.com/cwpearson/openwhisk-go/model"
)

const (
	actionPath = "/api/v1/namespaces/_/actions"
)

type Client struct {
	endpoint string
	username string
	password string
}

func New(endpoint, username, password string) *Client {
	return &Client{endpoint: "https://openwhisk.ng.bluemix.net", username: username, password: password}
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

func (c *Client) NewRequestWithBasicAuth(method string, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	return req, nil
}

func (c *Client) GetActions() ([]model.Action, error) {
	req, err := c.NewRequestWithBasicAuth("GET", c.endpoint+actionPath, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := doReq(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	var actions model.Actions
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (c *Client) ListEntities() ([]byte, error) {
	// req, err := http.NewRequest("GET", c.endpoint+"/api/v1/namespaces/whisk.system/packages", nil)
	req, err := http.NewRequest("GET", c.endpoint+"/api/v1/namespaces/_/actions", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)

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
