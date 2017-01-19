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
	"log"
	//"github.com/hyperledger/fabric/accesscontrol/crypto/attr"
	//"github.com/hyperledger/fabric/accesscontrol/impl"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/errors"
	"encoding/json"	
	//"github.com/op/go-logging"
)

//var myLogger = logging.MustGetLogger("accum_share")
var myLogger = shim.NewLogger("CLDChaincode")

const   MEDICAL      =  "medical"
const   PHARMACY   =  "pharmacy"
const   DENTAL =  "dental"


type V5C_Holder struct {
	V5Cs 	[]string `json:"v5cs"`
}

type User_and_eCert struct {
	Identity string `json:"identity"`
	eCert string `json:"ecert"`
}



type AccumShare struct {
	Claims struct {
		PolicyID string `json:"PolicyID"`
		SubscriberID string `json:"SubscriberID"`
		PolicyStartDate string `json:"PolicyStartDate"`
		PolicyEndDate string `json:"PolicyEndDate"`
		PolicyType string `json:"PolicyType"`
		DeductibleBalance string `json:"DeductibleBalance"`
		OOPBalance string `json:"OOPBalance"`
		Claim struct {
			ClaimID string `json:"ClaimID"`
			MemberID string `json:"MemberID"`
			CreateDTTM string `json:"CreateDTTM"`
			LastUpdateDTTM string `json:"LastUpdateDTTM"`
			Transaction struct {
				TransactionID string `json:"TransactionID"`
				Accumulator struct {
					Type string `json:"Type"`
					Amount string `json:"Amount"`
					UoM string `json:"UoM"`
				} `json:"Accumulator"`
				Participant string `json:"Participant"`
				TotalTransactionAmount string `json:"TotalTransactionAmount"`
				UoM string `json:"UoM"`
				Overage string `json:"Overage"`
			} `json:"Transaction"`
			TotalClaimAmount string `json:"TotalClaimAmount"`
			UoM string `json:"UoM"`
		} `json:"Claim"`
	} `json:"Claims"`
}


// SimpleChaincode example simple Chaincode implementation
type AccumShareChaincode struct {
}

func (t *AccumShareChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	
	
	//var SubscriberID, PolicyID, PolicyStartDate, PolicyEndDate, PolicyType, DeductibleBalance, OOPBalance string    // Entities
	//var err error

	/*if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}*/
	
	
	// Initialize the chaincode
	function, args := stub.GetFunctionAndParameters()
	if function == "init" {
		// Deletes an entity from its state
		fmt.Println("Inside Init-------------------------")
	}
	var SubscriberIDValue, PolicyIDValue, PolicyStartDateValue, PolicyEndDateValue, PolicyTypeValue, DeductibleBalanceValue, OOPBalanceValue string
	
	//SubscriberID = args[0]	
	SubscriberIDValue = args[1]
	
	//PolicyID = args[2]
	PolicyIDValue = args[3]
	
	//PolicyStartDate = args[4]
	PolicyStartDateValue = args[5]
	
	//PolicyEndDate = args[6]
	PolicyEndDateValue = args[7]
	
	//PolicyType = args[8]
	PolicyTypeValue = args[9]
	
	//DeductibleBalance = args[10]
	DeductibleBalanceValue = args[11]
	
	//OOPBalance = args[12]
	OOPBalanceValue = args[13]
	
	
	
	jsonResponse := `{   "Claims": {      "PolicyID": "",      "SubscriberID": "",      "PolicyStartDate": "",      "PolicyEndDate": "",      "PolicyType": "",            "DeductibleBalance":"",      "OOPBalance":"",    	  "BalanceUoM":"", 	 	        "Claim": {         "ClaimID": "",         "MemberID": "",         "CreateDTTM": "",         "LastUpdateDTTM": "",         "Transaction": {            "TransactionID": "",            "Accumulator": {               "Type": "",                              "Amount": "",               "UoM": ""            },   "Overage":"",         "Participant": "",            "TotalTransactionAmount": "",            "UoM": ""         },         "TotalClaimAmount": "",         "UoM": ""      }   }}`
	
	res := &AccumShare{}
	
	err := json.Unmarshal([]byte(jsonResponse), res)
        if(err!=nil) {
            log.Fatal(err)
        }
	
	res.Claims.SubscriberID = SubscriberIDValue
	res.Claims.PolicyID = PolicyIDValue
	res.Claims.PolicyStartDate = PolicyStartDateValue
	res.Claims.PolicyEndDate = PolicyEndDateValue
	res.Claims.PolicyType = PolicyTypeValue	
	res.Claims.DeductibleBalance = DeductibleBalanceValue
	res.Claims.OOPBalance = OOPBalanceValue
	
	body, err := json.Marshal(res)
	if err != nil {
           panic(err)
        }
        fmt.Println(string(body))
	
	err = stub.PutState(SubscriberIDValue, []byte(string(body)))
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}


