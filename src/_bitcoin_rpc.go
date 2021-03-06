package plutus

// import (
// 	"encoding/json"
// 	"errors"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/awnumar/memguard"
// 	"github.com/tidwall/gjson"
// )

// // BitcoinRPC is similar to the MoneroRPC struct except that the parameters come in a different format, thus I've decided to create a new struct for it.
// type BitcoinRPC struct {
// 	JSONRPCVersion float32       `json:"jsonrpc"`
// 	ID             uint64        `json:"id"`
// 	Method         string        `json:"method"`
// 	Params         []interface{} `json:"params"`
// }

// // Used for reconciliation of funds

// // GetBTCBalance will call the getbalance rpc method
// func GetBTCBalance() ([]byte, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "getbalance",
// 		Params:         []interface{}{"*", 2},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return rep, nil
// }

// // ListTransactionsResponse models the response returned by the listtransactions rpc call
// type ListTransactionsResponse struct {
// 	Result []struct {
// 		Address           string  `json:"address"`
// 		Category          string  `json:"category"`
// 		Amount            float64 `json:"amount"`
// 		Label             string  `json:"label"`
// 		Vout              int     `json:"vout"`
// 		Fee               float64 `json:"fee"`
// 		Confirmations     int     `json:"confirmations"`
// 		Trusted           bool    `json:"trusted"`
// 		BlockHash         string  `json:"blockhash"`
// 		BlockIndex        int     `json:"blockindex"`
// 		BlockTime         int64   `json:"blocktime"`
// 		TxID              string  `json:"txid"`
// 		Time              int64   `json:"time"`
// 		TimeReceived      int64   `json:"timereceived"`
// 		Comment           string  `json:"comment"`
// 		BIP125Replaceable string  `json:"bip125-replaceable"`
// 		Abandoned         bool    `json:"abandoned"`
// 	}
// }

// // ListBTCTransactions will take a label and call the listtransactions rpc call
// func ListBTCTransactions(labelEnclave *memguard.Enclave, depth int) (*memguard.Enclave, error) {

// 	label, err := labelEnclave.Open()
// 	defer label.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "listtransactions",
// 		Params:         []interface{}{label.String(), depth},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	respEnclave := memguard.NewEnclave(rep)

// 	return respEnclave, nil
// }

// // GetTransactionResponse models the response returned by the gettransaction rpc call
// type GetTransactionResponse struct {
// 	Amount            float64 `json:"amount"`
// 	Fee               float64 `json:"fee"`
// 	Confirmations     uint    `json:"confirmations"`
// 	Blockhash         string  `json:"blockhash"`
// 	BlockIndex        uint    `json:"blockindex"`
// 	BlockTime         uint    `json:"blocktime"`
// 	TxID              string  `json:"txid"`
// 	Time              uint    `json:"time"`
// 	TimeReceived      uint    `json:"timereceived"`
// 	BIP125Replaceable string  `json:"bip-125-replaceable"`
// 	Details           []struct {
// 		Address   string  `json:"address"`
// 		Category  string  `json:"category"`
// 		Amount    float64 `json:"amount"`
// 		Label     string  `json:"label"`
// 		Vout      int     `json:"vout"`
// 		Fee       float64 `json:"fee"`
// 		Abandoned bool    `json:"abandoned"`
// 	}
// 	Hex     string      `json:"hex"`
// 	Decoded interface{} `json:"decoded"`
// }

// // GetBTCTransaction will call the gettransaction rpc call
// func GetBTCTransaction(txIDEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	txID, err := txIDEnclave.Open()
// 	defer txID.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "gettransaction",
// 		Params:         []interface{}{txID.String()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	respEnclave := memguard.NewEnclave(rep)

// 	return respEnclave, nil
// }

// // CreateRawTransactionResponse contains information regarding the payto electrum rpc call
// type CreateRawTransactionResponse struct {
// 	Transaction string `json:"transaction"`
// }

// // CreateRawTxInput models the structure of an input to the createrawtransaction rpc call
// type CreateRawTxInput struct {
// 	TxID string `json:"txid"`
// 	Vout int    `json:"vout"`
// }

// // CreateRawTxInputs is a collection of CreateRawTxInput structs
// type CreateRawTxInputs []CreateRawTxInput

// // CreateRawBTCTransaction will take a txid, address, and amount and call the createrawtransaction rpc call
// func CreateRawBTCTransaction(txidEnclave, addressEnclave, amountEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	txid, err := txidEnclave.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	address, err := addressEnclave.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	amount, err := amountEnclave.Open()
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer txid.Destroy()
// 	defer address.Destroy()
// 	defer amount.Destroy()

// 	input := CreateRawTxInputs{
// 		CreateRawTxInput{
// 			TxID: txid.String(),
// 			Vout: 0,
// 		},
// 	}

// 	inputJSON, err := json.Marshal(input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	outputJSON, err := json.Marshal([]interface{}{map[string]interface{}{address.String(): amount.String()}})
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "createrawtransaction",
// 		Params:         []interface{}{inputJSON, outputJSON},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // SignRawTransactionResponse will hold information from the response of the signtransaction rpc call
// type SignRawTransactionResponse struct {
// 	Hex      string `json:"hex"`      // The hex-encoded raw transaction with signature(s)
// 	Complete bool   `json:"complete"` // If the transaction has a complete set of signatures
// 	Errors   []struct {
// 		TxID      string `json:"txid"`      // The hash of the referenced, previous transaction
// 		Vout      int    `json:"vout"`      // The index of the output to spent and used as input
// 		ScriptSig string `json:"scriptSig"` // The hex-encoded signature script
// 		Sequence  int    `json:"sequence"`  //  Script sequence number
// 		Error     string `json:"error"`     // Verification or signing error related to the input
// 	} `json:"errors"`
// }

