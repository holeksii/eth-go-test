package ethclient

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type Transaction struct {
	Hash  string `json:"hash"`
	Input string `json:"input"`
}

func GetTransaction(client *rpc.Client, txHash common.Hash) (*Transaction, error) {
	var tx Transaction
	err := client.Call(&tx, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func IsMatchingTransaction(tx *Transaction, matchingHash string) bool {
	return strings.Contains(tx.Input, matchingHash)
}
