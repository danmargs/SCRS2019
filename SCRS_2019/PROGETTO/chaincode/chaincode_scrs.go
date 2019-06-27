package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SCRSChaincode implementazione
type SCRSChaincode struct {
	Marito            string
	Moglie            string
	MaritoStatoCivile string
	MoglieStatoCivile string
	ContoInComune     string
}

//
// Init is the main function
//
func (t *SCRSChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Progetto Matrimonio Init")
	_, args := stub.GetFunctionAndParameters()
	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	// Initialize the chaincode
	t.Marito = args[0]
	if strings.Compare(args[1], "sposato") != 0 {
		return shim.Error("Devi inserire la stringa 'sposato' per la variabile Marito.")
	}
	t.MaritoStatoCivile = args[1]

	t.Moglie = args[2]
	if strings.Compare(args[3], "sposata") != 0 {
		return shim.Error("Devi inserire la stringa 'sposata' per la variabile Donna.")
	}
	t.MoglieStatoCivile = args[3]

	t.ContoInComune = args[4]

	fmt.Printf("MaritoStatoCivile = %s, MoglieStatoCivile = %s, ContoInComune = %s, ContoMarito = 0, ContoMoglie = 0\n", t.MaritoStatoCivile, t.MoglieStatoCivile, t.ContoInComune)

	// Write the state to the ledger
	err = stub.PutState(t.Marito, []byte(t.MaritoStatoCivile))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(t.Moglie, []byte(t.MoglieStatoCivile))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ContoInComune", []byte(t.ContoInComune))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ContoMarito", []byte("0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ContoMoglie", []byte("0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//
// Invoke can call the blockchain to outside
//
func (t *SCRSChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Progetto Matrimonio Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "aggiungi" {
		// richiama la funzione aggiungi
		return t.aggiungi(stub, args)
	} else if function == "divorzia" {
		// richiama la funzione divorzia
		return t.divorzia(stub, args)
	} else if function == "queryMarito" {
		// richiama la funzione queryMarito
		return t.queryMarito(stub, args)
	} else if function == "queryMoglie" {
		// richiama la funzione queryMoglie
		return t.queryMoglie(stub, args)
	} else if function == "aggiungiMarito" {
		// richiama la funzione aggiungiMarito
		return t.aggiungiMarito(stub, args)
	} else if function == "aggiungiMoglie" {
		// richiama la funzione aggiungiMoglie
		return t.aggiungiMoglie(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"aggiungi\" \"aggiungiMarito\" \"aggiungiMoglie\" \"divorzia\" \"query\"")
}

//
// aggiunge soldi al ContoMarito solo se ha già chiesto il divorzio
//
func (t *SCRSChaincode) aggiungiMarito(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ContoMarito string
	var X string

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	t.Marito = args[0]

	Maritovalbytes, err := stub.GetState(t.Marito)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Maritovalbytes == nil {
		return shim.Error("Entity not found")
	}
	t.MaritoStatoCivile = string(Maritovalbytes)

	if t.MaritoStatoCivile != "divorzio" {
		return shim.Error("Devi aver richiesto il divorzio per poter usufruire del conto separato")
	}

	X = args[1]

	ContoMaritovalbytes, err := stub.GetState("ContoMarito")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if ContoMaritovalbytes == nil {
		return shim.Error("Entity not found")
	}

	Somma, err := strconv.Atoi(X)
	Conto, err := strconv.Atoi(string(ContoMaritovalbytes))

	if Somma < 0 {
		if Conto < Somma {
			return shim.Error("Impossibile completare operazione, Conto inferiore alla somma da voler sottrarre.")
		}
	}

	Conto = Conto + Somma

	ContoMarito = strconv.Itoa(Conto)

	err = stub.PutState("ContoMarito", []byte(ContoMarito))
	if err != nil {
		return shim.Error(err.Error())
	}

	Risposta := "ContoMarito = " + ContoMarito
	fmt.Printf("ContoMarito = %s\n", ContoMarito)

	return shim.Success([]byte(Risposta))
}

//
// aggiunge soldi al ContoMoglie solo se ha già chiesto il divorzio
//
func (t *SCRSChaincode) aggiungiMoglie(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ContoMoglie string
	var X string

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	t.Moglie = args[0]

	Moglievalbytes, err := stub.GetState(t.Moglie)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Moglievalbytes == nil {
		return shim.Error("Entity not found")
	}
	t.MoglieStatoCivile = string(Moglievalbytes)

	if t.MoglieStatoCivile != "divorzio" {
		return shim.Error("Devi aver richiesto il divorzio per poter usufruire del conto separato")
	}

	X = args[1]

	ContoMoglievalbytes, err := stub.GetState("ContoMoglie")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if ContoMoglievalbytes == nil {
		return shim.Error("Entity not found")
	}

	Somma, err := strconv.Atoi(X)
	Conto, err := strconv.Atoi(string(ContoMoglievalbytes))

	if Somma < 0 {
		if Conto < Somma {
			return shim.Error("Impossibile completare operazione, Conto inferiore alla somma da voler sottrarre.")
		}
	}

	Conto = Conto + Somma

	ContoMoglie = strconv.Itoa(Conto)

	err = stub.PutState("ContoMoglie", []byte(ContoMoglie))
	if err != nil {
		return shim.Error(err.Error())
	}

	Risposta := "ContoMoglie = " + ContoMoglie
	fmt.Printf("ContoMoglie = %s\n", ContoMoglie)

	return shim.Success([]byte(Risposta))
}

//
// Consente di aggiungere soldi al conto se sia Marito che Moglie sono dichiarati come sposati
//
func (t *SCRSChaincode) aggiungi(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var X string // Transaction value
	var Conto int
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	t.Marito = args[0]
	t.Moglie = args[1]

	// Get the state from the ledger
	Maritovalbytes, err := stub.GetState(t.Marito)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Maritovalbytes == nil {
		return shim.Error("Entity not found")
	}
	t.MaritoStatoCivile = string(Maritovalbytes)

	Moglievalbytes, err := stub.GetState(t.Moglie)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Moglievalbytes == nil {
		return shim.Error("Entity not found")
	}
	t.MoglieStatoCivile = string(Moglievalbytes)

	if strings.Compare(t.MaritoStatoCivile, "sposato") != 0 || strings.Compare(t.MoglieStatoCivile, "sposata") != 0 {
		return shim.Error("Invalid transaction, tutte e due devono essere sposati")
	}

	X = args[2]

	Contovalbytes, err := stub.GetState("ContoInComune")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Contovalbytes == nil {
		return shim.Error("Entity not found")
	}

	Somma, err := strconv.Atoi(X)
	Conto, err = strconv.Atoi(string(Contovalbytes))

	if Somma < 0 {
		if Conto < Somma {
			return shim.Error("Impossibile completare operazione, Conto inferiore alla somma da voler sottrarre.")
		}
	}

	Conto = Conto + Somma

	t.ContoInComune = strconv.Itoa(Conto)

	fmt.Printf("MaritoStatoCivile = %s, MoglieStatoCivile = %s, ContoInComune = %s\n", t.MaritoStatoCivile, t.MoglieStatoCivile, t.ContoInComune)

	// Write the state back to the ledger
	err = stub.PutState(t.Marito, []byte(t.MaritoStatoCivile))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(t.Moglie, []byte(t.MoglieStatoCivile))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ContoInComune", []byte(t.ContoInComune))
	if err != nil {
		return shim.Error(err.Error())
	}

	Risposta := "MaritoStatoCivile = " + t.MaritoStatoCivile + ", MoglieStatoCivile = " + t.MoglieStatoCivile + ", ContoInComune = " + t.ContoInComune

	return shim.Success([]byte(Risposta))
}

// Fa cambiare lo stato civile di una controparte da sposato a divorziato
// (A,B) -> A divorzia da B (importante l'ordine)
func (t *SCRSChaincode) divorzia(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ConiugeA := args[0]
	ConiugeB := args[1]

	err := stub.PutState(ConiugeA, []byte("divorzio"))
	if err != nil {
		return shim.Error(err.Error())
	}

	//
	// controllo se tutte e due adesso sono 'divorizio'				*1*
	// e in caso prendo il ContoInComune, lo divido in 2 			*2*
	// lo metto metà in ContoMarito e metà in ContoMoglie			*2*
	// azzero ContoInComune											*3*
	// faccio putState di tutto quanto								*4*
	//

	ConiugeAvalbytes, err := stub.GetState(ConiugeA)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if ConiugeAvalbytes == nil {
		return shim.Error("Entity not found")
	}
	ConiugeAStatoCivile := string(ConiugeAvalbytes)

	//

	ConiugeBvalbytes, err := stub.GetState(ConiugeB)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if ConiugeBvalbytes == nil {
		return shim.Error("Entity not found")
	}
	ConiugeBStatoCivile := string(ConiugeBvalbytes)

	// *1*
	if ConiugeAStatoCivile == ConiugeBStatoCivile {

		Contovalbytes, err := stub.GetState("ContoInComune")
		if err != nil {
			return shim.Error("Failed to get state")
		}
		if Contovalbytes == nil {
			return shim.Error("Entity not found")
		}
		Conto, err := strconv.Atoi(string(Contovalbytes))

		//
		ContoMaritovalbytes, err := stub.GetState("ContoMarito")
		if err != nil {
			return shim.Error("Failed to get state")
		}
		if ContoMaritovalbytes == nil {
			return shim.Error("Entity not found")
		}
		ContoMa, err := strconv.Atoi(string(ContoMaritovalbytes))

		//
		ContoMoglievalbytes, err := stub.GetState("ContoMoglie")
		if err != nil {
			return shim.Error("Failed to get state")
		}
		if ContoMoglievalbytes == nil {
			return shim.Error("Entity not found")
		}
		ContoMo, err := strconv.Atoi(string(ContoMoglievalbytes))

		// *2*
		ContoMo = ContoMo + (Conto / 2)
		ContoMa = ContoMa + (Conto / 2)

		// *3*
		Conto = 0

		ContoInComune := strconv.Itoa(Conto)
		ContoMarito := strconv.Itoa(ContoMa)
		ContoMoglie := strconv.Itoa(ContoMo)

		// *4*
		err = stub.PutState("ContoInComune", []byte(ContoInComune))
		if err != nil {
			return shim.Error(err.Error())
		}
		err = stub.PutState("ContoMarito", []byte(ContoMarito))
		if err != nil {
			return shim.Error(err.Error())
		}
		err = stub.PutState("ContoMoglie", []byte(ContoMoglie))
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	fmt.Printf("ConiugeAStatoCivile = %s, ConiugeBStatoCivile = %s\n", ConiugeAStatoCivile, ConiugeBStatoCivile)
	Risposta := "ConiugeAStatoCivile = " + ConiugeAStatoCivile + ", ConiugeBStatoCivile = " + ConiugeBStatoCivile

	return shim.Success([]byte(Risposta))
}

//
// query callback representing the query of a chaincode
// per Marito ritorna tutte le informazioni
//
func (t *SCRSChaincode) queryMarito(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	t.Marito = args[0]

	// Get the state from the ledger
	Maritovalbytes, err := stub.GetState(t.Marito)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + t.Marito + "\"}"
		return shim.Error(jsonResp)
	}

	if Maritovalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for " + t.Marito + "\"}"
		return shim.Error(jsonResp)
	}

	ContoInComunevalbytes, err := stub.GetState("ContoInComune")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for ContoInComune\"}"
		return shim.Error(jsonResp)
	}

	if ContoInComunevalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for ContoInComune\"}"
		return shim.Error(jsonResp)
	}

	ContoMaritovalbytes, err := stub.GetState("ContoMarito")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for ContoMarito \"}"
		return shim.Error(jsonResp)
	}

	if ContoMaritovalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for ContoMarito\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + t.Marito + "\",\"Stato\":\"" + string(Maritovalbytes) + "\",\"ContoComune\":\"" + string(ContoInComunevalbytes) + "\",\"ContoSeparato\":\"" + string(ContoMaritovalbytes) + "\"}"
	fmt.Printf("QueryMarito Response:%s\n", jsonResp)
	Risposta := "Nome = " + t.Marito + ", Stato = " + t.MaritoStatoCivile + ", ContoComune = " + t.ContoInComune + ", ContoSeparato = " + string(ContoMaritovalbytes)

	return shim.Success([]byte(Risposta))
}

