
package main

import (
	"errors"
	"fmt"
	//"strconv"
	//"time"
	//"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

type RecruitmentChainCode struct {
}


type Demand struct{
	RRDNo string `json:"RRDNo"`
	Role string `json:"Role"`
	RoleDescription string `json:"RoleDescription"`
	YearsOfExperience string `json:"YearsOfExperience"`
	Location string `json:"Location"`
	Priority string `json:"Priority"`
	DemandStatus string `json:"DemandStatus"`
	Certification []Certification `json:"Certification"`
	JobDescription []JobDescription `json:"JobDescription"`
	MustHave []MustHave `json:"MustHave"`
	GoodToHave []GoodToHave `json:"GoodToHave"`
	SalaryRange []SalaryRange `json:"SalaryRange"`	
}

type Certification struct{
	Name []string `json:"Name"`		
}

type JobDescription struct{
	Description []string `json:"Description"`		
}

type MustHave struct{
	Skill []string `json:"Skill"`
}

type GoodToHave struct{
	Skill []string `json:"Skill"`
}

type SalaryRange struct{
	Min string `json:"Min"`
	Max string `json:"Max"`
	Currency string `json:"Currency"`
	
}

type Status struct{
	RRDNo string `json:"RRDNo"`
	Status string `json:"Status"`
}

func (self *RecruitmentChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("In Init start ")	
	fmt.Println("Nothing to be initialized ")
	return nil, nil
}

func (self *RecruitmentChainCode) Invoke(stub shim.ChaincodeStubInterface,
	function string, args []string) ([]byte, error) {
	fmt.Println("In Invoke with function  " + function)
	
	if function == "AddDemand" {
		fmt.Println("invoking AddDemand " + function)
		testBytes,err := AddDemand(args[0],stub)
		if err != nil {
			fmt.Println("Error performing AddDemand ")
			return nil, err
		}
		fmt.Println("Processed AddDemand successfully. ")
		return testBytes, nil
	}
	
	
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

func (self *RecruitmentChainCode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	fmt.Println("In Query with function " + function)
	//bytes, err:= query.Query(stub, function,args)
	//if err != nil {
		//fmt.Println("Error retrieving function  ")
		//return nil, err
	//}
	
	bytes, err:= QueryDetails(stub, function,args)
	if err != nil {
		fmt.Println("Error retrieving function  ")
		return nil, err
	}
	return bytes,nil
	
}

func QueryDetails(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "GetDemand" {
		fmt.Println("Invoking GetDemand " + function)
		var demand Demand
		demand,err := GetDemand(args[0], stub)
		if err != nil {
			fmt.Println("Error retrieving  the Demand")
			return nil, errors.New("Error retrieving  Demand")
		}
		fmt.Println("All success, returning demand")
		return json.Marshal(demand)
	}
	return nil, errors.New("Received unknown query function name")

}

func GetDemand(RRDNo string, stub shim.ChaincodeStubInterface)(Demand, error) {
	fmt.Println("In query.GetDemand start ")

	key := RRDNo
	var demand Demand
	demandBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving demand" , RRDNo)
		return demand, errors.New("Error retrieving demand" + RRDNo)
	}
	err = json.Unmarshal(demandBytes, &demand)
	fmt.Println("Demand   : " , demand);
	fmt.Println("In query.GetDemand end ")
	return demand, nil
}


func GetStatus(RRDNo string, stub shim.ChaincodeStubInterface)(Status, error) {
	fmt.Println("In query.GetStatus start ")
	key := RRDNo
	var status Status	
	demand,err := GetDemand(key, stub)
	if err != nil {
		fmt.Println("Error retrieving Status" , RRDNo)
		return status, errors.New("Error retrieving Status" + RRDNo)
	}
	status.RRDNo = key
	status.Status = demand.DemandStatus		
	fmt.Println("Demand Status   : " , status.Status);
	fmt.Println("In query.GetStatus end ")
	return status, nil
}


func AddDemand(demandJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.AddUser start ")
	//smartMeterId :=args[0]
	//userType 	:=args[1]
	
	//var user User
	res := &Demand{}
	//user := &User{}
	err := json.Unmarshal([]byte(demandJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal Demand ")
	}	
	fmt.Println("RRD Number : ",res.RRDNo)	
	
	body, err := json.Marshal(res)
	if err != nil {
        panic(err)
    }

    fmt.Println(string(body))	
	err = stub.PutState(res.RRDNo, []byte(string(body)))
	if err != nil {
		fmt.Println("Failed to create Demand ")
	}	
	err = stub.PutState(res.RRDNo + "_Status", []byte(res.DemandStatus))
	if err != nil {
		fmt.Println("Failed to set RRD Status ")
	}	
	fmt.Println("Created Demand  with Key : "+ res.RRDNo)
	fmt.Println("In initialize.AddDemand end ")
	return nil,nil		
}


func StatusUpdate(demandStatusJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.StatusUpdate start ")
		
	res := &Status{}	
	err := json.Unmarshal([]byte(demandStatusJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal Status ")
	}	
	fmt.Println("Status of Demand   : ",res.Status)	
	
	demandDet,err := GetDemand(res.RRDNo, stub)
	demandDet.DemandStatus = res.Status	
	
	demandStatusUpdate, err := json.Marshal(demandDet)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(demandStatusUpdate))	
	err = stub.PutState(res.RRDNo, []byte(string(demandStatusUpdate)))
	if err != nil {
		fmt.Println("Failed to update Demand Status ")
	}
	fmt.Println("Updated Demand Status  with Key : "+ res.RRDNo)
	fmt.Println("In initialize.StatusUpdate end ")
	return nil,nil			
}


func main() {
	err := shim.Start(new(RecruitmentChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}

}
