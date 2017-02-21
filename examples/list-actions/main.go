package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	openwhisk "github.com/c3sr/openwhisk-go"
	actions "github.com/c3sr/openwhisk-go/client/actions"
)

func main() {
	cli, err := openwhisk.NewBasicAuthClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	p := actions.NewGetActionByNameParams()
	p.SetNamespace("_")
	p.SetActionName("echo-go")

	ok, err := cli.Actions.GetActionByName(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*ok.Payload)
}