/*
func (t *AccumShareChaincode) get_ecert(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {

	ecert, err := stub.GetState(name)

	if err != nil { return nil, errors.New("Couldn't retrieve ecert for user " + name) }

	return ecert, nil
}

func (t *AccumShareChaincode) add_ecert(stub shim.ChaincodeStubInterface, name string, ecert string) ([]byte, error) {


	err := stub.PutState(name, []byte(ecert))

	if err == nil {
		return nil, errors.New("Error storing eCert for user " + name + " identity: " + ecert)
	}

	return nil, nil

}

func (t *AccumShareChaincode) get_username(stub shim.ChaincodeStubInterface) (string, error) {

    	username, err := impl.NewAccessControlShim(stub).ReadCertAttribute("username");
	if err != nil { return "", errors.New("Couldn't get attribute 'username'. Error: " + err.Error()) }
	return string(username), nil
}

func (t *AccumShareChaincode) check_affiliation(stub shim.ChaincodeStubInterface) (string, error) {
    affiliation, err := impl.NewAccessControlShim(stub).ReadCertAttribute("role");
	if err != nil { return "", errors.New("Couldn't get attribute 'role'. Error: " + err.Error()) }
	return string(affiliation), nil

}

func (t *AccumShareChaincode) get_caller_data(stub shim.ChaincodeStubInterface) (string, string, error){

	user, err := t.get_username(stub)

    // if err != nil { return "", "", err }

	// ecert, err := t.get_ecert(stub, user);

    // if err != nil { return "", "", err }

	affiliation, err := t.check_affiliation(stub);

    if err != nil { return "", "", err }

	return user, affiliation, nil
}
*/
func (t *AccumShareChaincode) processClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	
	
	var err error
	var DeductibleLimit int
	var SubscriberIDValue, ClaimIDValue, TransactionIDValue, TransactionAmountValue, UoMValue, CreateDTTMValue, LastUpdateDTTMValue, AccumTypeValue, ParticipantValue string
	
	SubscriberIDValue = args[1]
	ClaimIDValue = args[3]
	TransactionIDValue = args[5]
	TransactionAmountValue = args[7]
	UoMValue = args[9]
	CreateDTTMValue = args[11]
	LastUpdateDTTMValue = args[13]
	AccumTypeValue = args[15]
	ParticipantValue = args[17]
	
	
	DeductibleLimit = 500
	/*SubscriberAccums, err := t.query(stub, SubscriberIDValue)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	fmt.Printf("SubscriberAccums = %d\n", SubscriberAccums)	
	*/
	
	SubscriberAccums, err := stub.GetState(SubscriberIDValue)
	fmt.Printf("%v\n",SubscriberAccums)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + SubscriberIDValue + "\"}"
		return nil, errors.New(jsonResp)
	}

	if SubscriberAccums == nil {
		SubscriberAccums = []byte("No Records Found")
		//jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		//return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + SubscriberIDValue + "\",\"Value\":\"" + string(SubscriberAccums) + "\"}"
	fmt.Printf("%v\n",jsonResp)
	
	res := &AccumShare{}
	err = json.Unmarshal([]byte(SubscriberAccums), res)
        if(err!=nil) {
            log.Fatal(err)
        }

    	fmt.Printf("%v\n",res)
    	fmt.Printf("\tPolicy ID: %s\n",res.Claims.PolicyID)
    	fmt.Printf("\tSubscriber ID: %s\n",res.Claims.SubscriberID)
	fmt.Printf("\t Deductible Balance: %s\n",res.Claims.DeductibleBalance)
	
	var DedBalance, TransAmount, Overage int
	Overage = 0
	DedBalance, err = strconv.Atoi(res.Claims.DeductibleBalance)
	TransAmount, err = strconv.Atoi(TransactionAmountValue)
	
	
	if(AccumTypeValue == "IIDED"){
		if((TransAmount + DedBalance) <= DeductibleLimit){
			DedBalance = DedBalance + TransAmount	
		
		}else{
			DedBalance = DeductibleLimit
			Overage = TransAmount + DedBalance - DeductibleLimit
		}		
	}
	
	res.Claims.DeductibleBalance = strconv.Itoa(DedBalance)
	res.Claims.Claim.ClaimID = ClaimIDValue
	res.Claims.Claim.Transaction.TransactionID = TransactionIDValue
	res.Claims.Claim.Transaction.Accumulator.Type = AccumTypeValue
	res.Claims.Claim.Transaction.Accumulator.Amount = strconv.Itoa(TransAmount - Overage)
	res.Claims.Claim.Transaction.Accumulator.UoM = UoMValue
	res.Claims.Claim.Transaction.Overage = strconv.Itoa(Overage)
	res.Claims.Claim.Transaction.Participant = ParticipantValue
	res.Claims.Claim.Transaction.TotalTransactionAmount = TransactionAmountValue
	res.Claims.Claim.Transaction.UoM = UoMValue
	res.Claims.Claim.TotalClaimAmount = TransactionAmountValue
	res.Claims.Claim.UoM = UoMValue
	res.Claims.Claim.CreateDTTM = CreateDTTMValue
	res.Claims.Claim.LastUpdateDTTM = LastUpdateDTTMValue
	res.Claims.Claim.MemberID = SubscriberIDValue
	
	updatedBody, err := json.Marshal(res)
	if err != nil {
        	panic(err)
    	}
    	fmt.Println(string(updatedBody))
	err = stub.PutState(SubscriberIDValue, []byte(string(updatedBody)))
	if err != nil {
		return nil, err
	}
	
	 return nil, nil

}


