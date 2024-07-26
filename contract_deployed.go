package main

import (
	"context"
	"fmt"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func contractDeployed(ec ethclient.Client, cAddress string) bool {
	address := common.HexToAddress(cAddress)
	bytecode, err := ec.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatalf("Failed to retrieve contract bytecode: %v", err)
	}

	if len(bytecode) == 0 {
		fmt.Println("Contract is not deployed.")
		return false
	} else {
		fmt.Println("Contract is deployed.")
		return true
	}
}
