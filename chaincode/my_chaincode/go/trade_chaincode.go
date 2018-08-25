package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// TradeFinanceChaincode example Trade Finance Chaincode implementation
type TradeFinanceChaincode struct {
}

//Init is called during chaincode instantiation to initialize any data. and chaincode upgrade also calls this function to reset
func (t *TradeFinanceChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the args from the transaction proposal
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 4 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	// Set up any variables or assets here by calling stub.PutState()
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)

}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The 'set'
// method may create a new asset by specifying a new key-value pair.
func (t *TradeFinanceChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "adms" {
		return t.adms(stub, args)
	}
	return shim.Success(nil)
}

func (t *TradeFinanceChaincode) adms(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	a := args[0]
	b := args[1]

	valA, err := stub.GetState(a)
	if err != nil {
		return shim.Error("Failed to get state of Value A ")
	}
	if valA == nil {
		return shim.Error("Entity not found")
	}

	valB, err := stub.GetState(b)
	if err != nil {
		return shim.Error("Failed to get state of Value B ")
	}
	if valB == nil {
		return shim.Error("Entity not found")
	}

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments.")
	}

	var addition, divition, substruction, multiple string

	valueofa, errvalueOfA := strconv.Atoi(args[2])
	if errvalueOfA != nil {
		return shim.Error("Invalid transaction amount of a")
	}

	valueofb, errvalueOfB := strconv.Atoi(args[3])
	if errvalueOfB != nil {
		return shim.Error("Invalid transaction amount of b")
	}

	add := valueofa + valueofb
	div := valueofa / valueofb
	sub := valueofa - valueofb
	mul := valueofa * valueofb

	err1 := stub.PutState(addition, []byte(strconv.Itoa(add)))
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	err2 := stub.PutState(substruction, []byte(strconv.Itoa(sub)))
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	err3 := stub.PutState(multiple, []byte(strconv.Itoa(mul)))
	if err3 != nil {
		return shim.Error(err3.Error())
	}
	err4 := stub.PutState(divition, []byte(strconv.Itoa(div)))
	if err3 != nil {
		return shim.Error(err4.Error())
	}

	return shim.Success(nil)

}
func main() {

	err := shim.Start(new(TradeFinanceChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
