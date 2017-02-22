package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"bytes"

	log "github.com/Sirupsen/logrus"
)

func addOne(x int) int {
	return x + 1
}

func CallFunction(args []interface{}, function interface{}) ([]interface{}, error) {
	functionTy := reflect.TypeOf(function)
	functionVal := reflect.ValueOf(function)

	fmt.Println("Function has", functionTy.NumIn(), "parameters")
	fmt.Println("Function has", functionTy.NumOut(), "results")
	fmt.Println("Function sig:", functionTy)

	if functionTy.NumIn() != len(args) {
		return nil, errors.New("Function has " + strconv.Itoa(functionTy.NumIn()) + " params, got " + strconv.Itoa(len(args)) + " arguments.")
	}

	// Set up argument values
	argVals := make([]reflect.Value, len(args))
	for i, arg := range args {
		paramTy := functionTy.In(i)
		fmt.Println(paramTy)

		// Look up the arugment in the args list
		argTy := reflect.TypeOf(arg)
		if paramTy != argTy {
			return nil, errors.New("Argument " + strconv.Itoa(i) + " type mismatch: got " + argTy.String() + " but needed " + paramTy.String())
		}

		argVals[i] = reflect.ValueOf(arg)
	}

	// Call the function
	resultVals := functionVal.Call(argVals)

	// Convert the results from values to go types
	results := make([]interface{}, len(resultVals))
	for i, resultVal := range resultVals {
		resultTy := resultVal.Type()
		resultKind := resultVal.Kind()
		fmt.Println(resultTy, resultKind)

		switch resultKind {
		case reflect.Int:
			results[i] = resultVal.Int()
		case reflect.Bool:
			results[i] = resultVal.Bool()
		case reflect.String:
			results[i] = resultVal.String()
		default:
			return nil, errors.New("Unhandled kind " + resultKind.String())
		}
	}

	return results, nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Expected one argument")
	}

	var args map[string]interface{}
	dec := json.NewDecoder(bytes.NewBuffer([]byte(os.Args[1])))
	dec.UseNumber()
	err := dec.Decode(&args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", args)

	// Convert to slice
	argSlice := make([]interface{}, 0)
	for i := 0; i < len(args); i++ {
		idxStr := strconv.Itoa(i)
		if arg, ok := args[idxStr]; ok {
			switch arg := arg.(type) {
			case json.Number: // try parsing as int, if fail, assume float
				i32, err := strconv.ParseInt(arg.String(), 10, 32)
				if err == nil {
					argSlice = append(argSlice, int(i32))
					continue
				}
				f64, err := strconv.ParseFloat(arg.String(), 64)
				if err != nil {
					log.Fatal(err)
				}
				argSlice = append(argSlice, f64)
			default:
				argSlice = append(argSlice, arg)
			}
		} else {
			log.Info("Didn't find argument", i)
			break
		}
	}

	results, err := CallFunction(argSlice, addOne)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", results)

	// x := args["x"]
	// switch x := x.(type) {
	// case string:
	// 	addOne(strconv.Atoi(x))
	// default:

	// }

	// strconv.Atoi(arg["x"])

	fmt.Printf("{ \"msg\": \"Hello from arbitrary Go program!\", \"args\": %s }\n", os.Args[1:])
}
