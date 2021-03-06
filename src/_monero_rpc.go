package plutus

// import (
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// 	"os"

// 	"github.com/awnumar/memguard"
// )

// // MoneroRPC contains information about an RPC call
// type MoneroRPC struct {
// 	JSONRPCVersion float32     `json:"jsonrpc"`
// 	ID             uint64      `json:"id"`
// 	Method         string      `json:"method"`
// 	Params         interface{} `json:"params"`
// }

// // we only need to interact with the monero wallet rpc api
// // so that's what will be defined here

// // CreateXMRAccountResponse contains information that will be returned from the create_account rpc call
// type CreateXMRAccountResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		AccountIndex int    `json:"account_index"`
// 		Address      string `json:"address"`
// 	} `json:"result"`
// }

// // CreateXMRAccount accepts a label (mandatory in our case), and returns a response or an error
// func CreateXMRAccount(labelEnclave *memguard.Enclave) (*memguard.Enclave, error) {
// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	b, err := labelEnclave.Open()
// 	defer b.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "create_account",
// 		Params:         map[string]string{"label": b.String()},
// 	}

// 	// marhsal the above rpc call into json format
// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		log.Printf("Err: %s\n", err.Error())
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(resp)
// 	memguard.WipeBytes(resp)

// 	return responseEnclave, nil

// }

// // CreateXMRAddressResponse will hold information regarding our CreateAddress RPC call
// type CreateXMRAddressResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		Address      string `json:"address"`
// 		AddressIndex int    `json:"address_index"`
// 	} `json:"result"`
// }

// // CreateXMRAddress will call the create_address RPC call for us
// func CreateXMRAddress(indexEnclave *memguard.Enclave) (*memguard.Enclave, error) {
// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	b, err := indexEnclave.Open()
// 	defer b.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "create_address",
// 		Params:         map[string]interface{}{"account_index": b.Bytes()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil
// }

// // GetXMRAddressResponse will hold information regarding the get_address rpc call
// type GetXMRAddressResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		Address   string `json:"address"`
// 		Addresses []struct {
// 			Address      string `json:"address"`
// 			AddressIndex int    `json:"address_index"`
// 			Label        string `json:"label"`
// 			Used         bool   `json:"used"`
// 		} `json:"addresses"`
// 	} `json:"result"`
// }

// // GetXMRAddress will call the get_address rpc call, and return ALL addresses for an account
// func GetXMRAddress(indexEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	b, err := indexEnclave.Open()
// 	defer b.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "get_address",
// 		Params:         map[string]interface{}{"account_index": b.Bytes()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil
// }

// // GetXMRBalanceResponse contains infomration from the get_balance rpc call
// type GetXMRBalanceResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		Balance              int64 `json:"balance"`
// 		MultisigImportNeeded bool  `json:"multisig_import_needed"`
// 		PerSubaddress        []struct {
// 			Address           string `json:"address"`
// 			AddressIndex      int    `json:"address_index"`
// 			Balance           int64  `json:"balance"`
// 			Label             string `json:"label"`
// 			NumUnspentOutputs int    `json:"num_unspent_outputs"`
// 			UnlockedBalance   int64  `json:"unlocked_balance"`
// 		} `json:"per_subaddress"`
// 		UnlockedBalance int64 `json:"unlocked_balance"`
// 	} `json:"result"`
// }

// // GetXMRBalance will call the get_balance rpc call for us
// func GetXMRBalance(indexEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	b, err := indexEnclave.Open()
// 	defer b.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "get_balance",
// 		Params:         map[string]interface{}{"account_index": b.Bytes()},
// 	}

// 	data, err := json.Marshal(&call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil
// }

// // TransferInput will hold information necessary for making the transfer rpc call
// type TransferInput struct {
// 	Destinations []struct {
// 		Amount  int64  `json:"amount"`
// 		Address string `json:"address"`
// 	} `json:"destinations"`
// 	AccountIndex uint `json:"account_index"`
// 	Priority     uint `json:"priority"`
// 	RingSize     uint `json:"ring_size"`
// }

// // TransferXMRResponse holds information about the response from the transfer rpc call
// type TransferXMRResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		Amount        int64  `json:"amount"`
// 		Fee           int64  `json:"fee"`
// 		MultisigTxset string `json:"multisig_txset"`
// 		TxBlob        string `json:"tx_blob"`
// 		TxHash        string `json:"tx_hash"`
// 		TxKey         string `json:"tx_key"`
// 		TxMetadata    string `json:"tx_metadata"`
// 		UnsignedTxset string `json:"unsigned_txset"`
// 	} `json:"result"`
// }

