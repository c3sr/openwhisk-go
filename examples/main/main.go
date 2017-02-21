package main

import (
	"fmt"
	"log"

	apiclient "github.com/c3sr/openwhisk-go/client"
	actions "github.com/c3sr/openwhisk-go/client/actions"

	"os"

	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

func main() {

	formats := strfmt.Default
	endpoint := os.Getenv("OPENWHISK_ENDPOINT")

	wr := httptransport.BasicAuth(os.Getenv("OPENWHISK_USERNAME"), os.Getenv("OPENWHISK_PASSWORD"))
	transport := httptransport.New(endpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = wr

	cli := apiclient.New(transport, formats)

	p := actions.NewGetAllActionsParams()
	p.SetNamespace("_")

	ok, err := cli.Actions.GetAllActions(p)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal(err)
	}

	for _, m := range ok.Payload {
		fmt.Println(*m.Name)
	}
}
