package openwhisk

import (
	"errors"
	"fmt"
	"log"
	"os"

	client "github.com/c3sr/openwhisk-go/client"
	swclient "github.com/c3sr/openwhisk-go/swagger_client"
	swactions "github.com/c3sr/openwhisk-go/swagger_client/actions"
	swmodels "github.com/c3sr/openwhisk-go/swagger_models"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

type Client struct {
	swcli *swclient.OpenWhiskREST // client from swagger
	cli   *client.Client          // manual client for stuff that's broken in swagger
}

type overwriteActionOptions struct {
	version    string
	memLimitMb int32
	timeoutMs  int32
	publish    bool
	namespace  string
	overwrite  *string
}

type OverwriteActionOption func(o *overwriteActionOptions)

func defaultOverwriteActionOptions() *overwriteActionOptions {
	trueStr := "true"
	return &overwriteActionOptions{"0.0.1", 256, 60000, true, "_", &trueStr}
}

func NewClient(endpoint, username, password string) *Client {
	return &Client{newBasicAuthSwaggerClient(endpoint, username, password),
		client.NewClient(endpoint, username, password)}
}

func (c *Client) OverwriteAction(repo, tag, actionName string, options ...OverwriteActionOption) {

	// Apply any passed options
	opts := defaultOverwriteActionOptions()
	for _, o := range options {
		o(opts)
	}

	p := swactions.NewUpdateActionParams().
		WithOverwrite(opts.overwrite).
		WithActionName(actionName).
		WithNamespace(opts.namespace)
	// p.SetNamespace("IBM-ILLINOIS-C3SR_dev")

	ap := &swmodels.ActionPut{Version: opts.version, Publish: opts.publish}
	kindStr := swmodels.ActionExecKindBlackbox

	repoTag := repo
	if tag != "" {
		repoTag += ":" + tag
	}

	ap.Exec = &swmodels.ActionExec{Kind: &kindStr, Image: repoTag}
	memLimit := int32(256)
	timeout := int32(60000)
	ap.Limits = &swmodels.ActionLimits{Memory: &memLimit, Timeout: &timeout}
	p.SetAction(ap)

	ok, err := c.swcli.Actions.UpdateAction(p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func (c *Client) InvokeAction() {

}

func newBasicAuthSwaggerClientFromEnv() (*swclient.OpenWhiskREST, error) {
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

	return newBasicAuthSwaggerClient(endpoint, username, password), nil

}

func newBasicAuthSwaggerClient(endpoint, username, password string) *swclient.OpenWhiskREST {

	formats := strfmt.Default
	wr := httptransport.BasicAuth(username, password)
	transport := httptransport.New(endpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = wr
	cli := swclient.New(transport, formats)
	return cli
}

func NewKeyValue(key, value string) *swmodels.KeyValue {
	return &swmodels.KeyValue{
		Key:   &key,
		Value: &value,
	}
}
