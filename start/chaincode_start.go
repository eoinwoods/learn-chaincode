/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("EW query is running " + function)

	fmt.Println("EW function='" + function + "' equality=" + strconv.FormatBool(function == "read"))
	// Handle different functions
	if function == "read" {
		fmt.Println("EW query is in 'read' branch")
		return t.read(stub, args)
	} else if function == "dummy_query" { //read a variable
		fmt.Println("EW hi there " + function) //error
		return nil, nil
	}
	fmt.Println("EW query did not find func: " + function) //error

	return nil, errors.New("EW Received unknown function query: " + function)
}

///
/// Helper methods on SimpleChainCode
///
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error

	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Wrong number of arguments - expected key and value strings")
	}

	key = args[0]
	value = args[1]
	err = stub.PutState(key, []byte(value))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResult string
	var err error
	var valAsBytes []byte

	if len(args) != 1 {
		return nil, errors.New("Wrong number of arguments, expected key to search for")
	}

	key = args[0]
	valAsBytes, err = stub.GetState(key)
	if err != nil {
		jsonResult = "{\"Error\":\"Failed to find state for key '" + key + "'\"}"
		return nil, errors.New(jsonResult)
	}
	return valAsBytes, nil
}
