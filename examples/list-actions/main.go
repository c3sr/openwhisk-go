package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	openwhisk "github.com/c3sr/openwhisk-go"
	actions "github.com/c3sr/openwhisk-go/swagger_client/actions"
)

func main() {
	cli, err := openwhisk.NewBasicAuthClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	p := actions.NewGetAllActionsParams()
	p.SetNamespace("_")

	ok, err := cli.Actions.GetAllActions(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Actions:")
	for _, m := range ok.Payload {
		fmt.Println(*m.Name, *m.Namespace)
	}
}
