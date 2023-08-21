package db

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func CreateTransactionsTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS transactions (hash VARCHAR(127) PRIMARY KEY, timestamp TIMESTAMP)")
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	}
}

func InsertTransaction(txHash common.Hash) {
	currentTime := time.Now()
	_, err := DB.Exec("INSERT INTO transactions (hash, timestamp) VALUES ($1, $2)", txHash.Hex(), currentTime)
	if err != nil {
		log.Printf("Failed to save to database: %v", err)
	}
}
