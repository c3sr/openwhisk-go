package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cwpearson/openwhisk-go/model"
)

const (
	actionPath = "/api/v1/namespaces/_/actions/"
)

type Parameters map[string]string

type ActionInvocation struct {
	name     string
	params   Parameters
	blocking bool
	result   bool
}

type Client struct {
	endpoint string
	username string
	password string
}

func NewClient(endpoint, username, password string) *Client {
	return &Client{endpoint: "https://openwhisk.ng.bluemix.net", username: username, password: password}
}

func NewInvocation(actionName string, options ...func(*ActionInvocation)) *ActionInvocation {

	i := &ActionInvocation{name: actionName, params: Parameters{}}

	for _, option := range options {
		option(i)
	}

	return i
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

func (c *Client) GetActions() ([]model.Action, error) {
	req, err := c.NewBasicAuthRequest("GET", c.endpoint+actionPath, nil)
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

func (i *ActionInvocation) AddParameter(key string, value string) {
	i.params[key] = value
}

func Blocking(i *ActionInvocation) {
	i.blocking = true
}

func Nonblocking(i *ActionInvocation) {
	i.blocking = false
}

func ResultOnly(i *ActionInvocation) {
	i.result = true
}

func (c *Client) Invoke(i *ActionInvocation) ([]byte, error) {

	jsonBytes, err := json.Marshal(i.params)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonBytes))

	req, err := c.NewBasicAuthRequest("POST", c.endpoint+actionPath+i.name, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	// Set request parameters
	q := req.URL.Query()
	if i.blocking {
		q.Add("blocking", "true")
	} else {
		q.Add("blocking", "false")
	}
	if i.result {
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
