package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

// HealthCheck will check that we can successfully connect to all necessary services. If we cannot, these errors will be logged
func HealthCheck(c *gin.Context) {

	var errored bool
	errors := make(map[string]string)

	dbConnError := plutus.Database.DB().Ping()
	if dbConnError != nil {
		plutus.Logger.Error("Cannot connect to database")
		errors["database"] = "Cannot connect to database"
		errored = true
	}

	// bitcoindConnError := plutus.Ping(os.Getenv("BTCD_HOST"))
	// if bitcoindConnError != nil {
	// 	plutus.Logger.Error("Cannot connect to bitcoind")
	// 	errors["bitcoind"] = "Cannot connect to bitcoind"
	// 	errored = true
	// }

	monerodConnError := plutus.Ping(os.Getenv("MONEROD_SERVICE_ADDR"))
	if monerodConnError != nil {
		plutus.Logger.Error("Cannot connect to monero_wallet_rpc")
		errors["monero_wallet_rpc"] = "Cannot connect to monero_wallet_rpc"
		errored = true
	}

	if errored {
		c.JSON(200, errors)
		return
	}

	c.JSON(200, "OK")
}
