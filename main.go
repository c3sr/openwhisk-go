package main

import (
	"fmt"
	"log"
	"os"

	apiclient "github.com/cwpearson/openwhisk-go/client"
)

func main() {

	// create the API client
	client := apiclient.NewClient(os.Getenv("OPENWHISK_ENDPOINT"),
		os.Getenv("OPENWHISK_USERNAME"),
		os.Getenv("OPENWHISK_PASSWORD"))

	fmt.Println(client)

	actions, err := client.GetActions()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(actions)

	invocation := apiclient.NewInvocation("echo-go", apiclient.Blocking)
	resp, err := client.Invoke(invocation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
	// make the authenticated request to get all items
	// bearerTokenAuth := httptransport.BearerToken(os.Getenv("API_ACCESS_TOKEN"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))
	// resp, err := client.Operations.All(operations.AllParams{}, bearerTokenAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, basicAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyQueryAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyHeaderAuth)

	// fmt.Printf("%#v\n", String(entities)
}
