package openwhisk

import (
	"errors"

	apiclient "github.com/c3sr/openwhisk-go/swagger_client"
	models "github.com/c3sr/openwhisk-go/swagger_models"

	"os"

	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

func NewBasicAuthClientFromEnv() (*apiclient.OpenWhiskREST, error) {
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

	return NewBasicAuthClient(endpoint, username, password), nil

}

func NewBasicAuthClient(endpoint, username, password string) *apiclient.OpenWhiskREST {

	formats := strfmt.Default
	wr := httptransport.BasicAuth(username, password)
	transport := httptransport.New(endpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = wr
	cli := apiclient.New(transport, formats)
	return cli
}

func NewKeyValue(key, value string) *models.KeyValue {
	return &models.KeyValue{
		Key:   &key,
		Value: &value,
	}
}
