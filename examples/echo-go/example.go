package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("invocation starting")

	fmt.Println("There are ", len(os.Args[1:]), "args")

	var arg map[string]interface{}
	json.Unmarshal([]byte(os.Args[1]), &arg)

	fmt.Println(arg)

	fmt.Printf("{ \"msg\": \"Hello from arbitrary Go program!\", \"args\": %s }\n", os.Args[1:])
}
