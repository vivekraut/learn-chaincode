
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

type PlatformCharge struct{
	PlatformID string `json:"PlatformID"`
	Charge string `json:"Charge"`		
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
	Consumer []User `json:"Consumer"`
	Contract []Contract `json:"Contract"`
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
	Producer User `json:"Producer"`	
	Consumer User `json:"Consumer"`	
	Battery string `json:"Battery"`	
	Grid string `json:"Grid"`
	Platform string `json:"Platform"`
	
}

type Meter struct{
	EnergyReadingId string `json:"EnergyReadingId"`	
	Date string `json:"Date"`	
	SmartMeterID string `json:"SmartMeterID"`
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
	
	if function == "GetGridPrice" {
		fmt.Println("Invoking GetGridPrice " + function)
		var gridPrice GridPrice
		gridPrice,err := GetGridPrice(args[0], stub)
		if err != nil {
			fmt.Println("Error retrieving the Grid Price")
			return nil, errors.New("Error retrieving the Grid Price")
		}
		fmt.Println("All success, returning grid price")
		return json.Marshal(gridPrice)
		
	}
	
	if function == "GetPlatformCharge" {
		fmt.Println("Invoking GetPlatformCharge " + function)
		var platformCharge PlatformCharge
		platformCharge,err := GetPlatformCharge(args[0], stub)
		if err != nil {
			fmt.Println("Error retrieving the Platform Charge")
			return nil, errors.New("Error retrieving the Platform Charge")
		}
		fmt.Println("All success, returning Platform Charge")
		return json.Marshal(platformCharge)
		
	}
	
	if function == "GetProposal" {
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
	
	if function == "GetContract" {
		fmt.Println("Invoking GetContract " + function)
		var contracts Contract
		contracts,err := GetContract(args[0], stub)
		if err != nil {
			fmt.Println("Error retrieving the contract")
			return nil, errors.New("Error retrieving the contract")
		}
		fmt.Println("All success, returning contract")
		return json.Marshal(contracts)
	}
	
	if function == "GetMeterReading" {
		fmt.Println("Invoking GetMeterReading " + function)
		var meterReading Meter
		meterReading,err := GetMeterReading(args[0], stub)
		if err != nil {
			fmt.Println("Error retrieving the meter reading")
			return nil, errors.New("Error retrieving the meter reading")
		}
		fmt.Println("All success, returning meter reading")
		return json.Marshal(meterReading)
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


func GetGridPrice(userIDDate string, stub shim.ChaincodeStubInterface)(GridPrice, error) {
	fmt.Println("In query.GetGridPrice start ")

	key := userIDDate
	var gridPrice GridPrice
	gridPriceBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving gridPrice" , userIDDate)
		return gridPrice, errors.New("Error retrieving Grid Price for the given user in this date" + userIDDate)
	}
	err = json.Unmarshal(gridPriceBytes, &gridPrice)
	fmt.Println("Grid Price   : " , gridPrice);
	fmt.Println("In query.GetGridPrice end ")
	return gridPrice, nil

}

func GetPlatformCharge(platformID string, stub shim.ChaincodeStubInterface)(PlatformCharge, error) {
	fmt.Println("In query.GetPlatformCharge start ")
	key := platformID
	var platformCharge PlatformCharge
	platformChargeBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving gridPrice" , platformID)
		return platformCharge, errors.New("Error retrieving Grid Price for the given user in this date" + platformID)
	}
	err = json.Unmarshal(platformChargeBytes, &platformCharge)
	fmt.Println("Platform Charge   : " , platformCharge);
	fmt.Println("In query.GetPlatformCharge end ")
	return platformCharge, nil
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


func GetContract(contractID string, stub shim.ChaincodeStubInterface)(Contract, error) {
	fmt.Println("In query.GetContract start ")
	key := contractID
	var contracts Contract
	contractBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving contract" , contractID)
		return contracts, errors.New("Error retrieving contract" + contractID)
	}
	err = json.Unmarshal(contractBytes, &contracts)
	fmt.Println("Contract   : " , contracts);
	fmt.Println("In query.GetContract end ")
	return contracts, nil
}

func GetMeterReading(energyReadingID string, stub shim.ChaincodeStubInterface)(Meter, error) {
	fmt.Println("In query.GetMeterReading start ")
	key := energyReadingID
	var meterReading Meter
	meterReadingBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving meter reading" , energyReadingID)
		return meterReading, errors.New("Error retrieving meter reading" + energyReadingID)
	}
	err = json.Unmarshal(meterReadingBytes, &meterReading)
	fmt.Println("Contract   : " , meterReading);
	fmt.Println("In query.GetMeterReading end ")
	return meterReading, nil
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
	err = stub.PutState(res.UserID + "_" + res.UserType, []byte(string(body)))
	if err != nil {
		fmt.Println("Failed to create User ")
	}
	
	err = stub.PutState(res.SmartMeterID, []byte(res.UserID))
	if err != nil {
		fmt.Println("Failed to set smart meter ID ")
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
	dateValue := res.Date[:len(res.Date)-6]	
	err = stub.PutState(res.UserID + "_" + dateValue, []byte(gridPriceJSON))
	if err != nil {
		fmt.Println("Failed to create User ")
	}
	fmt.Println("Created User  with Key : "+ res.UserID)
	fmt.Println("In initialize.SetGridPrice end ")
	return nil,nil	
	
}


func SetPlatformCharge(platformJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.SetPlatformCharge start ")
	//chargeValue :=args[0]
	
	res := &PlatformCharge{}
	err := json.Unmarshal([]byte(platformJSON), res)
	if err != nil {
		fmt.Println("Failed to unmarshal PlatformCharge ")
	}	
	
	//key := "PlatformCharges"
	err = stub.PutState(res.PlatformID, []byte(platformJSON))
	if err != nil {
		fmt.Println("Failed to Set Platform Charge ")
	}
	fmt.Println("Created Charge  with Key : "+ res.PlatformID)
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
	
	var users User
	users,err = GetUsers(res.UserID + "_" + "Prosumer", stub)
	if err != nil {
		fmt.Println("Error receiving  the Users")
		return nil, errors.New("Error receiving  Users")
	}
	
	
	now := time.Now()
	//Getting the date only 	
	dateValue := res.Date[:len(res.Date)-6]
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
	
	var gridPriceKey string
	gridPriceKey = res.UserID + "_" + dateValue
	fmt.Println("Grid Price Key...."+gridPriceKey)
	
	gridPrice,err := GetGridPrice(gridPriceKey, stub)
	if err != nil {
		fmt.Println("Error retrieving grid price")
		return nil, errors.New("Error retrieving grid price")
	}
	
	fmt.Println("Grid Price -->"+res.Price)
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
		var priceFloat float64
		priceFloat = float64(priceInt)*1.1
		//var priceInt int64
		gridPriceInt, errGP := strconv.Atoi(gridPrice.Price)
		if errGP != nil {
			fmt.Println("Error converting grid price")
			return nil, errors.New("Error converting grid price")
		}
		if(priceFloat > float64(gridPriceInt)){
			fmt.Println("Error - Price too high")
			return nil, errors.New("Error - Price too high")		
		}
		
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
	
	/*
	var producer User
	var consumer User
	producer = res.Producer
	consumer = res.Consumer
	
	prodId := producer.UserID
	consumerId := consumer.UserID
	
	fmt.Println("Getting producer details ")
	producerInfo,err := GetUsers(prodId + "_" + "Prosumer", stub)
	if err != nil {
		fmt.Println("Error retrieving the producer details")
		return nil, errors.New("Error retrieving the producer details")
	}
	
	fmt.Println("Getting consumer details ")
	consumerInfo,err := GetUsers(consumerId + "_" + "Prosumer", stub)
	if err != nil {
		fmt.Println("Error retrieving the producer details")
		return nil, errors.New("Error retrieving the consumer details")
	}
	*/
	gridUserInfo,err := GetUsers("0" + "_" + "Grid", stub)
	if err != nil {
		fmt.Println("Error retrieving the grid user details")
		return nil, errors.New("Error retrieving the grid user details")
	}
	fmt.Println("Grid User")
	fmt.Println(gridUserInfo)
	
	proposal,err := GetProposals(res.ProposalID, stub)
	now := time.Now()
	//Getting the date only 	
	dateValue := res.Date[:len(res.Date)-6]
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
	
	fmt.Println("Grid Price Date --> "+(strconv.Itoa(yearVal)+strconv.Itoa(monthVal)+strconv.Itoa(dayVal)))
	fmt.Println("Getting grid price for date ....")
	fmt.Println(dateValue)
	gridPriceInfo,err := GetGridPrice("0" + "_" + dateValue, stub)
	if err != nil {
		fmt.Println("Error retrieving the grid price details")
		return nil, errors.New("Error retrieving the grid price details")
	}
	
	energySignedInt, errEP := strconv.Atoi(res.EnergySigned);
	if errEP != nil {
		fmt.Println("Error converting Energy Signed")
		return nil, errors.New("Error converting Energy Signed")
	}
	fmt.Println("Energy Signed - "+(strconv.Itoa(energySignedInt)))
	proposalEnergyRemInt, errER := strconv.Atoi(proposal.EnergyRemaining);
	if errER != nil {
		fmt.Println("Error converting Energy Remaining")
		return nil, errors.New("Error converting Energy Remaining")
	}
	fmt.Println("Energy Proposed - "+(strconv.Itoa(proposalEnergyRemInt)))
	proposalEnergySignedInt, errPES := strconv.Atoi(proposal.EnergySigned);
	if errPES != nil {
		fmt.Println("Error converting Energy Signed")
		return nil, errors.New("Error converting Energy Signed")
	}
	fmt.Println("Proposal Energy Signed - "+(strconv.Itoa(proposalEnergySignedInt)))
	fmt.Println("Grid User ID.....")
	fmt.Println(gridUserInfo.UserID)
	fmt.Println("Grid Price User ID.....")
	fmt.Println(gridPriceInfo.UserID)
	
	//if(producerInfo.UserID != "" && consumerInfo.UserID != ""){		
		if(proposal.ProposalID != ""){		
			if(formattedDate.After(now) && energySignedInt > 0)	{	
				if(gridUserInfo.UserID != "" && gridPriceInfo.UserID != ""){	
					fmt.Println("User Found. Signing the contract")
					gridUserPrice := gridPriceInfo.Price
					platformChargeInfo,err := GetPlatformCharge("0", stub)
					if err != nil {
						fmt.Println("Error retrieving the platform charge details")
						return nil, errors.New("Error retrieving the platform charge details")
					}
					
					platformCharge := platformChargeInfo.Charge					
					fmt.Println("Platform Charge -->"+platformCharge)
					if(energySignedInt > proposalEnergyRemInt){
						fmt.Println("Error - Energy not available")
						return nil, errors.New("Error - Energy not available")					
					}
					
					res.Status = "SIGNED"
					//res.EnergySigned = "0"
					res.EnergyConsumed = "0"
					
					//Price = <proposal>-Price. Get Proposal price
					res.Price = proposal.Price
					res.BatteryBuyPrice = "0"
					res.BatterySellPrice = "0"
					//GridPrice = <gridUser>-price
					fmt.Println("Grid Price -->"+gridUserPrice)
					res.GridPrice = gridUserPrice
					res.PlatformComission = platformCharge
					//producer = <proposal>-producer
					//res.Producer = producerInfo.UserID
					//res.Consumer = consumerInfo.UserID
					res.Battery = "null"
					//Grid = <gridUser>-usrid
					res.Grid = gridUserInfo.UserID
					//platform = <platform>-platformid
					res.Platform = platformChargeInfo.PlatformID
					
					body, err := json.Marshal(res)
					if err != nil {
						panic(err)
					}
					fmt.Println(string(body))	
					err = stub.PutState(res.ContractID, []byte(string(body)))
					if err != nil {
						fmt.Println("Failed to create contract ")
					}
					fmt.Println("Signed Contract  with Key : "+ res.ContractID)
				}
			}
			
			energyRem := proposalEnergyRemInt - energySignedInt
			
			if(energyRem == 0){
				proposal.Status = "CLOSED"			
			}else{
				proposal.Status = "OPEN"
			}
			
			energySign := proposalEnergySignedInt + energySignedInt
			proposal.EnergySigned = strconv.Itoa(energySign)
			proposal.EnergyRemaining = strconv.Itoa(energyRem)
			
			var consumerRecord User
			consumerRecord.UserID = res.UserID
			var contractRecord Contract
			contractRecord.ContractID = res.ContractID
			
			proposalUpdate, err := json.Marshal(proposal)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(proposalUpdate))	
			err = stub.PutState(res.ContractID, []byte(string(proposalUpdate)))
			if err != nil {
				fmt.Println("Failed to update proposal ")
			}
			fmt.Println("Updated Proposal  with Key : "+ proposal.ProposalID)
			
		}		
	//}
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
	
	//var readingDate = res.Date	
	
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
	
	now := time.Now()
	formattedDate := time.Date(yearVal, time.Month(monthVal), dayVal, hourVal, minutesVal, secondsVal, 0, time.UTC)	
	//datString := res.Date[:len(res.Date)-10] + "-" + res.Date[4:len(res.Date)-8] + "-" + res.Date[6:len(res.Date)-6]
	
	nowString := now.String()
	dateString := formattedDate.String()
	
	energyAmtInt, errEM := strconv.Atoi(res.EnergyAmount);
	if errEM != nil {
		fmt.Println("Error converting Energy Amount")
		return nil, errors.New("Error converting Energy Amount")
	}
	
	if(nowString[:10] == dateString[:10]){
		if(energyAmtInt != 0){
			//var userIDByte res.UserID
			userIDValBytes, err := stub.GetState(res.SmartMeterID)
			if err != nil {
				fmt.Println("Error retrieving user ID")
				return nil, errors.New("Error retrieving User Details")
			}
			userIDVal := string(userIDValBytes)
			
			var users User			
			users,err = GetUsers(userIDVal + "_" + "Prosumer", stub)
			if err != nil {
				fmt.Println("Error receiving  the Users")
				return nil, errors.New("Error receiving  Users")
			}	
						
			body, err := json.Marshal(res)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))	
			err = stub.PutState(res.EnergyReadingId, []byte(string(body)))
			if err != nil {
				fmt.Println("Failed to create Meter Reading ")
			}
			fmt.Println("Created Meter Reading  with Key : "+ res.EnergyReadingId)					
			
			energyConsInt, errEC := strconv.Atoi(users.EnergyConsumed);
			if errEC != nil {
				fmt.Println("Error converting Energy Consumed")
				return nil, errors.New("Error converting Energy Consumed")
			}
			
			energyProdInt, errEP := strconv.Atoi(users.EnergyProduced);
			if errEP != nil {
				fmt.Println("Error converting Energy Produced")
				return nil, errors.New("Error converting Energy Produced")
			}
			
			
			if(energyAmtInt > 0){
				energyConsumedVal := energyConsInt + energyAmtInt
				energyProdVal := energyProdInt
				users.EnergyConsumed = 	strconv.Itoa(energyConsumedVal)			
				users.EnergyProduced = 	strconv.Itoa(energyProdVal)		
			} else{
				energyConsumedVal := energyConsInt 
				energyProdVal := energyProdInt + energyAmtInt*(-1)
				users.EnergyConsumed = 	strconv.Itoa(energyConsumedVal)			
				users.EnergyProduced = 	strconv.Itoa(energyProdVal)		
			}		

			
			bodyUser, err := json.Marshal(users)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bodyUser))	
			err = stub.PutState(users.UserID, []byte(string(bodyUser)))
			if err != nil {
				fmt.Println("Failed to update User Details ")
			}
			fmt.Println("Updated user details  with Key : "+ users.UserID)			
			
		}			
	}		
	
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
