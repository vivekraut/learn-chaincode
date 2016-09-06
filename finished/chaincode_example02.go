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
	"bytes"

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


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, value string	
    var err error
    fmt.Println("Storing the parameters in hyperledger fabric...")

    /*if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }*/
	
	if(len(args)%2 != 0) {
		  fmt.Printf("Incorrect number of arguments. One of the keys or values is missing.")
		  fmt.Println("")
				  
    }else{
	     for i := 0; i < len(args); i++ {
	     if(i%2 == 0){
		     if args[i] != "" {
                  fmt.Printf("Key: %s", args[i])
				  fmt.Println("")
				  key = args[i]    
				  i++
             }
		     if(i!=len(args)) {
			      fmt.Printf("Value: %s", args[i])
			      fmt.Println("")
				  value = args[i]
			 }
		     err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
			 if err != nil {
				return nil, err
			 }
		 }
    }
	}

	/*
    key = args[0]                            //rename for fun
    value = args[1]
    err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
	*/
		
    return nil, nil
}


// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}
	if function == "write" {
		fmt.Println("Calling write()")
        return t.write(stub, args)
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
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	
	//Store state for transactions
	var transacted, historyval string
	transacted = "T_"+args[0]+"|"+args[1]
	Tvalbytes, err := stub.GetState(transacted)
	if err != nil {
		return nil, errors.New("Failed to get transacted state")
	}
	if Tvalbytes == nil {
		historyval = args[2]
	}else{
		historyval = string(Tvalbytes)
		historyval = historyval+","+args[2]		
	}	
	err = stub.PutState(transacted, []byte(historyval))
	if err != nil {
		return nil, err
	}
	
	
	//Store state for sponsor transactions
	var s_transactions, s_history string
	s_transactions = "T_"+args[0]
	Svalbytes, err := stub.GetState(s_transactions)
	if err != nil {
		return nil, errors.New("Failed to get sponsor transacted state")
	}
	if Svalbytes == nil {
		s_history = args[1]+"|"+args[2]
	}else{
		s_history = string(Svalbytes)
		s_history = s_history+","+args[1]+"|"+args[2]	
	}	
	err = stub.PutState(s_transactions, []byte(s_history))
	if err != nil {
		return nil, err
	}
	
	
	//Store state for Idea transactions
	var i_transactions, i_history string
	i_transactions = "T_"+args[1]
	Ivalbytes, err := stub.GetState(i_transactions)
	if err != nil {
		return nil, errors.New("Failed to get idea transacted state")
	}
	if Ivalbytes == nil {
		i_history = args[0]+"|"+args[2]
	}else{
		i_history = string(Ivalbytes)
		i_history = i_history+","+args[0]+"|"+args[2]	
	}	
	err = stub.PutState(i_transactions, []byte(i_history))
	if err != nil {
		return nil, err
	}
	
	
	

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
	/*if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}*/
	
	if function == "queryAll" {
		fmt.Println("Calling QueryAll()")
        return t.queryAll(stub, args)
    }
	
	if function == "queryTransact" {
		fmt.Println("Calling QueryTransact()")
        return t.queryTransact(stub, args)
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
		Avalbytes = []byte("0")
		//jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		//return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}





// Query callback representing the query of a chaincode
func (t *SimpleChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		
	//var A string // Entities
	//var err error

	/*if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}*/
    var RetValue []byte
	var buffer bytes.Buffer 
    var jsonRespString string    
		for i := 0; i < len(args); i++ {	     
		    Avalbytes, err := stub.GetState(args[i])
			if err != nil {
				jsonResp := "{\"Error\":\"Failed to get state for " + args[i] + "\"}"
				return nil, errors.New(jsonResp)
			}

			if Avalbytes == nil {
				Avalbytes = []byte("0")
				//jsonResp := "{\"Error\":\"Nil amount for " + args[i] + "\"}"
				//return nil, errors.New(jsonResp)
			}
			if(i!=len(args)-1) {
			   jsonRespString =  string(Avalbytes)+","
			}else{
			   jsonRespString =  string(Avalbytes)
			}
			buffer.WriteString(jsonRespString)			
			RetValue = []byte(buffer.String())			
			
		}
		jsonResp := "{"+buffer.String()+"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return RetValue, nil
		
	
	/*
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
	*/
}





// Query callback representing the query of a chaincode
func (t *SimpleChaincode) queryTransact(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		
	//var A string // Entities
	//var err error

	/*if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}*/
    var RetValue []byte
	var buffer bytes.Buffer 
    var jsonRespString, queryparam string  
	queryparam = "T_"+args[0]
			     
		Tvalbytes, err := stub.GetState(queryparam)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
			return nil, errors.New(jsonResp)
		}

		if Tvalbytes == nil {
			Tvalbytes = []byte("0")
			//jsonResp := "{\"Error\":\"Nil amount for " + args[i] + "\"}"
			//return nil, errors.New(jsonResp)
		}
		
		jsonRespString =  string(Tvalbytes)		
		buffer.WriteString(jsonRespString)			
		RetValue = []byte(buffer.String())
		jsonResp := "{"+buffer.String()+"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return RetValue, nil
		
	
	/*
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
	*/
}




func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
