package transactionclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// NewTransactionPayload creates a new transaction payload.
func NewTransactionPayload(symbol string, price uint64) TransactionPayload {
	// Get the current timestamp
	currentTimeStamp := uint64(time.Now().Unix())

	// Return the payload
	return TransactionPayload{
		Symbol:    symbol,
		Price:     price,
		Timestamp: currentTimeStamp,
	}
}

// BroadcastTransaction sends a POST request to broadcast a transaction and returns the transaction hash.
func BroadcastTransaction(symbol string, price uint64) (string, error) {
	// Create a new transaction payload
	payload := NewTransactionPayload(symbol, price)

	// Serialize the payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Send the POST request
	resp, err := http.Post(broadcastURL, "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("broadcast failed with status code: %d", resp.StatusCode)
	}

	// Decode the response
	var response TransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// Return the transaction hash
	return response.TxHash, nil
}

// CheckTransactionStatus sends a GET request to check the status of a transaction and returns the status.
func CheckTransactionStatus(txHash string) (string, error) {
	// Send the GET request to check the status
	resp, err := http.Get(fmt.Sprintf(checkStatusURL, txHash))

	// Check for errors
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status check failed with status code: %d", resp.StatusCode)
	}

	// Decode the response
	var response TransactionStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// Return the transaction status
	return response.TxStatus, nil
}

// CheckTransactionStatus periodically checks the status of a transaction and returns the status.
func PollTransactionStatus(txHash string, intervalSeconds uint64, timeoutSeconds uint64) (string, error) {
	// Loop forever
	for {
		// Check the transaction status
		status, err := CheckTransactionStatus(txHash)

		// Check for transaction errors
		if err != nil {
			fmt.Printf("Error checking transaction status: %v\n", err)
		}

		// If the transaction is not pending, return
		if status != "PENDING" {
			return status, err
		}

		// Sleep for a while before checking again
		time.Sleep(time.Duration(intervalSeconds) * time.Second)

		// Decrement the timeout
		timeoutSeconds -= intervalSeconds

		// Check if the timeout has expired
		if timeoutSeconds <= 0 {
			return status, err
		}
	}
}
