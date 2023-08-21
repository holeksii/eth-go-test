package ethclient

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
)

func GetNodeWsUrl() string {
	jsonFile, err := os.Open("/root/config.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	var config map[string]interface{}
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&config)
	nodeUrl := config["config"].(map[string]interface{})["web3"].(map[string]interface{})["eth"].(map[string]interface{})["node"].(map[string]interface{})["urls"].(map[string]interface{})["ws"].(string)

	return nodeUrl
}

func Connect(url string) (*rpc.Client, error) {
	return rpc.Dial(url)
}
