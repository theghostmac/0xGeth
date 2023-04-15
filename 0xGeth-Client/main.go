package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var infuraURL = "https://mainnet.infura.io/v3/66c41c21577b45b0848e04dc31fe30d4"

func main() {
	// create the eth client
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error creating ethereum client: %v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error checking block number: %v", err)
	}
	fmt.Println(block.Number())
}

// cmd: curl --url https://mainnet.infura.io/v3/66c41c21577b45b0848e04dc31fe30d4 \
//-X POST \
//-H "Content-Type: application/json" \
//-d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
