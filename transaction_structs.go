package transactionclient

// TransactionPayload represents the JSON payload for broadcasting a transaction.
type TransactionPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

// TransactionResponse represents the response from the transaction broadcast.
type TransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

// TransactionStatusResponse represents the response from checking the transaction status.
type TransactionStatusResponse struct {
	TxStatus string `json:"tx_status"`
}