func (t *AccumShareChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, value string	
    //var err error
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
			 
			 //check if the state exists. If not initialize the state
			Avalbytes, err := stub.GetState(key)
			if err != nil {
				jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
				return nil, errors.New(jsonResp)
			}
		
			if Avalbytes == nil {
				Avalbytes = []byte("0")
				//err = stub.PutState(key, Avalbytes)
				
				//jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
				//return nil, errors.New(jsonResp)
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


func (t *AccumShareChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	function, args := stub.GetFunctionAndParameters()
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}
	if function == "write" {
		fmt.Println("Calling write()")
        return t.write(stub, args)
        }
	if function == "processClaim" {
		fmt.Println("Calling processClaim()")
        return t.processClaim(stub, args)
        }	
	
	return nil, nil
}

// Deletes an entity from state
func (t *AccumShareChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
func (t *AccumShareChaincode) Query(stub shim.ChaincodeStubInterface) ([]byte, error) {
	/*if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}*/
	var err error
	function, args := stub.GetFunctionAndParameters()
	/*
	caller, caller_affiliation, err := t.get_caller_data(stub)
	if err != nil { fmt.Printf("QUERY: Error retrieving caller details", err); 
		       return nil, errors.New("QUERY: Error retrieving caller details: "+err.Error()) 
		      }
	
    	myLogger.Debug("function: ", function)
    	myLogger.Debug("caller: ", caller)
   	myLogger.Debug("affiliation: ", caller_affiliation)
	*/
	
	if function == "queryAll" {
		fmt.Println("Calling QueryAll()")
        return t.queryAll(stub, args)
    }
	
	if function == "queryTransact" {
		fmt.Println("Calling QueryTransact()")
        return t.queryTransact(stub, args)
    }
	
	var A string // Entities
	

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
func (t *AccumShareChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		
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
func (t *AccumShareChaincode) queryTransact(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		
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
	myLogger.Debugf("***********Accum Share*************")
	err := shim.Start(new(AccumShareChaincode))
	if err != nil {
		//err = errors.ErrorWithCallStack(errors.Logging, errors.LoggingUnknownError, err.Error())
		fmt.Printf("Error starting Simple chaincode: %s", err)
		//return nil, err
	}
}
