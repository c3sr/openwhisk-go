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

	trueStr := "true"
	// falseStr := "false"

	p := actions.NewInvokeActionParams().
		WithNamespace("_").
		WithActionName("echo-go")
		// WithResult(&falseStr)
	p.SetBlocking(&trueStr)

	kv := openwhisk.NewKeyValue("keyString", "valueString")
	p.SetPayload(kv)

	ok, accepted, err := cli.Actions.InvokeAction(p)
	if err != nil {
		log.Fatal(err)
	}
	if ok != nil {
		fmt.Println("ActionOK")
		activation := *ok.Payload
		fmt.Printf("%#v\n", activation)
	}
	if accepted != nil {
		fmt.Println("ActionAccepted")
		id := *accepted.Payload
		fmt.Printf("%#v\n", id)
	}

}
