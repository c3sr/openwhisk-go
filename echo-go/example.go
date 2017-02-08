package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is an example log message from a Go program.")
	fmt.Printf("{ \"msg\": \"Hello from arbitrary Go program!\", \"args\": %s }\n", os.Args[1:])
}
