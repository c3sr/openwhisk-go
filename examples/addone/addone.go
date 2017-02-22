package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	jsoncall "github.com/c3sr/openwhisk-go/jsoncall"
)

func div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by 0!")
	} else {
		return x / y, nil
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Expected one argument")
	}

	argsJSON := []byte(os.Args[1])
	resultJSON, err := jsoncall.CallWithJSON(div, argsJSON)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resultJSON))

	fmt.Printf("{ \"msg\": \"Hello from arbitrary Go program!\", \"args\": %s }\n", os.Args[1:])
}
