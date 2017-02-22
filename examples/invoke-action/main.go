package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	openwhisk "github.com/c3sr/openwhisk-go"
	apiclient "github.com/c3sr/openwhisk-go/client"
	models "github.com/c3sr/openwhisk-go/models"
	swagger_actions "github.com/c3sr/openwhisk-go/swagger_client/actions"
)

func main() {
	c, err := apiclient.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ai := models.NewActionInvocation("echo-go", models.Blocking(true), models.ResultOnly(false))

	ai.AddParameter("key1", "value1")

	bytes, err := c.Invoke(ai)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", string(bytes))
}

func brokenSwaggerExample() {
	cli, err := openwhisk.NewBasicAuthClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	trueStr := "true"
	// falseStr := "false"

	p := swagger_actions.NewInvokeActionParams().
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
