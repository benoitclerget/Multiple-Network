package main

// Note: external vendor lib must be added before deploying the smartcontract:
// $GOPATH/bin/govendor init
// $GOPATH/bin/govendor init
// $GOPATH/bin/govendor add +external

import (
	"encoding/base64"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("SmartContract demo-sc")

// Demo object
type Demo struct {
}

var params = map[string]int{
	"init1": 1,
	"hello": 1,
	"read":  1,
	"write": 2,
}

// Init ==> Chaincode init (called each time a smartcontrat is instantiated)
func (t *Demo) Init(stub shim.ChaincodeStubInterface) peer.Response {

	logger.Info("Initilization smart contract demo-sc")

	var err error

	var function, args = stub.GetFunctionAndParameters()

	switch function {
	case "init1":
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}
		err = stub.PutState("MyData", []byte(args[0]))
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success([]byte("Contrat demo-sc initialized"))
}

// Invoke ==> invoke functions (params array is used to identify the function)
func (t *Demo) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	//logger.Info("Invoke recu")

	// Verification des arguments
	var args = stub.GetArgs()

	var n, present = params[string(args[0])]

	//logger.Info("Function: ",string(args[0]))
	if present == false {
		return shim.Error("unkwnon function")
	}

	//logger.Info("Args length: ",len(args))
	if len(args)-1 != n {
		return shim.Error("Incorrect number of arguments. Expecting " + string(n))
	}

	// traitement de la function
	switch string(args[0]) {
	case "hello":
		return shim.Success([]byte("Hello World " + string(args[1])))
	case "read":
		id, _ := cid.GetID(stub)
		mspid, _ := cid.GetMSPID(stub)
		decodedid, _ := base64.StdEncoding.DecodeString(id)

		roleid, ok, err := cid.GetAttributeValue(stub, "role")
		if err != nil {
			return shim.Error("Error retrieving the role attribute")
			// There was an error trying to retrieve the attribute
		}
		Avalbytes, err := stub.GetState(string(args[1]))
		logger.Info(" read function, Key:  ", string(args[1]), " Value: ", string(Avalbytes), " Value bytes: ", Avalbytes)
		if err != nil {
			return shim.Error("Failed to get state")
		}
		//err = stub.SetEvent("READKEY", []byte(`{"key":"`+string(args[1])+`","value":"`+string(Avalbytes)+`"}`))
		err = stub.SetEvent("READKEY", []byte(`{"key":"`+string(args[1])+`","value":"NOT SHOWN"}`))
		if err != nil {
			return shim.Error(err.Error())
		}
		if !ok {
			// The client identity does not possess the role attribute
			return shim.Success([]byte(`{"key":"` + string(args[1]) + `","value":"` + string(Avalbytes) + `","userid":"` + string(decodedid) + `","orgid":"` + mspid + `","role":"unknown"}`))
		}
		return shim.Success([]byte(`{"key":"` + string(args[1]) + `","value":"` + string(Avalbytes) + `","userid":"` + string(decodedid) + `","orgid":"` + mspid + `","role":"` + roleid + `"}`))

	case "write":
		logger.Info(" write function, Key:  ", string(args[1]), " Value: ", string(args[2]))

		id, _ := cid.GetID(stub)
		mspid, _ := cid.GetMSPID(stub)

		roleid, ok, err := cid.GetAttributeValue(stub, "role")
		if err != nil {
			return shim.Error("Error retrieving the role attribute")
			// There was an error trying to retrieve the attribute
		}

		err = stub.PutState(string(args[1]), args[2])
		if err != nil {
			return shim.Error("Failed to write attribytes")
		}
		//err = stub.SetEvent("WRITEKEY", []byte(`{"key":"`+string(args[1])+`","value":"`+string(args[2])+`"}`))
		err = stub.SetEvent("WRITEKEY", []byte(`{"key":"`+string(args[1])+`","value":"NOT SHOWN"}`))
		if err != nil {
			return shim.Error(err.Error())
		}
		if !ok {
			// The client identity does not possess the role attribute
			return shim.Success([]byte(`{"key":"` + string(args[1]) + `","value":"` + string(args[2]) + `","userid":"` + id + `","orgid":"` + mspid + `","role":"unknown"}`))
		}
		return shim.Success([]byte(`{"key":"` + string(args[1]) + `","value":"` + string(args[2]) + `","userid":"` + id + `","orgid":"` + mspid + `","role":"` + roleid + `"}`))

	default:
		return shim.Success([]byte("Hello World "))
	}
}

func main() {
	err := shim.Start(new(Demo))

	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
