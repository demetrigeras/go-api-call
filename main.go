package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/tx/:txHash", getTransaction)
	r.Run()

}

func getTransaction(c *gin.Context) {
	// Extract the transaction hash from the URL parameter
	txHash := c.Param("txHash")

	// Connect to Ethereum client (Infura)
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/f1bfabaa66614342a34543701a76b373")
	if err != nil {
		log.Printf("Error fetching transaction: %v", err)
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}

	if !pending {
		// Render the transaction using a template or return as JSON
		c.JSON(200, tx)
	} else {
		c.JSON(404, gin.H{"error": "Transaction pending or not found"})
	}
}