// // SignRawBTCTransactionWithWallet will call the signrawtransactionwithwallet rpc call
// func SignRawBTCTransactionWithWallet(txEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	tx, err := txEnclave.Open()
// 	defer tx.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "signrawtransactionwithwallet",
// 		Params:         []interface{}{tx.String()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // SignRawBTCTransactionWithKey will call the signrawtransactionwithkey rpc call
// func SignRawBTCTransactionWithKey(txEnclave, keyEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	tx, err := txEnclave.Open()
// 	defer tx.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	key, err := keyEnclave.Open()
// 	defer key.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "signrawtransactionwithkey",
// 		Params:         []interface{}{tx.String(), []string{key.String()}},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // SendRawTransactionResponse models the response returned by the signrawtransaction rpc call
// type SendRawTransactionResponse struct {
// 	Hex string `json:"hex"`
// }

// // SendRawBTCTransaction will call the sendrawtransaction rpc call
// func SendRawBTCTransaction(txEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("ELECTRUM_WALLET_RPC_URL")

// 	tx, err := txEnclave.Open()
// 	defer tx.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "sendrawtransaction",
// 		Params:         []interface{}{tx.String()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // SendToAddressResponse models the response returned by the sendtoaddress rpc method
// type SendToAddressResponse struct {
// 	Txid string `json:"txid"`
// }

// // SendBTCToAddress will call the sendtoaddress rpc method
// func SendBTCToAddress(addressEnclave, amountEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	address, err := addressEnclave.Open()
// 	defer address.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}
// 	amount, err := amountEnclave.Open()
// 	defer amount.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "sendtoaddress",
// 		Params:         []interface{}{address.String(), amount.String()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // GetNewBTCAddress will aceept a label and call the getnewaddress rpc call
// func GetNewBTCAddress(label string) ([]byte, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "getnewaddress",
// 		Params:         []interface{}{label},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return rep, nil
// }

// // AddMultisigAddressResponse will contain information from the createmultisig rpc call
// type AddMultisigAddressResponse struct {
// 	Address      string `json:"address"`
// 	RedeemScript string `json:"redeemScript"`
// }

// // AddMultisigBTCAddress will call the createmultisig rpc call
// func AddMultisigBTCAddress(merchant, customer *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	merchantBuf, err := merchant.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer merchantBuf.Destroy()

// 	merchantKey, err := FindPublicKeyByUsername(merchant)
// 	if err != nil {
// 		return nil, errors.New("Could not find public key for username: " + merchantBuf.String())
// 	}

// 	customerBuf, err := customer.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer customerBuf.Destroy()

// 	customerKey, err := FindPublicKeyByUsername(customer)
// 	if err != nil {
// 		return nil, errors.New("Could not find public key for username: " + customerBuf.String())
// 	}

// 	g, err := GetStaffKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	key, err := g.Open()
// 	defer key.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "addmultisigaddress",
// 		Params:         []interface{}{2, []string{merchantKey.PubKey, customerKey.PubKey, key.String()}, "multisig:" + merchantKey.Username + "," + customerKey.Username, "bech32"},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := plutus.DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	address := gjson.Get(string(rep), "address").String()
// 	redeemScript := gjson.Get(string(rep), "redeemScript").String()
// 	err = NewMultisigWallet([]byte(merchantKey.PubKey), []byte(customerKey.PubKey), key.Bytes(), address, redeemScript)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // BackupBTCWallet will call the backupwallet rpc. This will be exposed from the admin panel. It has no response, and simply copies the existing wallet file safely to the destination filepath. I believe this filepath can either be relative or absolute, but the documentation shows it being relative
// func BackupBTCWallet(destinationEnclave *memguard.Enclave) error {
// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	destination, err := destinationEnclave.Open()
// 	defer destination.Destroy()
// 	if err != nil {
// 		return err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "backupwallet",
// 		Params:         []interface{}{destination.String()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return err
// 	}

// 	// There's no response from this rpc so no need for a variable to hold it
// 	_, err = plutus.DirectJSONPOST(url, data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // LockUnspentResponse is just a bool, because that's all the lockunspent rpc method returns
// type LockUnspentResponse bool

// // LockUnspentBTC calls the lockunspent rpc method
// func LockUnspentBTC(txIDEnclave, vOutEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	tx, err := txIDEnclave.Open()
// 	defer tx.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	vOut, err := vOutEnclave.Open()
// 	defer vOut.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	vout, err := strconv.Atoi(vOut.String())
// 	if err != nil {
// 		return nil, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "lockunspent",
// 		Params:         []interface{}{false, tx.String(), vout},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := plutus.DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }

// // ListLockUnspentResponse models the response returned by the listlockunspent rpc method
// type ListLockUnspentResponse []struct {
// 	Txid string `json:"txid"`
// 	Vout int    `json:"vout"`
// }

// // ListLockedUnspentBTC will call the listlockunspent rpc method
// func ListLockedUnspentBTC() (*memguard.Enclave, error) {
// 	url := os.Getenv("BITCOIND_RPC_URL")

// 	rand.Seed(time.Now().UnixNano())
// 	call := &BitcoinRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             rand.Uint64(),
// 		Method:         "listlockunspent",
// 		Params:         []interface{}{},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := plutus.DirectJSONPOST(url, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)

// 	return responseEnclave, nil
// }
