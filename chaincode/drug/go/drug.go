 package main
 
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 type Patient struct {
	Id string `json:"Id"`
	RiskLevel string `json:"RiskLevel"`
	Prescriptions []Prescription	
 }
 type Prescription struct {
	Name string `json:"Name"`
	CreateDate string `json:"CreateDate"`
	ExpireDate string `json:"ExpireDate"`
	ControlledSubstance string `json:"ControlledSubstance"`
	Schedule string `json:"Schedule"`
	Dosage string `json:"Dosage"`
	Brand string `json:"Brand"`
	LastDispenseDate string `json:"LastDispenseDate"`
	NumberOfRefills string `json:"NumberOfRefills"`
	Phamacy string `json:"Phamacy"`
 }
 

 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 
 type Drug struct {
	PrescriptionName   string `json:"PrescriptionName"`
	PrescriptionDate  string `json:"PrescriptionDate"`
	Status string `json:"Status"`
	ControlledSubstance  string `json:"ControlledSubstance"`
 }

 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryDrug" {
		 return s.queryDrug(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createDrug" {
		 return s.createDrug(APIstub, args)
	 } else if function == "queryAllDrugs" {
		 return s.queryAllDrugs(APIstub)
	 } else if function == "changeDrugStatus" {
		 return s.changeDrugStatus(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryDrug(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 drugAsBytes, _ := APIstub.GetState(args[0])
	 fmt.Println();
	 return shim.Success(drugAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

 patients := []Patient{
	{
	Id:"111",
	RiskLevel:"2",
	Prescriptions:
	[]Prescription{
		{
	Name:"Hydrocodone",
	CreateDate:"08/01/2018",
	ExpireDate:"01/20/2019",
	ControlledSubstance:"Yes",
	Schedule:"Schedule II",
	Dosage:"50",
	Brand:"Generic",
	LastDispenseDate:"08/03/2018",
	NumberOfRefills:"1",
	Phamacy:"BostonMA",
		},
},
	},
	{
		Id:"111",
		RiskLevel:"2",
		Prescriptions:
		[]Prescription{
			{
		Name:"Simvastatin",
		CreateDate:"07/11/2018",
		ExpireDate:"12/20/2018",
		ControlledSubstance:"Yes",
		Schedule:"Schedule III",
		Dosage:"50",
		Brand:"Generic",
		LastDispenseDate:"08/03/2018",
		NumberOfRefills:"1",
		Phamacy:"BostonMA",
	},
	},
},
 }
	i := 0
	for i < len(patients) {
		fmt.Println("i is ", i)
		patientsAsBytes, _ := json.Marshal(patients[i])
		APIstub.PutState("DRUG"+strconv.Itoa(i), patientsAsBytes)
		fmt.Println("Added", patients[i])
		i = i + 1
	}
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createDrug(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	log.Println("Create Drug inside go file")
	var Patient = 	Patient{
		Id:args[1],
		RiskLevel:args[2],
		Prescriptions:
		[]Prescription{
			{
		Name:args[3],
		CreateDate:args[4],
		ExpireDate:args[5],
		ControlledSubstance:args[6],
		Schedule:args[7],
		Dosage:args[8],
		Brand:args[9],
		LastDispenseDate:args[10],
		NumberOfRefills:args[11],
		Phamacy:args[12],
			},
	},
		}
	 
	 PatientAsBytes, _ := json.Marshal(Patient)
	 APIstub.PutState(args[0], PatientAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryAllDrugs(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "DRUG0"
	 endKey := "DRUG99"

	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryAllDrugs:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) changeDrugStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 drugAsBytes, _ := APIstub.GetState(args[0])
	 drug := Drug{}
 
	 json.Unmarshal(drugAsBytes, &drug)
	 drug.Status = args[1]
 
	 drugAsBytes, _ = json.Marshal(drug)
	 APIstub.PutState(args[0], drugAsBytes)
 
	 return shim.Success(nil)
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 