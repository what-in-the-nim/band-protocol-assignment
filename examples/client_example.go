package main

import (
	"fmt"
	client "github.com/what-in-the-nim/transactionclient"
)

func main() {
	// Create some transaction data
	var symbol string = "ETH"
	var price uint64 = 4500

	// Broadcast the transaction
	fmt.Printf("Broadcasting transaction for %s at %d\n", symbol, price)
	txHash, err := client.BroadcastTransaction(symbol, price)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Confirmed Transaction Hash: %s\n", txHash)

	var status string
	// Check the transaction status
	status, _ = client.CheckTransactionStatus(txHash)
	fmt.Printf("Transaction Hash: %s\n", txHash)
	fmt.Printf("Transaction Status: %s\n", status)

	// Keep checking the transaction status every 5 seconds for 60 seconds
	status, _ = client.PollTransactionStatus(txHash, 5, 30)
	fmt.Printf("Transaction Hash: %s\n", txHash)
	fmt.Printf("Transaction Status: %s\n", status)
}
