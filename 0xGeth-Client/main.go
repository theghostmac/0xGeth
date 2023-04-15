package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/66c41c21577b45b0848e04dc31fe30d4"
var ganacheURL = "http://localhost:8545"

func main() {
	// create the eth client
	client, err := ethclient.DialContext(context.Background(), ganacheURL)
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
