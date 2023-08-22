package db

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func CreateTransactionsTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS transactions (hash VARCHAR(127) PRIMARY KEY, timestamp1 TIMESTAMP, timestamp2 TIMESTAMP)")
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	}
}

func InsertTransaction(txHash common.Hash, timestamp time.Time) {
	currentTime := time.Now()
	_, err := DB.Exec("INSERT INTO transactions (hash, timestamp1, timestamp2) VALUES ($1, $2, $3)", txHash.Hex(), timestamp, currentTime)
	if err != nil {
		log.Printf("Failed to save to database: %v", err)
	}
}
