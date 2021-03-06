package main

	import (
		"errors"
		"fmt"
		"strconv"
		"time"
		"strings"
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
		AccountBalance string `json:"AccountBalance"`
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
		Battery User `json:"Battery"`
		Grid User `json:"Grid"`
		Platform User `json:"Platform"`
		GridVolume string `json:"GridVolume"`
		BatteryVolume string `json:"BatteryVolume"`
		ProducerVolume string `json:"BatteryVolume"`
		ChangeInGridBalance     string `json:"ChangeInGridBalance"`
		ChangeInProducerBalance string `json:"ChangeInProducerBalance"`
		ChangeInBatteryBalance  string `json:"ChangeInBatteryBalance"`
		ChangeInConsumerBalance string `json:"ChangeInConsumerBalance"`
		ChangeInPlatformBalance string `json:"ChangeInPlatformBalance"`
		ChangeInGridBalanceOld  string `json:"ChangeInGridBalanceOld"`
		ChangeInProducerBalanceOld string `json:"ChangeInProducerBalanceOld"`
		ChangeInBatteryBalanceOld string `json:"ChangeInBatteryBalanceOld"`
		ChangeInConsumerBalanceOld string `json:"ChangeInConsumerBalanceOld"`
		ChangeInPlatformBalanceOld string `json:"ChangeInPlatformBalanceOld"`
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

	type ContractList []string

	type ProposalList []string

	func (self *EnergyTradingChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
		fmt.Println("In Init start ")

		var UserID, SmartMeterID, UserType string

		if len(args) != 3 {
			return nil, errors.New("Incorrect number  of  arguments. Expecting 3 - UserID, SmartMeterID and UserType. UserType should be Prosumer, Battery or Grid")
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
		fmt.Println("Initialization No functions found  ")
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

		if function == "SubmitProposal" {
			fmt.Println("invoking ListProposal " + function)
			testBytes,err := SubmitProposal(args[0],stub)
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
				fmt.Println("Error performing SignContract - ")
				fmt.Println(err)
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

		if function == "PerformSettlement" {
			fmt.Println("invoking PerformSettlement " + function)
			testBytes,err := PerformSettlement(args[0],stub)
			if err != nil {
				fmt.Println("Error performing PerformSettlement ")
				return nil, err
			}
			fmt.Println("Processed PerformSettlement successfully. ")
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

		if function == "GetAllContract" {
			fmt.Println("Invoking GetAllContract " + function)
			var contracts []Contract
			contracts,err := GetAllContract(args[0], stub)
			if err != nil {
				fmt.Println("Error retrieving the contract")
				return nil, errors.New("Error retrieving the contract")
			}
			fmt.Println("All success, returning All contract")
			return json.Marshal(contracts)
		}

		if function == "GetAllProposal" {
			fmt.Println("Invoking GetAllProposal " + function)
			var proposals []Proposal
			proposals,err := GetAllProposal(args[0], stub)
			if err != nil {
				fmt.Println("Error retrieving the proposals")
				return nil, errors.New("Error retrieving the proposals")
			}
			fmt.Println("All success, returning All proposals")
			return json.Marshal(proposals)
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

		if function == "GetBalance" {
			fmt.Println("Invoking GetBalance " + function)
			var balanceUser Balance
			balanceUser,err := GetBalance(args[0], stub)
			if err != nil {
				fmt.Println("Error retrieving the user balance")
				return nil, errors.New("Error retrieving the user balance")
			}
			fmt.Println("All success, returning user balance")
			return json.Marshal(balanceUser)
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
		fmt.Println("Getting the contract for ID..." , key)
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


	func GetProposal(ProposalID string, stub shim.ChaincodeStubInterface)(Proposal, error) {
		fmt.Println("In query.GetProposal start ")
		key := ProposalID
		fmt.Println("Getting the Proposal for ID..." , key)
		var Proposals Proposal
		ProposalBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving Proposal" , ProposalID)
			return Proposals, errors.New("Error retrieving Proposal" + ProposalID)
		}
		err = json.Unmarshal(ProposalBytes, &Proposals)
		fmt.Println("Proposal   : " , Proposals);
		fmt.Println("In query.GetProposal end ")
		return Proposals, nil
	}


	func GetAllContract(ContractDate string, stub shim.ChaincodeStubInterface)([]Contract, error) {
		fmt.Println("In query.GetAllContract start ")
		key := "contract_"+ContractDate
		fmt.Println("Getting the contract for Date..." , key)
		var contracts []Contract
		var contractsIds ContractList

		contractListBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving contract List for Date" , ContractDate)
			return contracts, errors.New("Error retrieving contract list for Date" + ContractDate)
		}
		err = json.Unmarshal(contractListBytes, &contractsIds)
		fmt.Println("Contract Ids   : " , contractsIds);
		for _, contractId := range contractsIds {
			fmt.Println("Getting Contract ",contractId)
			v_contract,err := GetContract(contractId, stub)
			if err != nil {
				fmt.Println("Error retrieving the contract", contractId)
			} else {
				fmt.Println("Contract Retrived for", contractId)
				contracts = append(contracts,v_contract)
			}
		}

		fmt.Println("In query.GetContract end ")
		return contracts, nil
	}

	func GetAllProposal(ProposalDate string, stub shim.ChaincodeStubInterface)([]Proposal, error) {
		fmt.Println("In query.GetAllProposal start ")
		key := "proposal_"+ProposalDate
		fmt.Println("Getting the Proposal for Date..." , key)
		var proposals []Proposal
		var proposalIds ProposalList

		proposalListBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving proposals List for Date" , ProposalDate)
			return proposals, errors.New("Error retrieving contract list for Date" + ProposalDate)
		}
		err = json.Unmarshal(proposalListBytes, &proposalIds)
		fmt.Println("proposals Ids   : " , proposalIds);
		for _, proposalId := range proposalIds {
			fmt.Println("Getting Contract ",proposalId)
			v_proposal,err := GetProposal(proposalId, stub)
			if err != nil {
				fmt.Println("Error retrieving the contract", proposalId)
			} else {
				fmt.Println("Contract Retrived for", proposalId)
				proposals = append(proposals,v_proposal)
			}
		}

		fmt.Println("In query.GetContract end ")
		return proposals, nil
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
		fmt.Println("Meter Reading   : " , meterReading);
		fmt.Println("In query.GetMeterReading end ")
		return meterReading, nil
	}

	func GetBalance(userID string, stub shim.ChaincodeStubInterface)(Balance, error) {
		fmt.Println("In query.GetBalance start ")
		key := userID
		var balanceUser Balance
		balanceUserBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving user balance" , userID)
			return balanceUser, errors.New("Error retrieving user balance" + userID)
		}
		err = json.Unmarshal(balanceUserBytes, &balanceUser)
		fmt.Println("User Balance   : " , balanceUser);
		fmt.Println("In query.GetBalance end ")
		return balanceUser, nil
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

	func SubmitProposal(proposalJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
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


		//now := time.Now()
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
		
		fmt.Println(formattedDate)
		var gridPriceKey string
		//gridPriceKey = res.UserID + "_" + dateValue
		gridPriceKey =  "0_" + dateValue
		fmt.Println("Grid Price Key...."+gridPriceKey)

		gridPrice,err := GetGridPrice(gridPriceKey, stub)
		if err != nil {
			fmt.Println("Error retrieving grid price")
			return nil, errors.New("Error retrieving grid price")
		}

		fmt.Println("Proposal Price -->"+res.Price)
		priceInt, errP := strconv.ParseFloat(res.Price,64);
		if errP != nil {
			fmt.Println("Error converting price")
			return nil, errors.New("Error converting price")
		}

		energyProposedInt, errEP := strconv.ParseFloat(res.EnergyProposed,64);
		if errEP != nil {
			fmt.Println("Error converting Energy proposed")
			return nil, errors.New("Error converting Energy proposed")
		}

		if( priceInt > 0 && energyProposedInt > 0 && users.UserID != "")	{
			var priceFloat float64
			priceFloat = priceInt*1.1
			//var priceInt int64
			gridPriceInt, errGP := strconv.ParseFloat(gridPrice.Price,64)
			if errGP != nil {
				fmt.Println("Error converting grid price")
				return nil, errors.New("Error converting grid price")
			}
			if(priceFloat > gridPriceInt){
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
			} else {

				err2 := SetAllproposal(dateValue,res.ProposalID,stub)
				if(err2 != nil){
					fmt.Println("Failed to save proposal"+ res.ProposalID + "in proposal list")
				}

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
		//now := time.Now()
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
			if( energySignedInt > 0)	{
				if(gridUserInfo.UserID != "" && gridPriceInfo.UserID != ""){
					fmt.Println("User Found. Signing the contract")
					gridUserPrice := gridPriceInfo.Price
					// Default Platform ID - P001
					platformChargeInfo,err := GetPlatformCharge("P001", stub)
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
					res.Producer.UserID = proposal.UserID
					res.Producer.EnergyProduced = proposal.EnergyProposed

					producerUser,err := GetUsers(res.Producer.UserID+"_Prosumer", stub)
					res.Producer.EnergyConsumed = producerUser.EnergyConsumed
					//res.Producer.EnergyProduced = producerUser.EnergyProduced
					res.Producer.UserType = 	producerUser.UserType
					res.Producer.EnergyAccountBalance = producerUser.EnergyAccountBalance
					res.Producer.SmartMeterID = producerUser.SmartMeterID
					producerUpdate, err := json.Marshal(res.Producer)
					if err != nil {
						panic(err)
					}

					fmt.Println(string(producerUpdate))
					err = stub.PutState(res.Producer.UserID+"_Prosumer", []byte(string(producerUpdate)))
					if err != nil {
						fmt.Println("Failed to update producer ")
					}
					fmt.Println("Updated producer  with Key : "+ res.Producer.UserID+"_Prosumer")

					res.Consumer.UserID = res.UserID
					// Default Battery ID - B001
					consumerUser,err := GetUsers(res.UserID+"_Prosumer", stub)
					res.Consumer.EnergyConsumed = consumerUser.EnergyConsumed
					res.Consumer.EnergyProduced = consumerUser.EnergyProduced
					res.Consumer.UserType = 	consumerUser.UserType
					res.Consumer.EnergyAccountBalance = consumerUser.EnergyAccountBalance
					res.Consumer.SmartMeterID = consumerUser.SmartMeterID

					batteryUser,err := GetUsers("B001_Battery", stub)
					res.Battery.UserID = "B001"
					res.Battery.EnergyConsumed = batteryUser.EnergyConsumed
					res.Battery.EnergyProduced = batteryUser.EnergyProduced
					res.Battery.UserType = 	batteryUser.UserType

					/*
					var batteryUser User
					batteryUser,err = GetUsers("B001" + "_" + "Battery", stub)
					if err != nil {
						fmt.Println("Error getting  the battery user details")
						return nil, errors.New("Error getting  the battery user details")
					}
					*/


					//Grid = <gridUser>-usrid
					res.Grid.UserID = gridUserInfo.UserID
					// Default platform ID - P001
					res.Platform.UserID = platformChargeInfo.PlatformID

					//Initializing Contract Values
					res.ChangeInBatteryBalanceOld  = "0"
					res.ChangeInProducerBalanceOld = "0"
					res.ChangeInGridBalanceOld = "0"
					res.ChangeInConsumerBalanceOld = "0"
					res.ChangeInPlatformBalanceOld = "0"
					res.ChangeInGridBalance = "0"
					res.ChangeInProducerBalance = "0"
					res.ChangeInBatteryBalance = "0"
					res.ChangeInConsumerBalance = "0"
					res.ChangeInPlatformBalance = "0"

					body, err := json.Marshal(res)
					if err != nil {
						panic(err)
					}
					fmt.Println(string(body))
					err = stub.PutState(res.ContractID, []byte(string(body)))
					fmt.Println(res.ContractID)
					if err != nil {
						fmt.Println("Failed to create contract ")
					} else {

						err2 := SetAllContract(dateValue,res.ContractID,stub)
						if(err2 != nil){
							fmt.Println("Failed to save contract"+ res.ContractID + "in contract list")
						}

					}

					fmt.Println("Signed Contract  with Key : "+ res.ContractID)

					// Setting the contract for the date
					fmt.Println("Contract Date -- ")
					fmt.Println(formattedDate)
					dateString := formattedDate.String()
					dateStringUpd := dateString[:10]

					fmt.Println(dateStringUpd)
					var contractIDsString string

					// Getting the values from ledger
					contractIDsBytes, err := stub.GetState(dateStringUpd)
					if err != nil {
						fmt.Println("Error retrieving contract IDs")
						return nil, errors.New("Error retrieving contract IDs")
						contractIDsString = ""
					}else{
						contractIDsVal := string(contractIDsBytes)
						contractIDsString = contractIDsVal + ","
					}

					fmt.Println(res.ContractID)
					err = stub.PutState(dateStringUpd, []byte(contractIDsString + res.ContractID))
					if err != nil {
						fmt.Println("Failed to create contract list ")
					}
					fmt.Println("Added contract to the list")
				} else {
					return nil,errors.New("Grid User not found or Grid Price is not Set")
				}

			} else {
				return nil,errors.New("Energy Signed Amount is Zero")
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
			err = stub.PutState(proposal.ProposalID, []byte(string(proposalUpdate)))
			if err != nil {
				fmt.Println("Failed to update proposal ")
			}
			fmt.Println("Updated Proposal  with Key : "+ proposal.ProposalID)

		} else {
			return nil,errors.New("Proposal Id is Emply")
		}
		//}
		fmt.Println("In initialize.SignContract end ")
		return nil,nil
	}

	func SetAllproposal(proposalDate string, proposalId string, stub shim.ChaincodeStubInterface)(error){
		errors.New("Unable to update proposal list for date " + proposalDate)

		key := "proposal_" + proposalDate
		var proposalIds ProposalList
		proposalListBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving proposal List for Date" , proposalDate)
			//Create One
			proposalIds[0] = proposalId
			proposalListByte, err := json.Marshal(proposalIds)
			err = stub.PutState(key,[]byte(string(proposalListByte)))
			if err!= nil{
				fmt.Println("Unable to Put proposal List for Date" , proposalDate)
			}
		} else {
			err = json.Unmarshal(proposalListBytes,&proposalIds)
			isExist := StringInSlice(proposalId, proposalIds)
			if(isExist){
				fmt.Println("proposal Id " + proposalId + "already exist in List for Date "+ proposalDate)
				return nil
			} else {
				proposalIds = append (proposalIds, proposalId )
				proposalListByte, err := json.Marshal(proposalIds)
				err = stub.PutState(key,[]byte(string(proposalListByte)))
				if err!= nil{
					fmt.Println("Unable to Put proposal List for Date" , proposalDate)
				}
			}
		}
		return nil
	}

	func SetAllContract(contractDate string, contractId string, stub shim.ChaincodeStubInterface)(error){
		errors.New("Unable to update contract list for date " + contractDate)

		key := "contract_" + contractDate
		var contractIds ContractList
		ContractListBytes, err := stub.GetState(key)
		if err != nil {
			fmt.Println("Error retrieving Contract List for Date" , contractDate)
			//Create One
			contractIds[0] = contractId
			contractListByte, err := json.Marshal(contractIds)
			err = stub.PutState(key,[]byte(string(contractListByte)))
			if err!= nil{
				fmt.Println("Unable to Put Contract List for Date" , contractDate)
			}
		} else {
			err = json.Unmarshal(ContractListBytes,&contractIds)
			isExist := StringInSlice(contractId, contractIds)
			if(isExist){
				fmt.Println("Contract Id " + contractId + "already exist in List for Date "+ contractDate)
				return nil
			} else {
				contractIds = append (contractIds, contractId )
				contractListByte, err := json.Marshal(contractIds)
				err = stub.PutState(key,[]byte(string(contractListByte)))
				if err!= nil{
					fmt.Println("Unable to Put Contract List for Date" , contractDate)
				}
			}
		}
		return nil
	}


	func StringInSlice(a string, list []string) bool {
		for _, b := range list {
			if b == a {
				return true
			}
		}
		return false
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
		fmt.Println("Dates....")
		fmt.Println(nowString)
		fmt.Println(dateString)
		energyAmtInt, errEM := strconv.ParseFloat(res.EnergyAmount,64);
		if errEM != nil {
			fmt.Println("Error converting Energy Amount")
			return nil, errors.New("Error converting Energy Amount")
		}

		fmt.Println(energyAmtInt)

		if(nowString[:10] == dateString[:10]){
			if(energyAmtInt != 0){
				fmt.Println("condition successful ")
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

				energyConsInt, errEC := strconv.ParseFloat(users.EnergyConsumed,64);
				if errEC != nil {
					fmt.Println("Error converting Energy Consumed")
					return nil, errors.New("Error converting Energy Consumed")
				}

				energyProdInt, errEP := strconv.ParseFloat(users.EnergyProduced,64);
				if errEP != nil {
					fmt.Println("Error converting Energy Produced")
					return nil, errors.New("Error converting Energy Produced")
				}

				fmt.Println("Energy Amount (Integer)")
				fmt.Println(energyAmtInt)

				if(energyAmtInt > 0){
					energyConsumedVal := energyConsInt + energyAmtInt
					energyProdVal := energyProdInt
					users.EnergyConsumed = 	strconv.FormatFloat(energyConsumedVal ,'f', 2, 32)
					users.EnergyProduced = 	strconv.FormatFloat(energyProdVal ,'f', 2, 32)
				} else{
					energyConsumedVal := energyConsInt
					energyProdVal := energyProdInt + energyAmtInt*(-1)
					users.EnergyConsumed = 	strconv.FormatFloat(energyConsumedVal ,'f', 2, 32)
					users.EnergyProduced = 	strconv.FormatFloat(energyProdVal ,'f', 2, 32)

				}
				fmt.Println("Energy Consumed")
				fmt.Println(users.EnergyConsumed)
				fmt.Println("Energy Produced")
				fmt.Println(users.EnergyProduced)

				bodyUser, err := json.Marshal(users)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(bodyUser))
				err = stub.PutState(users.UserID+"_Prosumer", []byte(string(bodyUser)))
				if err != nil {
					fmt.Println("Failed to update User Details ")
				}
				fmt.Println("Updated user details  with Key : "+ (users.UserID+"_Prosumer"))

			}
		}

		fmt.Println("In initialize.MeterReading end ")

		presentDate := dateString[:10]
		PerformSettlement(presentDate,stub)
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

		consumerUser,err := GetUsers(res.UserID, stub)
		consumerUser.EnergyAccountBalance = res.Balance

		consumerUpdate, err := json.Marshal(consumerUser)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(consumerUpdate))
		err = stub.PutState(res.UserID, []byte(string(consumerUpdate)))
		if err != nil {
			fmt.Println("Failed to update User Balance ")
		}
		fmt.Println("Updated User Balance  with Key : "+ res.UserID)
		fmt.Println("In initialize.BalanceUpdate end ")
		return nil,nil
	}


	func PerformSettlement(dateVal string, stub shim.ChaincodeStubInterface) ([]byte, error) {
		fmt.Println("In services.PerformSettlement start ")


		contractIDsBytes, err := stub.GetState(dateVal)
		if err != nil {
			fmt.Println("Error retrieving contract IDs")
			return nil, errors.New("Error retrieving contract IDs")
		}

		var contractIDs [] string
		contractIDs = strings.Split(string(contractIDsBytes), ",")

		var contracts [] Contract;

		for i := 1; i < len(contractIDs); i++{
			contractIDVal := contractIDs[i];
			var contractDetails Contract
			contractDetails,err = GetContract(contractIDVal, stub)
			if err != nil {
				fmt.Println("Error receiving  the contract details")
				return nil, errors.New("Error receiving  contract details")
			}
			contracts = append(contracts, contractDetails)

		}


		//Get Contract 1 - con1
		//Get Contract 2 - con2
		//var con1 Contract
		//var con2 Contract
		//contracts = append(contracts, con1)
		//contracts = append(contracts, con2)

		var updatedContracts[] Contract;

		var ContractIdMap  map[string]Contract;
		ContractIdMap = make(map[string]Contract);

		var ContractPerConsumer  map[string]string;
		ContractPerConsumer = make(map[string]string);

		var ContractPerProducer  map[string]([]string);
		ContractPerProducer = make(map[string]([]string));

		var ProducerContractVolume  map[string]float64;
		ProducerContractVolume  = make(map[string]float64);

		var ConsumerContractVolume map[string]float64;
		ConsumerContractVolume = make(map[string]float64);


		var ConsumerTotalConsumption map[string]float64;
		ConsumerTotalConsumption = make(map[string]float64);


		var ProducerTotalProduction map[string]float64;
		ProducerTotalProduction = make(map[string]float64);

		var BatterySupply float64
		BatterySupply = 0.0;
		var GridSupply float64
		GridSupply = 0.0
		var totalOverConsumed float64
		totalOverConsumed = 0.0
		var totalUnderConsumed float64
		totalUnderConsumed = 0.0
		var BatteryBuyPrice float64
		BatteryBuyPrice = 0.0

		var BatterySellPrice float64
		BatterySellPrice = 0.0

		//var totalProduced = 0.0;
		//var totalConsumed = 0.0;

		//var userActualVolume map[string]float64;
		//userActualVolume = make(map[string]float64);

		//var ConsumerActualVolume = new Map();

		if len(contracts)  > 0 {
			Battery := contracts[0].Battery
			val2 , _ := strconv.ParseFloat(Battery.EnergyProduced, 64)
			val3 , _ := strconv.ParseFloat(Battery.EnergyConsumed, 64)
			val5 , _ := strconv.ParseFloat(contracts[0].Price, 64);
			BatterySupply = val2 - val3
			fmt.Print("BatterySupply:");
			fmt.Println(BatterySupply);
			//fmt.Println(err);

			Grid := contracts[0].Grid
			val4 , _ := strconv.ParseFloat(Grid.EnergyProduced, 64)
			GridSupply = val4
			fmt.Print("GridSupply:");
			fmt.Println(GridSupply);
			//fmt.Println(err)
			BatteryBuyPrice = val5
			BatterySellPrice = val5
		}

		for i := 0; i < len(contracts); i++{
			var l_contract = contracts[i];
			producer := l_contract.Producer;
			consumer := l_contract.Consumer;

			producerUser,err := GetUsers(producer.UserID+"_Prosumer", stub)
			if err != nil {
				fmt.Println("Failed to retrieve producer ")
			}

			producer.EnergyConsumed = producerUser.EnergyConsumed
			producer.EnergyProduced = producerUser.EnergyProduced
			producer.UserType = 	producerUser.UserType
			producer.EnergyAccountBalance = producerUser.EnergyAccountBalance
			producer.SmartMeterID = producerUser.SmartMeterID

			consumerUser,err := GetUsers(consumer.UserID+"_Prosumer", stub)
			if err != nil {
				fmt.Println("Failed to retrieve consumer ")
			}
			consumer.EnergyConsumed = consumerUser.EnergyConsumed
			consumer.EnergyProduced = consumerUser.EnergyProduced
			consumer.UserType = 	consumerUser.UserType
			consumer.EnergyAccountBalance = consumerUser.EnergyAccountBalance
			consumer.SmartMeterID = consumerUser.SmartMeterID


			val1 , _ := strconv.ParseFloat(l_contract.EnergySigned, 64);
			val2 , _ := strconv.ParseFloat(producer.EnergyProduced, 64);
			val3 , _ := strconv.ParseFloat(producer.EnergyConsumed, 64);
			val4 , _ := strconv.ParseFloat(consumer.EnergyProduced, 64);
			val5 , _ := strconv.ParseFloat(consumer.EnergyConsumed, 64);
			val6 , _ := strconv.ParseFloat(l_contract.Price, 64);
			if val, ok := ProducerContractVolume[producer.UserID]; !ok {
				ProducerContractVolume[producer.UserID] = val1;
			} else{
				ProducerContractVolume[producer.UserID] = val1 + val;
			}

			if val, ok := ConsumerContractVolume[consumer.UserID]; !ok {
				ConsumerContractVolume[consumer.UserID] = val1;
			} else{
				ConsumerContractVolume[consumer.UserID] = val1 + val;
			}

			if _, ok := ConsumerTotalConsumption[consumer.UserID]; !ok {
				ConsumerTotalConsumption[consumer.UserID] = val5  - val4;
			}

			if _, ok := ProducerTotalProduction[producer.UserID]; !ok {
				ProducerTotalProduction[producer.UserID] = val2  - val3;
			}

			if BatteryBuyPrice > val6 {
				BatteryBuyPrice = val6
			}

			if BatterySellPrice < val6 {
				BatterySellPrice = val6
			}
			l_contract.EnergyConsumed = strconv.FormatFloat(val5 ,'f', 2, 32)

			ContractIdMap[l_contract.ContractID] = l_contract;
			ContractPerProducer[producer.UserID] = append(ContractPerProducer[producer.UserID],l_contract.ContractID)
			ContractPerConsumer[consumer.UserID] = l_contract.ContractID

			//fmt.Println(err);
		}


		BatterySellPrice = BatterySellPrice * 1.1;
		BatteryBuyPrice = BatteryBuyPrice * 0.9 ;

		//--Dividing Volume of Production
		for userid := range ProducerTotalProduction{
			if ProducerTotalProduction[userid] > 0 {
				for _,conid := range ContractPerProducer[userid] {
					m_contract := ContractIdMap[conid]
					val1 , _ := strconv.ParseFloat(m_contract.EnergySigned, 64);
					m_contract.ProducerVolume =  strconv.FormatFloat( (val1 / ProducerContractVolume[userid] * ProducerTotalProduction[userid]),'f', 2, 32)
					ContractIdMap[conid] = m_contract
				}
			}
		}

		for userid := range ConsumerTotalConsumption {
			if ConsumerTotalConsumption[userid] > ConsumerContractVolume[userid] {
				totalOverConsumed = totalOverConsumed + ConsumerTotalConsumption[userid] - ConsumerContractVolume[userid]
			} else if ConsumerTotalConsumption[userid] < ConsumerContractVolume[userid] {
				totalUnderConsumed = totalOverConsumed - ConsumerTotalConsumption[userid] + ConsumerContractVolume[userid]
			} else {
				totalOverConsumed = totalOverConsumed + ConsumerTotalConsumption[userid] - ConsumerContractVolume[userid]
			}
		}

		if totalOverConsumed > 0 {
			for userid := range ConsumerTotalConsumption {
				if ConsumerTotalConsumption[userid] > ConsumerContractVolume[userid] {
					m_contract := ContractIdMap[ContractPerConsumer[userid]]
					m_contract.GridVolume = strconv.FormatFloat(( ConsumerTotalConsumption[userid] - ConsumerContractVolume[userid] ) / totalOverConsumed ,'f', 2, 32)
				} else if ConsumerTotalConsumption[userid] < ConsumerContractVolume[userid] {
					totalUnderConsumed = totalOverConsumed - ConsumerTotalConsumption[userid] + ConsumerContractVolume[userid]
				} else {
					totalOverConsumed = totalOverConsumed + ConsumerTotalConsumption[userid] - ConsumerContractVolume[userid]
				}
			}

		}
		for conid := range ContractIdMap{
			t_con := ContractIdMap[conid]
			t_con.Status = "PROCESSED"
			t_con.BatteryBuyPrice = strconv.FormatFloat(BatteryBuyPrice, 'f', 2, 32)
			t_con.BatterySellPrice = strconv.FormatFloat(BatterySellPrice, 'f', 2, 32)
			val1 , _ := strconv.ParseFloat(t_con.GridVolume, 64);
			//val2 , _ := strconv.ParseFloat(t_con.BatteryVolume, 64);
			val3 , _ := strconv.ParseFloat(t_con.EnergyConsumed, 64);
			val4 , _ := strconv.ParseFloat(t_con.EnergySigned, 64);
			val5 , _ := strconv.ParseFloat(t_con.BatteryBuyPrice, 64);
			val6 , _ := strconv.ParseFloat(t_con.BatterySellPrice, 64);
			val7 , _ := strconv.ParseFloat(t_con.GridPrice, 64);
			val8 , _ := strconv.ParseFloat(t_con.Price, 64);
			//val9 , _ := strconv.ParseFloat(t_con.ProducerVolume, 64);
			val2 := val3 - val4 - val1
			var ChangeInBatteryBalance float64
			if val2  < 0 {
				ChangeInBatteryBalance = val2 * val5
			} else if val2 > 0 {
				ChangeInBatteryBalance = val2 * val6
			}

			ChangeInGridBalance := val1 * val7
			ChangeInProducerBalance := val4 * val8
			ChangeInConsumerBalance := - ( ChangeInProducerBalance + ChangeInGridBalance + ChangeInBatteryBalance)
			t_con.ChangeInBatteryBalanceOld  = t_con.ChangeInBatteryBalance
			t_con.ChangeInProducerBalanceOld = t_con.ChangeInProducerBalance
			t_con.ChangeInGridBalanceOld 	 = t_con.ChangeInGridBalance
			t_con.ChangeInConsumerBalanceOld = t_con.ChangeInConsumerBalance
			t_con.ChangeInPlatformBalanceOld = t_con.ChangeInPlatformBalance
			platformCharge, _ := GetPlatformCharge("P001",stub)
			platformChargeFloat ,_ := strconv.ParseFloat(platformCharge.Charge, 64);
			var ChangeInPlatformBalance float64
			if( ChangeInProducerBalance > 0 ){
				ChangeInProducerBalance = ChangeInProducerBalance * ( 1 - platformChargeFloat )
				ChangeInPlatformBalance = ChangeInProducerBalance * platformChargeFloat
			} else {
				ChangeInProducerBalance = ChangeInProducerBalance * ( 1 + platformChargeFloat )
				ChangeInPlatformBalance = -1 * ChangeInProducerBalance * platformChargeFloat
			}

			t_con.ChangeInBatteryBalance = strconv.FormatFloat(ChangeInBatteryBalance,'f', 2, 32)
			t_con.ChangeInProducerBalance = strconv.FormatFloat(ChangeInProducerBalance ,'f', 2, 32)
			t_con.ChangeInGridBalance = strconv.FormatFloat(ChangeInGridBalance ,'f', 2, 32)
			t_con.ChangeInConsumerBalance = strconv.FormatFloat(ChangeInConsumerBalance ,'f', 2, 32)
			t_con.ChangeInPlatformBalance = strconv.FormatFloat(ChangeInPlatformBalance ,'f', 2, 32)
			t_con.BatteryVolume = strconv.FormatFloat(val2 ,'f', 2, 32)
			updatedContracts = append(updatedContracts , t_con)

		}

		fmt.Print("GridSupply:");
		fmt.Println(GridSupply);
		fmt.Print("totalUnderConsumed:")
		fmt.Println(totalUnderConsumed)
		fmt.Print("totalOverConsumed:")
		fmt.Println(totalOverConsumed)
		fmt.Print("BatteryBuyPrice:")
		fmt.Println(BatteryBuyPrice)
		fmt.Print("BatterySellPrice:")
		fmt.Println(BatterySellPrice)
		fmt.Print("Contract:")
		//fmt.Println(ContractIdMap)

		for i := 0; i < len(updatedContracts); i++ {
			var con = updatedContracts[i];
			fmt.Println("")
			fmt.Println("")
			fmt.Print("ProposalID:")
			fmt.Println(con.ProposalID)
			fmt.Print("ContractID:")
			fmt.Println(con.ContractID)
			fmt.Print("EnergySigned:")
			fmt.Println(con.EnergySigned)
			fmt.Print("EnergyConsumed:")
			fmt.Println(con.EnergyConsumed)
			fmt.Print("Status:")
			fmt.Println(con.Status)
			fmt.Print("Price:")
			fmt.Println(con.Price)
			fmt.Print("BatteryBuyPrice:")
			fmt.Println(con.BatteryBuyPrice)
			fmt.Print("BatterySellPrice:")
			fmt.Println(con.BatterySellPrice)
			fmt.Print("GridPrice:")
			fmt.Println(con.GridPrice)
			fmt.Print("PlatformComission:")
			fmt.Println(con.PlatformComission)
			fmt.Print("GridVolume:")
			fmt.Println(con.GridVolume)
			fmt.Print("BatteryVolume:")
			fmt.Println(con.BatteryVolume)
			fmt.Print("ProducerVolume:")
			fmt.Println(con.ProducerVolume)
			fmt.Print("ChangeInGridBalance:")
			fmt.Println(con.ChangeInGridBalance)
			fmt.Print("ChangeInProducerBalance:")
			fmt.Println(con.ChangeInProducerBalance)
			fmt.Print("ChangeInBatteryBalance:")
			fmt.Println(con.ChangeInBatteryBalance)
			fmt.Print("ChangeInConsumerBalance:")
			fmt.Println(con.ChangeInConsumerBalance)

			// Updating the contract after the settlement

			// Producer balance update starts here
			producerUser, err := GetUsers(con.Producer.UserID + "_Prosumer", stub)
			if err != nil {
				fmt.Println("Failed to retrieve producer ")
			}
			con.Producer = producerUser

			prodBalanceFloat, errPB := strconv.ParseFloat(producerUser.EnergyAccountBalance, 64)
			if errPB != nil {
				fmt.Println("Error converting Producer Balance")
				return nil, errors.New("Error converting Producer Balance")
			}

			changeProdBalanceFloat, errCPB := strconv.ParseFloat(con.ChangeInProducerBalance, 64);
			if errCPB != nil {
				fmt.Println("Error converting Changed Producer Balance")
				return nil, errors.New("Error converting Changed Producer Balance")
			}

			changeProdBalanceOldFloat, errCPB := strconv.ParseFloat(con.ChangeInProducerBalanceOld, 64);
			if errCPB != nil {
				fmt.Println("Error converting Changed Producer Old Balance")
				return nil, errors.New("Error converting Changed Producer Old Balance")
			}

			prodBalanceFloat = prodBalanceFloat + changeProdBalanceFloat - changeProdBalanceOldFloat
			producerUser.EnergyAccountBalance = strconv.FormatFloat(prodBalanceFloat, 'f', 6, 64)

			// Producer balance update ends here


			//Updating Platform Balance
			platformUser, err := GetPlatformCharge("P001", stub)
			if err != nil {
				fmt.Println("Failed to retrieve Platform ")
			}

			platformBalanceFloat, errPB := strconv.ParseFloat(platformUser.AccountBalance, 64)
			if errPB != nil {
				fmt.Println("Error converting Platform Balance")
				return nil, errors.New("Error converting Platform Balance")
			}

			changePlatformBalanceFloat, errCPB := strconv.ParseFloat(con.ChangeInPlatformBalance, 64);
			if errCPB != nil {
				fmt.Println("Error converting Changed Platform Balance")
				return nil, errors.New("Error converting Changed Platform Balance")
			}

			changePlatformBalanceOldFloat, errCPB := strconv.ParseFloat(con.ChangeInPlatformBalanceOld, 64);
			if errCPB != nil {
				fmt.Println("Error converting Changed Platform Old Balance")
				return nil, errors.New("Error converting Changed Platform Old Balance")
			}

			platformBalanceFloat = platformBalanceFloat + changePlatformBalanceFloat - changePlatformBalanceOldFloat
			platformUser.AccountBalance = strconv.FormatFloat(platformBalanceFloat, 'f', 6, 64)
			PlatformUpdate, err := json.Marshal(platformUser)
			if err != nil {
				panic(err)
			}
			SetPlatformCharge(string(PlatformUpdate),stub)
			//Updating Platform Balance ends here

			// Updating Consumer Balance
			consumerUser,err := GetUsers(con.Consumer.UserID+"_Prosumer", stub)
			if err != nil {
				fmt.Println("Failed to retrieve consumer ")
			}
			con.Consumer = consumerUser
			consumerBalanceFloat, errCB := strconv.ParseFloat(consumerUser.EnergyAccountBalance, 64)
			if errCB != nil {
				fmt.Println("Error converting Consumer Balance")
				return nil, errors.New("Error converting Consumer Balance")
			}

			changeConBalanceFloat, errCCB := strconv.ParseFloat(con.ChangeInConsumerBalance, 64);
			if errCCB != nil {
				fmt.Println("Error converting Changed Consumer Balance")
				return nil, errors.New("Error converting Changed Consumer Balance")
			}

			changeConBalanceOldFloat, errCCB := strconv.ParseFloat(con.ChangeInConsumerBalanceOld, 64);
			if errCCB != nil {
				fmt.Println("Error converting Changed Consumer Old Balance")
				return nil, errors.New("Error converting Changed Consumer Old Balance")
			}
			consumerBalanceFloat = consumerBalanceFloat + changeConBalanceFloat - changeConBalanceOldFloat
			consumerUser.EnergyAccountBalance = strconv.FormatFloat(consumerBalanceFloat, 'f', 6, 64)

			// Consumer Balance Update ends here

			batteryUser,err := GetUsers(con.Battery.UserID+"_Battery", stub)
			if err != nil {
				fmt.Println("Failed to retrieve battery ")
			}
			con.Battery = batteryUser
			batteryBalanceFloat, errBB := strconv.ParseFloat(batteryUser.EnergyAccountBalance, 64)
			if errBB != nil {
				fmt.Println("Error converting Battery Balance")
				return nil, errors.New("Error converting Battery Balance")
			}

			changeBatteryBalanceFloat, errCBB := strconv.ParseFloat(con.ChangeInBatteryBalance, 64)
			if errCBB != nil {
				fmt.Println("Error converting Changed Battery Balance")
				return nil, errors.New("Error converting Changed Battery Balance")
			}

			changeBatteryBalanceOldFloat, errCBB := strconv.ParseFloat(con.ChangeInBatteryBalanceOld, 64)
			if errCBB != nil {
				fmt.Println("Error converting Changed Battery Old Balance")
				return nil, errors.New("Error converting Changed Battery Old Balance")
			}


			batteryBalanceFloat = batteryBalanceFloat + changeBatteryBalanceFloat - changeBatteryBalanceOldFloat
			batteryUser.EnergyAccountBalance = strconv.FormatFloat(batteryBalanceFloat, 'f', 6, 64)

			// Battery balance update ends here


			// Grid balance update starts here
			gridUser,err := GetUsers("0_Grid", stub)
			if err != nil {
				fmt.Println("Failed to retrieve grid ")
			}
			con.Grid = gridUser

			gridBalanceFloat, errGB := strconv.ParseFloat(gridUser.EnergyAccountBalance, 64)
			if errGB != nil {
				fmt.Println("Error converting Grid Balance")
				return nil, errors.New("Error converting Grid Balance")
			}

			changeGridBalanceFloat, errCPB := strconv.ParseFloat(con.ChangeInGridBalance, 64)
			if errCPB != nil {
				fmt.Println("Error converting Changed Grid Balance")
				return nil, errors.New("Error converting Changed Grid Balance")
			}

			changeGridBalanceOldFloat, errCPB := strconv.ParseFloat(con.ChangeInGridBalanceOld, 64)
			if errCPB != nil {
				fmt.Println("Error converting Changed Grid Old Balance")
				return nil, errors.New("Error converting Changed Grid Old Balance")
			}


			gridBalanceFloat = gridBalanceFloat + changeGridBalanceFloat - changeGridBalanceOldFloat
			gridUser.EnergyAccountBalance = strconv.FormatFloat(gridBalanceFloat, 'f', 6, 64)

			// Grid balance update ends here

			// Updating producer balance
			producerUpdate, err := json.Marshal(producerUser)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(producerUpdate))
			err = stub.PutState(con.Producer.UserID+"_Prosumer", []byte(string(producerUpdate)))
			if err != nil {
				fmt.Println("Failed to update producer ")
			}
			fmt.Println("Updated Producer  with Key : "+ con.Producer.UserID+"_Prosumer")

			// Updating consumer balance
			consumerUpdate, err := json.Marshal(consumerUser)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(consumerUpdate))
			err = stub.PutState(con.Consumer.UserID+"_Prosumer", []byte(string(consumerUpdate)))
			if err != nil {
				fmt.Println("Failed to update consumer ")
			}
			fmt.Println("Updated consumer  with Key : "+ con.Consumer.UserID+"_Prosumer")

			// Updating Battery balance
			batteryUpdate, err := json.Marshal(batteryUser)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(batteryUpdate))
			err = stub.PutState(con.Battery.UserID+"_Battery", []byte(string(batteryUpdate)))
			if err != nil {
				fmt.Println("Failed to update battery balance ")
			}
			fmt.Println("Updated battery balance  with Key : "+ con.Battery.UserID+"_Battery")

			// Updating Grid Balance
			gridUpdate, err := json.Marshal(gridUser)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(gridUpdate))
			err = stub.PutState("0_Grid", []byte(string(gridUpdate)))
			if err != nil {
				fmt.Println("Failed to update grid balance ")
			}
			fmt.Println("Updated grid balance  with Key : 0_Grid")


			// Updating the contract
			conUpdate, err := json.Marshal(con)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(conUpdate))
			err = stub.PutState(con.ContractID, []byte(string(conUpdate)))
			if err != nil {
				fmt.Println("Failed to update contract ")
			}
			fmt.Println("Updated Contract  with Key : "+ con.ContractID)


		}
		fmt.Println("In initialize.PerformSettlement end ")
		return nil,nil

	}


	func main() {
		err := shim.Start(new(EnergyTradingChainCode))
		if err != nil {
			fmt.Printf("Error starting Simple chaincode: %s", err)
		}


	}
