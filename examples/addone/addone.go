package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func addOne(x int) int {
	return x + 1
}

func CallFunction(args map[string]interface{}, function interface{}) (interface{}, error) {
	functionTy := reflect.TypeOf(function)

	fmt.Println("Function has", functionTy.NumIn(), "arguments")

	fmt.Println(functionTy)
	for i := 0; i < functionTy.NumIn(); i++ {
		paramTy := functionTy.In(i)
		fmt.Println(paramTy)

		// Look up the arugment in the args list
		providedTy := reflect.TypeOf(args[strconv.Itoa(i)])
		if paramTy != providedTy {
			fmt.Println("Argument type mismatch: got", providedTy, "and expected", paramTy)
		}

		// Call the function
		function.Call()

	}

	fmt.Println("Function has", functionTy.NumOut(), "results")

	for o := 0; o < functionTy.NumOut(); o++ {
		resultTy := functionTy.Out(o)
		fmt.Println(resultTy)
	}

	return nil, nil
}

func main() {

	var args map[string]interface{}
	json.Unmarshal([]byte(os.Args[1]), &args)

	fmt.Println(args)

	_, _ = CallFunction(args, addOne)

	// x := args["x"]
	// switch x := x.(type) {
	// case string:
	// 	addOne(strconv.Atoi(x))
	// default:

	// }

	// strconv.Atoi(arg["x"])

	fmt.Printf("{ \"msg\": \"Hello from arbitrary Go program!\", \"args\": %s }\n", os.Args[1:])
}
