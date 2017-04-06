
package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/energy/contracts/data"
	//"github.com/energy/contracts/query"
	"encoding/json"
)

type EnergyTradingChainCode struct {
}

type User struct{
	UserID string `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	SmartMeterID string `json:"SmartMeterID"`
	UserType string `json:"UserType"`
	BuyPrice string `json:"BuyPrice"`
	SellPrice string `json:"SellPrice"`
	EnergyConsumed string `json:"EnergyConsumed"`
	EnergyProduced string `json:"EnergyProduced"`
	EnergyAccountBalance string `json:"EnergyAccountBalance"`
}

type GridPrice struct{
	UserID string `json:"UserID"`
	Date string `json:"Date"`
	Price string `json:"Price"`		
}


type Proposal struct{
	UserID string `json:"UserID"`
	ProposalID string `json:"ProposalID"`
	Date string `json:"Date"`
	Price string `json:"Price"`
	EnergyProposed string `json:"EnergyProposed"`
	Status string `json:"Status"`
	EnergySigned string `json:"EnergySigned"`
	EnergyRemaining string `json:"EnergyRemaining"`
	//Consumer []Consumer `json:"Consumer"`
	//Contract []Contract `json:"Contract"`
}

type Contract struct{
	UserID string `json:"UserID"`
	ProposalID string `json:"ProposalID"`
	ContractID string `json:"ContractID"`
	Date string `json:"Date"`	
	EnergySigned string `json:"EnergySigned"`
	EnergyConsumed string `json:"EnergyConsumed"`
	Status string `json:"Status"`
	Price string `json:"Price"`
	BatteryBuyPrice string `json:"BatteryBuyPrice"`
	BatterySellPrice string `json:"BatterySellPrice"`
	GridPrice string `json:"GridPrice"`
	PlatformComission string `json:"PlatformComission"`
	Producer string `json:"Producer"`	
	Consumer string `json:"Consumer"`	
	Battery string `json:"Battery"`	
	Grid string `json:"Grid"`
	Platform string `json:"Platform"`
	
}

type Meter struct{
	EnergyReadingId string `json:"EnergyReadingId"`	
	Date string `json:"Date"`	
	SmartMeterId string `json:"SmartMeterId"`
	EnergyAmount string `json:"EnergyAmount"`
	EnergyUnit string `json:"EnergyUnit"`	
	
}


type Balance struct{
	UserID string `json:"UserID"`	
	Balance string `json:"Balance"`		
}


func (self *EnergyTradingChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("In Init start ")
	
	var UserID, SmartMeterID, UserType string 
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 - UserID, SmartMeterID and UserType. UserType should be Prosumer, Battery or Grid")
	}
	
	UserID = args[0]
	SmartMeterID = args[1]
	UserType = args[2]
	
	res := &User{}
	res.UserID = UserID
	res.SmartMeterID = SmartMeterID
	res.UserType = UserType
	
	body, err := json.Marshal(res)
	if err != nil {
        panic(err)
    }
    fmt.Println(string(body))		
	
	if function == "InitializeUser" {
		userBytes, err := AddUser(string(body),stub)
		if err != nil {
			fmt.Println("Error receiving  the User Details")
			return nil, err
		}
		fmt.Println("Initialization of User complete")
		return userBytes, nil
	}
	fmt.Println("Initialization No functions found ")
	return nil, nil
}


func (self *EnergyTradingChainCode) Invoke(stub shim.ChaincodeStubInterface,
	function string, args []string) ([]byte, error) {
	fmt.Println("In Invoke with function  " + function)
	
	if function == "AddUser" {
		fmt.Println("invoking AddUser " + function)
		testBytes,err := AddUser(args[0],stub)
		if err != nil {
			fmt.Println("Error performing AddUser ")
			return nil, err
		}
		fmt.Println("Processed AddUser successfully. ")
		return testBytes, nil
	}
	
	if function == "SetGridPrice" {
		fmt.Println("invoking SetGridPrice " + function)
		testBytes,err := SetGridPrice(args[0],stub)
		if err != nil {
			fmt.Println("Error performing SetGridPrice ")
			return nil, err
		}
		fmt.Println("Processed SetGridPrice successfully. ")
		return testBytes, nil
	}
	
	if function == "SetPlatformCharge" {
		fmt.Println("invoking SetPlatformCharge " + function)
		testBytes,err := SetPlatformCharge(args[0],stub)
		if err != nil {
			fmt.Println("Error performing SetPlatformCharge ")
			return nil, err
		}
		fmt.Println("Processed SetPlatformCharge successfully. ")
		return testBytes, nil
	}
	
	if function == "ListProposal" {
		fmt.Println("invoking ListProposal " + function)
		testBytes,err := ListProposal(args[0],stub)
		if err != nil {
			fmt.Println("Error performing ListProposal ")
			return nil, err
		}
		fmt.Println("Processed ListProposal successfully. ")
		return testBytes, nil
	}
	
	if function == "SignContract" {
		fmt.Println("invoking SignContract " + function)
		testBytes,err := SignContract(args[0],stub)
		if err != nil {
			fmt.Println("Error performing SignContract ")
			return nil, err
		}
		fmt.Println("Processed SignContract successfully. ")
		return testBytes, nil
	}
	
	if function == "MeterReading" {
		fmt.Println("invoking MeterReading " + function)
		testBytes,err := MeterReading(args[0],stub)
		if err != nil {
			fmt.Println("Error performing MeterReading ")
			return nil, err
		}
		fmt.Println("Processed MeterReading successfully. ")
		return testBytes, nil
	}
	
	if function == "BalanceUpdate" {
		fmt.Println("invoking BalanceUpdate " + function)
		testBytes,err := BalanceUpdate(args[0],stub)
		if err != nil {
			fmt.Println("Error performing BalanceUpdate ")
			return nil, err
		}
		fmt.Println("Processed BalanceUpdate successfully. ")
		return testBytes, nil
	}
	
	
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

func (self *EnergyTradingChainCode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
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

	if function == "GetUser" {
		fmt.Println("Invoking getUser " + function)
		var users User
		users,err := GetUsers(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving  the Users")
			return nil, errors.New("Error receiving  Users")
		}
		fmt.Println("All success, returning users")
		return json.Marshal(users)
	}
	
	if function == "GetProposals" {
		fmt.Println("Invoking GetProposals " + function)
		var proposals Proposal
		proposals,err := GetProposals(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving  the proposals")
			return nil, errors.New("Error receiving  proposals")
		}
		fmt.Println("All success, returning proposals")
		return json.Marshal(proposals)
	}
	
	return nil, errors.New("Received unknown query function name")

}

func GetUsers(userID string, stub shim.ChaincodeStubInterface)(User, error) {
	fmt.Println("In query.GetUsers start ")

	key := userID
	var users User
	userBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving Users" , userID)
		return users, errors.New("Error retrieving Users" + userID)
	}
	err = json.Unmarshal(userBytes, &users)
	fmt.Println("Users   : " , users);
	fmt.Println("In query.GetUsers end ")
	return users, nil
}

func GetProposals(proposalID string, stub shim.ChaincodeStubInterface)(Proposal, error) {
	fmt.Println("In query.GetProposals start ")
	key := proposalID
	var proposals Proposal
	proposalBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving proposals" , proposalID)
		return proposals, errors.New("Error retrieving Proposals" + proposalID)
	}
	err = json.Unmarshal(proposalBytes, &proposals)
	fmt.Println("Proposals   : " , proposals);
	fmt.Println("In query.GetProposals end ")
	return proposals, nil
}


func AddUser(userJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.AddUser start ")
	//smartMeterId :=args[0]
	//userType 	:=args[1]
	
	//var user User
	res := &User{}
	//user := &User{}
	err := json.Unmarshal([]byte(userJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal user ")
	}	
	fmt.Println("User ID : ",res.UserID)
	
	res.BuyPrice = "0"
	res.SellPrice = "0"
	res.EnergyConsumed = "0"
	res.EnergyProduced = "0"
	res.EnergyAccountBalance = "0"
	
	
	body, err := json.Marshal(res)
	if err != nil {
        panic(err)
    }
    fmt.Println(string(body))	
	err = stub.PutState(res.UserID, []byte(string(body)))
	if err != nil {
		fmt.Println("Failed to create User ")
	}
	fmt.Println("Created User  with Key : "+ res.UserID)
	fmt.Println("In initialize.AddUser end ")
	return nil,nil	
	
}


func SetGridPrice(gridPriceJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.SetGridPrice start ")
	//smartMeterId :=args[0]
	//userType 	:=args[1]
	
	
	res := &GridPrice{}
	
	err := json.Unmarshal([]byte(gridPriceJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal GridPrice ")
	}	
	fmt.Println("User ID : ",res.UserID)	
		
	err = stub.PutState(res.UserID + "_" + res.Date, []byte(gridPriceJSON))
	if err != nil {
		fmt.Println("Failed to create User ")
	}
	fmt.Println("Created User  with Key : "+ res.UserID)
	fmt.Println("In initialize.SetGridPrice end ")
	return nil,nil	
	
}


func SetPlatformCharge(chargeValue string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.SetPlatformCharge start ")
	//chargeValue :=args[0]
	
	key := "PlatformCharges"
	err := stub.PutState(key, []byte(chargeValue))
	if err != nil {
		fmt.Println("Failed to Set Platform Charge ")
	}
	fmt.Println("Created Charge  with Key : "+ key)
	fmt.Println("In initialize.SetPlatformCharge end ")
	return nil,nil	
	
}

func ListProposal(proposalJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.ListProposal start ")
		
	res := &Proposal{}
	
	err := json.Unmarshal([]byte(proposalJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal proposal ")
	}	
	fmt.Println("ProposalID  : ",res.ProposalID)
	
	users,err := GetUsers(res.UserID, stub)
	if err != nil {
		fmt.Println("Error receiving  the Users")
		return nil, errors.New("Error receiving  Users")
	}
	
	
	
	now := time.Now()
	//Getting the date only 	
	//dateValue := res.Date[:len(res.Date)-6]
	//20170406122460
	
	yearVal, err  := strconv.Atoi(res.Date[:len(res.Date)-10])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}	
	
	monthVal, err  := strconv.Atoi(res.Date[4:len(res.Date)-8])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}
	dayVal, err  := strconv.Atoi(res.Date[6:len(res.Date)-6])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}
	hourVal, err  := strconv.Atoi(res.Date[8:len(res.Date)-4])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}
	minutesVal, err  := strconv.Atoi(res.Date[10:len(res.Date)-2])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}
	secondsVal, err  := strconv.Atoi(res.Date[12:len(res.Date)])
	if (err != nil ){
		fmt.Println("Please pass integer ")
	}
	
	//formattedDate, errD := time.Date(strconv.Atoi(res.Date[:len(res.Date)-10]), strconv.Atoi(res.Date[4:len(res.Date)-8]), //strconv.Atoi(res.Date[6:len(res.Date)-6]), strconv.Atoi(res.Date[8:len(res.Date)-4]), strconv.Atoi(res.Date[10:len(res.Date)-2]), //strconv.Atoi(res.Date[12:len(res.Date)]), 000000000, time.UTC)
	
	formattedDate := time.Date(yearVal, time.Month(monthVal), dayVal, hourVal, minutesVal, secondsVal, 0, time.UTC)
	
	
	priceInt, errP := strconv.Atoi(res.Price);
	if errP != nil {
		fmt.Println("Error converting price")
		return nil, errors.New("Error converting price")
	}
	
	energyProposedInt, errEP := strconv.Atoi(res.EnergyProposed);
	if errEP != nil {
		fmt.Println("Error converting Energy proposed")
		return nil, errors.New("Error converting Energy proposed")
	}
	
	if(formattedDate.After(now) && priceInt > 0 && energyProposedInt > 0 && users.UserID != "")	{
		res.Status = "OPEN"
		res.EnergySigned = "0"
		res.EnergyRemaining = res.EnergyProposed	
		
		body, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))	
		err = stub.PutState(res.ProposalID, []byte(string(body)))
		if err != nil {
			fmt.Println("Failed to create Proposal ")
		}
		fmt.Println("Created User  with Key : "+ res.ProposalID)
	}else{
		fmt.Println("Proposal is not Valid. Enter future date or valid price or proposed energy value.")
		return nil, errors.New("Error listing proposal")
	}	
	
	fmt.Println("In initialize.ListProposal end ")
	return nil,nil	
	
}


func SignContract(signContractJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.SignContract start ")
		
	res := &Contract{}
	
	err := json.Unmarshal([]byte(signContractJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal Contract ")
	}	
	fmt.Println("ContractID  : ",res.ContractID)
	
	res.Status = "SIGNED"
	//res.EnergySigned = "0"
	res.EnergyConsumed = "0"
	//Price = <proposal>-Price. Get Proposal price
	res.Price = "0"
	res.BatteryBuyPrice = "0"
	res.BatterySellPrice = "0"
	//GridPrice = <gridUser>-price
	res.GridPrice = "0"
	res.PlatformComission = "0"
	//producer = <proposal>-producer
	res.Producer = "0"
	res.Consumer = res.UserID
	res.Battery = "null"
	//Grid = <gridUser>-usrid
	res.Grid = "0"
	//platform = <platform>-platformid
	res.Platform = "0"
	
	
	body, err := json.Marshal(res)
	if err != nil {
        panic(err)
    }
    fmt.Println(string(body))	
	err = stub.PutState(res.ContractID, []byte(string(body)))
	if err != nil {
		fmt.Println("Failed to create User ")
	}
	fmt.Println("Signed Contract  with Key : "+ res.ContractID)
	fmt.Println("In initialize.SignContract end ")
	return nil,nil	
	
}

func MeterReading(meterReadingJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.MeterReading start ")
		
	res := &Meter{}
	
	err := json.Unmarshal([]byte(meterReadingJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal Meter ")
	}	
	fmt.Println("EnergyReadingId  : ",res.EnergyReadingId)
	
	
	err = stub.PutState(res.EnergyReadingId, []byte(meterReadingJSON))
	if err != nil {
		fmt.Println("Failed to create Meter Reading ")
	}
	fmt.Println("Created Meter Reading  with Key : "+ res.EnergyReadingId)
	fmt.Println("In initialize.MeterReading end ")
	return nil,nil	
	
	
}


func BalanceUpdate(balanceUpdateJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.BalanceUpdate start ")
		
	res := &Balance{}
	
	err := json.Unmarshal([]byte(balanceUpdateJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal Balance ")
	}	
	
	fmt.Println("Balance of user  : ",res.Balance)	
	
	err = stub.PutState(res.UserID + "_" + "Balance", []byte(balanceUpdateJSON))
	if err != nil {
		fmt.Println("Failed to Update Balance ")
	}
	fmt.Println("Updated Balance with Key : "+ res.UserID + "_" + "Balance")
	fmt.Println("In initialize.BalanceUpdate end ")
	return nil,nil		
	
}


func main() {
	err := shim.Start(new(EnergyTradingChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}


}
