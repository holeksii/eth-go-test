package ethclient

import (
	"github.com/ethereum/go-ethereum/rpc"
)

func Connect() (*rpc.Client, error) {
	return rpc.Dial("ws://135.181.214.217:8546")
}
