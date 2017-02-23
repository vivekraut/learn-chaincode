/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}
	
	if function == "addTable" {
		// Deletes an entity from its state
		return t.addTable(stub, args)
	}
	
	if function == "getTable" {
		// Deletes an entity from its state
		return t.getTable(stub, args)
	}

	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) addTable(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	err := stub.CreateTable("Customer", []*shim.ColumnDefinition{
	&shim.ColumnDefinition{Name: "Customer_ID", Type: shim.ColumnDefinition_STRING, Key: true},
	&shim.ColumnDefinition{Name: "Customer_Name", Type: shim.ColumnDefinition_STRING, Key: false},
	&shim.ColumnDefinition{Name: "Customer_Gender", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	
	if err != nil {
		return nil, err
	}

	success1, err := stub.InsertRow("Customer", shim.Row{
	Columns: []*shim.Column{
	&shim.Column{Value: &shim.Column_String_{String_: "C1001"}},
	&shim.Column{Value: &shim.Column_String_{String_: "Vivek"}},
	&shim.Column{Value: &shim.Column_String_{String_: "Male"}},
	},
	})
	
	if !success1 {
		return nil, errors.New("Entity not found")
	}
	
	if err != nil {
		return nil, err
	}
	
	success2, err := stub.InsertRow("Customer", shim.Row{
	Columns: []*shim.Column{
	&shim.Column{Value: &shim.Column_String_{String_: "C1002"}},
	&shim.Column{Value: &shim.Column_String_{String_: "John"}},
	&shim.Column{Value: &shim.Column_String_{String_: "Male"}},
	},
	})
	
	if !success2 {
		return nil, errors.New("Entity not found")
	}
	if err != nil {
		return nil, err
	}
	
	success3, err := stub.InsertRow("Customer", shim.Row{
	Columns: []*shim.Column{
	&shim.Column{Value: &shim.Column_String_{String_: "C1003"}},
	&shim.Column{Value: &shim.Column_String_{String_: "Simone"}},
	&shim.Column{Value: &shim.Column_String_{String_: "Female"}},
	},
	})
	
	if !success3 {
		return nil, errors.New("Entity not found")
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) getTable(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "C1001"}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Customer", columns)
	if err != nil {
		return nil, fmt.Errorf("getRows operation failed. %s", err)
	}
	
	cust := row.Columns[1].GetBytes()
	fmt.Printf("Customer = %d\n", cust)
	//myLogger.Debugf(" customer is [% x]",  cust)
	
	var columns2 []shim.Column
	col2 := shim.Column{Value: &shim.Column_String_{String_: "Male"}}
	columns = append(columns2, col2)
	
	rowChannel, err := stub.GetRows("Customer", columns2)
	if err != nil {
		return nil, fmt.Errorf("getRows operation failed. %s", err)
	}
	var rows []shim.Row
		for {
			select {
			case row, ok := <-rowChannel:
				if !ok {
					rowChannel = nil
				} else {
					rows = append(rows, row)
				}
			}
			if rowChannel == nil {
				break
			}
		}
	
	jsonRows, err := json.Marshal(rows)
		if err != nil {
			return nil, fmt.Errorf("getRows operation failed. Error marshaling JSON: %s", err)
		}
	fmt.Printf("Query Response:%s\n", jsonRows)
	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
