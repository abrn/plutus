# plutus

a bitcoin and monero payment processing server/wallet with a built in REST API to query transactions, wallets and addresses  
this can be run on the clearnet or optionally as an onion service to ensure total anonymity between the server and the API caller  
  
with this 17mb binary, you can accept payments for btc/xmr without paying any fees to services like bitpay, meanwhile staying anonymous and in total control of your coin

## features

- "secure enclave" system using [memguard](https://github.com/awnumar/memguard) that stores sensitive information encrypted in memory without writing to the disk, such as private keys
- full support for multi-signature and time-locked bitcoin transactions
- REST API with an http server using [gin](https://github.com/gin-gonic/gin#benchmarks) (the most performant server you'll find on the internet, which is useful for the next point)
- ability to run the API as a hidden service on the Tor network
- API keys with optional permissions that can be remotely managed
- ability to lock the wallet and effectively turn it into cold storage
  
check out the `docker` directory to see example setups for each service required and the `docker-compose.yml` in the root which can tie them all together (this is the optimal setup for Tor)  
  
the docker files are built for debian/ubuntu linux but the compiled binary is compatible with any windows or unix system

## how it works

the binary acts as an interface between bitcoind, monerod and the respective wallets for each coin then exposes a REST API for interacting with them  
this allows you to create new wallets, addresses, invoices for receiving transactions or sending coin without any RPC interaction, meanwhile logging everything to a database  
  
plutus then handles everything in the meantime and makes sure the database matches what is held in each wallet. querying this database is lightning fast compared to querying the nodes
or wallet instances themselves, which can take up to 3 minutes to return a full balance report

an http API is exposed and the routes are below (more information in `router.go` and the `routes` folder):

### general  

**POST** `/revoke` - revoke an API key  
**GET**  `/health` - return the server health status  

### bitcoin  

**GET**  `/btc/balance` - get the total balance of all accounts  
**POST** `/btc/createaccount` - create a new account  
**POST** `/btc/listtransactions` - list all txs or txs from a certain address/account  
**POST** `/btc/sendtoaddress` - send coin to an address using the optimal coin inputs    
**POST** `/btc/sendfrom` - send coin from a certain address/account  
**POST** `/btc/sendmany` - send multiple transactions  
**POST** `/btc/move` - move coin between an internal address/account  
**POST** `/btc/getaccountaddress` - get a new address for the wallet or a given account  
**POST** `/btc/addmultisig` - add a multisig setup to a transaction before broadcasting  
**POST** `/btc/submit_pubkey` - add a new public key to sign a transaction with  
**POST** `/btc/unlock` - unlock the wallet from cold storage  

### monero  

**GET**  `/xmr/notify/:txid`  
**POST** `/xmr/createaccount`  
**POST** `/xmr/createaddress`  
**POST** `/xmr/getbalance`  
**POST** `/xmr/transfer`  
  
## requirements  

- `golang` to compile the server into a single binary
- `btcd` for a [full bitcoin node](https://github.com/btcsuite/btcd) (more lightweight than bitcoind)
- `monerod` to run the standard monero server
- `postgresql` to store the wallets/accounts, addresses and transaction history

more coming soon...
