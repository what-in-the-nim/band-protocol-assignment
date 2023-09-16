# transactionclient

This Go package provides a client for broadcasting transactions and checking status on a remote service. It simplifies the process of creating transaction payloads, broadcasting transactions, and monitoring status.

## Table of Contents
- [transactionclient](#transactionclient)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Broadcasting a Transaction](#broadcasting-a-transaction)
    - [Checking Transaction Status](#checking-transaction-status)
    - [Polling Transaction Status](#polling-transaction-status)
  - [Example](#example)


## Installation

To use this in your project, import the `client` package into your project by adding the following import statement in your Go code:

```go
import client "github.com/what-in-the-nim/transactionclient"
```

Then, install the package using `go get`:

```go
go get -u github.com/what-in-the-nim/transactionclient"
```

## Usage

### Broadcasting a Transaction

To broadcast a transaction, use the `BroadcastTransaction` function. It takes the takes the following parameters:

- `symbol`: The symbol of the token to be transferred.
- `price`: The price of the token to be transferred.

Example:

```go
symbol := "BTC"
price := 1000

txHash, err := client.BroadcastTransaction(symbol, price)

if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Transaction Hash: %s\n", txHash)
}
```

### Checking Transaction Status

To check the status of a transaction, use the `CheckTransactionStatus` function. It takes the following parameters:

- `txHash`: The transaction hash to check.

Example:

```go
txHash := "your-transaction-hash"
status, err := client.CheckTransactionStatus(txHash)

if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Transaction Status: %s\n", status)
}
```

### Polling Transaction Status

To poll the status of a transaction, use the `PollTransactionStatus` function. It takes the following parameters:

- `txHash`: The transaction hash to check.
- `intervalSeconds`: The amount of time to wait between each check.
- `timeoutSeconds`: The amount of time to wait before timing out.

Example:

```go
txHash := "your-transaction-hash"
intervalSeconds := 5
timeoutSeconds := 60

status, err := client.PollTransactionStatus(txHash, intervalSeconds, timeoutSeconds)

if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Transaction Status: %s\n", status)
}
```

## Example

Example code can be found in the [examples](examples) directory.
