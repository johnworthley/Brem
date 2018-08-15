package main

import (
	"log"

	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
	"math/big"
	"strconv"
	"time"
	"fmt"

	"../server/data"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	)

var   BREMcontractAddress	= ""
const etherscanApiKey       = "194M4NC49RXX52KMM1W5U91Y6UK2RM5WUW"
const intervalInMinutes		= 1
const etherscanApiUrl		= "api-rinkeby.etherscan.io"
const infuraApiUrl			= "rinkeby.infura.io"

type JsonTransaction struct{
	BlockNumber 		int64 	`json:"blockNumber"`
	TimeStamp			int64 	`json:"timeStamp"`
	Hash				string	`json:"hash"`
	Nonce				string	`json:"nonce"`
	BlockHash			string	`json:"blockHash"`
	TransactionIndex	int64	`json:"transactionIndex"`
	From				string	`json:"from"`
	To					string	`json:"to"`
	Value				string	`json:"value"`
	Gas					string	`json:"gas"`
	GasPrice			string	`json:"gasPrice"`
	IsError				string	`json:"isError"`
	TxReceipt_status	string	`json:"txReceipt_status"`
	Input				string	`json:"input"`
	ContractAddress		string	`json:"contractAddress"`
	CumulativeGasUsed	string	`json:"cumulativeGasUsed"`
	GasUsed				string	`json:"gasUsed"`
	Confirmations		int64	`json:"confirmations"`
}

type JsonListTransactions struct {
	Status		int64				`json:"status"`
	Message		string				`json:"message"`
	Result		[]JsonTransaction	`json:"result"`
}

type JsonNumLastBlock struct {
	Result				string	`json:"result"`
}
type JsonBlockTimeResult struct {
	Timestamp 			string	`json:"timestamp"`
}
type JsonBlockTime struct {
	Result 				JsonBlockTimeResult	`json:"result"`
}

func getAddressFromBytes32(bts string) string {
	for i, r := range bts{
		c := string(r)
		if c != "0" {
			bts = bts[i:]
			break
		}
	}

	if len(bts) == 40 {
		bts = "0x" + bts
	}
	if len(bts) == 41 {	//
		bts = "0" + bts
	}
	return bts
}

func getListTransactions(address string, startBlock int64, endBlock int64) JsonListTransactions {
	resp, err := http.Get("http://"+etherscanApiUrl+"/api?module=account&action=txlist&address="+address+"&startblock="+strconv.FormatInt(startBlock, 16)+"&endblock="+strconv.FormatInt(endBlock, 16)+"&sort=asc&apikey="+etherscanApiKey)
	var transactions JsonListTransactions
	if err != nil{
		log.Println(err)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal([]byte(contents), &transactions)

	}
	return transactions
}
func getFunctionByHashSignature(hashSignature string) string {
	switch hashSignature {
	case getHashSignatureFunction("addAuditor(address)"):
		return "addAuditor"
	case getHashSignatureFunction("addDeveloper()"):
		return "addDeveloper"
	case getHashSignatureFunction("signUp(string)"):
		return "signUp"
	case getHashSignatureFunction("createBREMICO(string,string,uint256,uint256,uint256,string,string)"):
		return "createBREMICO"
	}
	return ""
}
func getHashSignatureFunction(signatureFunction string) string {
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write([]byte(signatureFunction))
	buf = hash.Sum(buf)
	res := hex.EncodeToString(buf)
	return "0x" + res[0:8]
}

func BREMaddAuditor(auditorAddress string){

	var auditor data.Auditor

	auditor.Address = auditorAddress
	auditor.ID = -1
	err := auditor.GetAuditor()
	if err != nil {
		log.Println(err)
	}

	if auditor.ID == -1 {
		auditor.AddAuditor()
	}
}
func BREMaddDeveloper(devAddress string){

	var dev data.Developer

	dev.Address = devAddress
	dev.ID = -1
	err := dev.GetDeveloper()
	if err != nil {
		log.Println(err)
	}

	if dev.ID == -1 {
		dev.Create()
	}
}

func BREMICOaddAuditor(icoAddress string, auditorAddress string){

	var ico data.ICO

	ico.Address = icoAddress
	err := ico.GetICO()
	if err != nil {
		log.Println(err)
	}

	var auditor data.Auditor
	auditor.Address = auditorAddress
	err = auditor.GetAuditor()
	if err != nil {
		log.Println(err)
	}

	currentAuditors, _ := ico.GetICOAuditors()
	for i := range currentAuditors{
		if currentAuditors[i].Address == auditor.Address {
			return
		}
	}

	ico.AddAuditorToICO(auditor)

}

