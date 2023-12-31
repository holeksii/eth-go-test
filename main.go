package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"holeksii.com/eth-go-test/db"
	"holeksii.com/eth-go-test/ethclient"
)

func main() {
	// db initialization
	db.InitDB()
	db.WaitDB()
	db.CreateTransactionsTable()
	defer db.CloseDB()

	// connect to eth
	nodeUrl := ethclient.GetNodeWsUrl()
	log.Printf("Connecting to Ethereum node: %s\n", nodeUrl)
	client, err := ethclient.Connect(nodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Subscribe to transactions
	ch := make(chan common.Hash)
	sub, err := client.EthSubscribe(context.Background(), ch, "newPendingTransactions")
	if err != nil {
		log.Fatalf("Failed to subscribe to pending transactions: %v", err)
	}
	defer sub.Unsubscribe()

	// Listen for transactions
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case txHash := <-ch:
			tx, err := ethclient.GetTransaction(client, txHash)
			timestamp := time.Now()
			if err != nil {
				log.Printf("Failed to fetch transaction details: %v", err)
				continue
			}

			// Check if transaction is matching
			matchingHash := "0x"
			if ethclient.IsMatchingTransaction(tx, matchingHash) {
				db.InsertTransaction(txHash, timestamp)
				log.Printf("Saved matching transaction hash: %s\n", tx.Hash)
			}
		}
	}
}
