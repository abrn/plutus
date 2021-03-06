package main

import (
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/awnumar/memguard"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	parser := argparse.NewParser("plutus", "A bitcoin and monero capable payment server")
	flag := parser.Flag("g", "gen-api", &argparse.Options{Required: false, Help: "Generate an API key to be used by a client. The postgres service must be up to run this command", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Print(parser.Usage(err))
		memguard.SafeExit(1)
	}

	if *flag {
		key, err := plutus.NewAPIKey()
		if err != nil {
			log.Println(err)
			memguard.SafeExit(1)
		}

		log.Println("DO NOT UNDER ANY CIRCUMSTANCES SHARE THIS KEY WITH ANYONE. STORE IT IN A KEEPASSXC VAULT ON AN ENCRYPTED MEDIUM (SOMETHING WITH PLATTERS IF YOU CAN HELP IT)")
		log.Println("Key: " + key)

		memguard.SafeExit(0)
	}
	runServer()
}