func BREMresolveTransaction(act JsonTransaction){

	fn := getFunctionByHashSignature(act.Input[0:10]);

	switch fn {
		case "addAuditor" :
			BREMaddAuditor(getAddressFromBytes32(act.Input[10:]))
//		case "addDeveloper" :
//			BREMaddDeveloper(getAddressFromBytes32(act.From))
//		case "signUp" :
//			BREMaddDeveloper(getAddressFromBytes32(act.From))
	}
}
func BREMICOresolveTransaction(act JsonTransaction){

	fn := getFunctionByHashSignature(act.Input[0:10]);

	switch fn{
		case "addAuditor" :
			BREMICOaddAuditor(getAddressFromBytes32(act.To), getAddressFromBytes32(act.Input[10:]))
	}
}

func updateBREMparticipants(startBlock int64){

	transactions := getListTransactions(BREMcontractAddress, startBlock, 99999999)

	for i := 0; i < len(transactions.Result); i++ {
		if transactions.Result[i].TxReceipt_status == "0" { // check failed transaction
			continue
		}
		if len(transactions.Result[i].Input) < 10 {
			continue
		}
		BREMresolveTransaction(transactions.Result[i])
	}
}
func updateAuditorsIco(icoAddress string, startBlock int64){

	transactions := getListTransactions(icoAddress, startBlock, 99999999)

	for i := 0; i < len(transactions.Result); i++ {
		if transactions.Result[i].TxReceipt_status == "0" { //check failed transaction
			continue
		}
		if len(transactions.Result[i].Input) < 10 {
			continue
		}
		BREMICOresolveTransaction(transactions.Result[i])
	}
}

func getBREMICOstatus(contract *BREMICO) string{

	var status string

	res, err := contract.AuditSelected(&bind.CallOpts{})
	if err != nil {
		log.Println(err)
	}

	if res == true {
		res, err := contract.HasClosed(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}

		if res == true {

			res, err := contract.IsOverdue(&bind.CallOpts{})
			if err != nil {
				log.Println(err)
			}

			if res == false {

				res, err := contract.CapReached(&bind.CallOpts{})
				if err != nil {
					log.Println(err)
				}

				if res == true {
					res, err := contract.IsRequested(&bind.CallOpts{})
					if err != nil {
						log.Println(err)
					}

					if res == false {
						res, err := contract.IsWithdrawn(&bind.CallOpts{})
						if err != nil {
							log.Println(err)
						}

						if res == false {
							status = "success"

						} else {
							status = "withdrawn"
						}

					} else {
						status = "requested"
					}

				} else {
					status = "failed"
				}

			} else {
				status = "overdue"
			}

		} else {
			status = "opened"
		}

	} else {
		status = "created"
	}

	return status
}
func checkIco(icoAddress string, client bind.ContractBackend, startBlock int64){

	contract, err := NewBREMICO(common.HexToAddress(icoAddress), client)
	if err != nil {
		log.Println(err)
	}

	var ico data.ICO
	ico.Address = icoAddress

	err = ico.GetICO()
	if err != nil {
		devAddress, err := contract.Wallet(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}

//		BREMaddDeveloper("0x" + hex.EncodeToString(devAddress[:]))

		var dev data.Developer
		dev.Address = "0x" + hex.EncodeToString(devAddress[:])
		err = dev.GetDeveloper()
		if err != nil {
			log.Println(err)
		}

		ico.Developer = dev

		closingTime, err := contract.ClosingTime(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}
		unixTime, err := strconv.Atoi(closingTime.String())
		if err != nil {
			log.Println(err)
		}
		ico.ClosingTime = time.Unix(int64(unixTime), 0) //

		feePercent, err := contract.WithdrawFeePercent(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}
		ico.FeePercent, err = strconv.Atoi(feePercent.String()) //

		tokenAddress, err := contract.Token(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}
		ico.TokenAddress = "0x" + hex.EncodeToString(tokenAddress[:])


		tokenContract, err := NewBREMToken(common.HexToAddress(ico.TokenAddress), client)
		if err != nil {
			log.Println(err)
		}

		name, err := tokenContract.Name(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}
		ico.Name = name

		symbol, err := tokenContract.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Println(err)
		}
		ico.Symbol = symbol

		ico.AddICO()
	}


	status := getBREMICOstatus(contract)	//



	res, err := contract.AuditorsAmount(&bind.CallOpts{})
	if err != nil {
		log.Println(err)
	}

	auditors, err := ico.GetICOAuditors()
	ln := int64(len(auditors))
	if res.Cmp(big.NewInt(ln)) > 0 {
		updateAuditorsIco(icoAddress, startBlock)
	}


	ico.SetStatusICO(status)

}
func updateBREMProjects(startBlock int64){

	client, err := ethclient.Dial("https://" + infuraApiUrl)
	if err != nil {
		log.Println(err)
	}

	BREMcontract, err := NewBREM(common.HexToAddress(BREMcontractAddress), client)
	if err != nil {
		log.Println(err)
	}

	projectsAmount, _ := BREMcontract.ProjectsAmount(&bind.CallOpts{})
	for i := big.NewInt(0); i.Cmp(projectsAmount) < 0; i.Add(i, big.NewInt(1)){

		BREMICOcontractAddress, _ := BREMcontract.GetProject(&bind.CallOpts{}, i)

		checkIco("0x" + hex.EncodeToString(BREMICOcontractAddress[:]), client, startBlock)
	}
}

