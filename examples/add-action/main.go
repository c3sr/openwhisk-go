package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	openwhisk "github.com/c3sr/openwhisk-go"
	actions "github.com/c3sr/openwhisk-go/swagger_client/actions"
	models "github.com/c3sr/openwhisk-go/swagger_models"
)

func main() {
	cli, err := openwhisk.NewBasicAuthClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	p := actions.NewUpdateActionParams()

	overwriteStr := "false"
	p.Overwrite = &overwriteStr
	p.SetActionName("example-action")
	p.SetNamespace("IBM-ILLINOIS-C3SR_dev")

	ap := &models.ActionPut{Version: "0.0.2", Publish: true}
	kindStr := models.ActionExecKindBlackbox

	ap.Exec = &models.ActionExec{Kind: &kindStr, Image: "c3sr/echo-go"}
	memLimit := int32(256)
	timeout := int32(60000)
	ap.Limits = &models.ActionLimits{Memory: &memLimit, Timeout: &timeout}
	p.SetAction(ap)

	ok, err := cli.Actions.UpdateAction(p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

}