// // Transfer will call the transfer rpc call for us
// func Transfer(inputEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	input, err := inputEnclave.Open()
// 	defer input.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	params, err := json.Marshal(input.Bytes())
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "transfer",
// 		Params:         params,
// 	}

// 	data, err := json.Marshal(call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil
// }

// // GetTransfersResponse models the response returned by the get_transfers rpc call
// type GetTransfersResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		In []struct {
// 			Address         string `json:"address"`
// 			Amount          int64  `json:"amount"`
// 			Confirmations   int    `json:"confirmations"`
// 			DoubleSpendSeen bool   `json:"double_spend_seen"`
// 			Fee             int64  `json:"fee"`
// 			Height          int    `json:"height"`
// 			Note            string `json:"note"`
// 			PaymentID       string `json:"payment_id"`
// 			SubaddrIndex    struct {
// 				Major int `json:"major"`
// 				Minor int `json:"minor"`
// 			} `json:"subaddr_index"`
// 			SuggestedConfirmationsThreshold int    `json:"suggested_confirmations_threshold"`
// 			Timestamp                       int    `json:"timestamp"`
// 			Txid                            string `json:"txid"`
// 			Type                            string `json:"type"`
// 			UnlockTime                      int    `json:"unlock_time"`
// 		} `json:"in"`
// 		Out []struct {
// 			Address         string `json:"address"`
// 			Amount          int64  `json:"amount"`
// 			Confirmations   int    `json:"confirmations"`
// 			DoubleSpendSeen bool   `json:"double_spend_seen"`
// 			Fee             int64  `json:"fee"`
// 			Height          int    `json:"height"`
// 			Note            string `json:"note"`
// 			PaymentID       string `json:"payment_id"`
// 			SubaddrIndex    struct {
// 				Major int `json:"major"`
// 				Minor int `json:"minor"`
// 			} `json:"subaddr_index"`
// 			SuggestedConfirmationsThreshold int    `json:"suggested_confirmations_threshold"`
// 			Timestamp                       int    `json:"timestamp"`
// 			Txid                            string `json:"txid"`
// 			Type                            string `json:"type"`
// 			UnlockTime                      int    `json:"unlock_time"`
// 		} `json:"out"`
// 	} `json:"result"`
// }

// // GetXMRTransfers will accept an index and call the get_transfers rpc call
// func GetXMRTransfers(indexEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	index, err := indexEnclave.Open()
// 	defer index.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "get_transfers",
// 		Params:         map[string]interface{}{"in": true, "out": true, "account_index": index.String()},
// 	}

// 	data, err := json.Marshal(call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil

// }

// // GetTransferByTxIDResponse models the response returned by the get_transfer_by_txid rpc method
// type GetTransferByTxIDResponse struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  struct {
// 		Transfer struct {
// 			Address       string `json:"address"`
// 			Amount        int64  `json:"amount"`
// 			Confirmations int    `json:"confirmations"`
// 			Destinations  []struct {
// 				Address string `json:"address"`
// 				Amount  int64  `json:"amount"`
// 			} `json:"destinations"`
// 			DoubleSpendSeen bool   `json:"double_spend_seen"`
// 			Fee             int64  `json:"fee"`
// 			Height          int    `json:"height"`
// 			Note            string `json:"note"`
// 			PaymentID       string `json:"payment_id"`
// 			SubaddrIndex    struct {
// 				Major int `json:"major"`
// 				Minor int `json:"minor"`
// 			} `json:"subaddr_index"`
// 			SuggestedConfirmationsThreshold int    `json:"suggested_confirmations_threshold"`
// 			Timestamp                       int    `json:"timestamp"`
// 			Txid                            string `json:"txid"`
// 			Type                            string `json:"type"`
// 			UnlockTime                      int    `json:"unlock_time"`
// 		} `json:"transfer"`
// 	} `json:"result"`
// }

// // GetXMRTransferByTxID will call the get_transfer_by_txid rpc method
// func GetXMRTransferByTxID(txIDEnclave *memguard.Enclave) (*memguard.Enclave, error) {

// 	MoneroWalletRPCURL := os.Getenv("MONERO_WALLET_RPC_URL")

// 	txID, err := txIDEnclave.Open()
// 	defer txID.Destroy()
// 	if err != nil {
// 		return nil, err
// 	}

// 	call := &MoneroRPC{
// 		JSONRPCVersion: 2.0,
// 		ID:             uint64(rand.Int63()),
// 		Method:         "get_transfer_by_txid",
// 		Params:         map[string]interface{}{"txid": txID.String()},
// 	}

// 	data, err := json.Marshal(call)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rep, err := DirectJSONPOST(MoneroWalletRPCURL, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	responseEnclave := memguard.NewEnclave(rep)
// 	memguard.WipeBytes(rep)

// 	return responseEnclave, nil
// }