func getNumLastBlock() int64 {
	resp, err := http.Get("https://"+etherscanApiUrl+"/api?module=proxy&action=eth_blockNumber&apikey="+etherscanApiKey)
	var numLastBlock JsonNumLastBlock
	if err != nil{
		log.Println(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal([]byte(contents), &numLastBlock)

	}

	if len(numLastBlock.Result) < 2 {
		log.Println("Error: len(numLastBlock.Result) < 2")
		return 0
	}
	result, err := strconv.ParseInt(numLastBlock.Result[2:], 16, 64)
	if err != nil{
		log.Println(err)
	}
	return result
}
func getTimeFromKBlock(K int64) int64 {
	resp, err := http.Get("https://"+etherscanApiUrl+"/api?module=proxy&action=eth_getBlockByNumber&tag="+strconv.FormatInt(K, 16)+"&boolean=true&apikey="+etherscanApiKey)
	var blockTime JsonBlockTime
	if err != nil{
		log.Println(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal([]byte(contents), &blockTime)
	}

	result, err := strconv.ParseInt(blockTime.Result.Timestamp[2:], 16, 64)
	if err != nil{
		log.Println(err)
	}
	return result
}
//get the block by age K minutes
func getKminuteBlock(K int) int64{
	const blocksPerMinute = 5

	currentTime := (time.Now()).Unix()
	wantedTime  := currentTime - int64( K * 60 )

	numLastBlock := getNumLastBlock()

	numBlock := numLastBlock - int64(int64(K)*int64(blocksPerMinute))
	if numLastBlock == 0 {
		numBlock = 0
	}
	time := getTimeFromKBlock(numBlock)
	if time < wantedTime {
		return numBlock
	}

	for {
		numBlock := numBlock - int64(int64(K/10)*int64(blocksPerMinute))
		time := getTimeFromKBlock(numBlock)
		if time < wantedTime {
			return numBlock
		}
	}

	return 0
}

type JsonBremAddress struct {
	Address			string	`json:"address"`
}
type JsonBremNetwork struct {
	Network			JsonBremAddress	`json:"4"`		//!!!
}
type JsonBremJson struct {
	Networks		JsonBremNetwork	`json:"networks"`
}
func getBREMcontractAddressFromBremJson() string {
	file, err := ioutil.ReadFile("./../build/contracts/BREM.json")
	if err != nil {
		log.Println(err)
	}
	var jsonBremJson JsonBremJson
	json.Unmarshal(file, &jsonBremJson)

	return jsonBremJson.Networks.Network.Address

}

func updateSuperuser(BREMAddress string){
	data.ClearSuperuser()

	transactions := getListTransactions(BREMAddress, 0, 99999999)

	var superuser data.Superuser

	superuser.Address = transactions.Result[0].From

	superuser.AddSuperuser()

}

func updateDB(minutes int){

	startBlock := getKminuteBlock(minutes)
	updateBREMparticipants(startBlock)
	updateBREMProjects(startBlock)

}

// DB updater
func RunUpdater() {

	BREMcontractAddress = getBREMcontractAddressFromBremJson()
	fmt.Println("BREMcontractAddress:" + BREMcontractAddress)

	updateSuperuser(BREMcontractAddress)

	updateBREMparticipants(0)
	updateBREMProjects(0)

	for {

		updateDB(intervalInMinutes * 3)

		time.Sleep(time.Minute * intervalInMinutes)

	}
}
