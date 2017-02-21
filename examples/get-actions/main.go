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

	action := *ok.Payload

	fmt.Println("Annotations:")
	for _, kv := range action.Annotations {
		fmt.Printf("  %s: %#v\n", *kv.Key, kv.Value)
	}

	fmt.Printf("namespace: %s\n", *action.Namespace)
	fmt.Printf("name: %s\n", *action.Name)
	fmt.Printf("mem limit: %d\n", *(*action.Limits).Memory)
	fmt.Printf("time limit: %d\n", *(*action.Limits).Timeout)
	fmt.Printf("exec: %#v\n", *action.Exec)
	fmt.Printf("exec kind: %s\n", *(*action.Exec).Kind)
}