//
// query callback representing the query of a chaincode
// per Moglie ritorna tutte le informazioni
//
func (t *SCRSChaincode) queryMoglie(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	t.Moglie = args[0]

	// Get the state from the ledger
	Moglievalbytes, err := stub.GetState(t.Moglie)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + t.Moglie + "\"}"
		return shim.Error(jsonResp)
	}

	if Moglievalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for " + t.Moglie + "\"}"
		return shim.Error(jsonResp)
	}

	ContoInComunevalbytes, err := stub.GetState("ContoInComune")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for ContoInComune\"}"
		return shim.Error(jsonResp)
	}

	if ContoInComunevalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for ContoInComune\"}"
		return shim.Error(jsonResp)
	}

	ContoMoglievalbytes, err := stub.GetState("ContoMoglie")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for ContoMoglie \"}"
		return shim.Error(jsonResp)
	}

	if ContoMoglievalbytes == nil {
		jsonResp := "{\"Error\":\"Nil for ContoMoglie\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + t.Moglie + "\",\"Stato\":\"" + string(Moglievalbytes) + "\",\"ContoComune\":\"" + string(ContoInComunevalbytes) + "\",\"ContoSeparato\":\"" + string(ContoMoglievalbytes) + "\"}"
	fmt.Printf("QueryMoglie Response:%s\n", jsonResp)

	Risposta := "Nome = " + t.Moglie + ", Stato = " + t.MoglieStatoCivile + ", ContoComune = " + t.ContoInComune + ", ContoSeparato = " + string(ContoMoglievalbytes)

	return shim.Success([]byte(Risposta))
}

func main() {
	err := shim.Start(new(SCRSChaincode))
	if err != nil {
		fmt.Printf("Error starting SCRS chaincode: %s", err)
	}
}
