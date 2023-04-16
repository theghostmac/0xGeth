package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/66c41c21577b45b0848e04dc31fe30d4"

func main() {
	// create the eth client.
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error creating ethereum client: %v", err)
	}
	defer client.Close()

	// access the block.
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error checking block number: %v", err)
	}
	fmt.Println("The block number is: ", block.Number())

	// addr is a random address that was gotten from Etherscan.io.
	walletAddress := "0xC771172AE08B5FC37B3AC3D445225928DE883876"
	convertedWalletAddress := common.HexToAddress(walletAddress) //HexToAddress converts the address to something processable.

	// getting balance.
	balance, err := client.BalanceAt(context.Background(), convertedWalletAddress, nil)
	if err != nil {
		log.Fatalf("Error getting the balance: %v", err)
	}
	fmt.Println("The balance is: ", balance)

	// convert the balance to correct Ether units
	correctBalance := new(big.Float)
	correctBalance.SetString(balance.String())
	fmt.Println("The correct balance is: ", correctBalance)

	// express balance in ether
}