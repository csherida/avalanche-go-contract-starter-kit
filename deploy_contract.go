package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/ethclient"
)

func deployContract(ctx context.Context, ec ethclient.Client, opts *bind.TransactOpts) string {
	// deploy the contract
	storageAddress, storageTransaction, _, err := DeployStorage(opts, ec)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	fmt.Println("storage address: ", storageAddress)

	// wait for the transaction to be accepted
	for {
		r, err := ec.TransactionReceipt(ctx, storageTransaction.Hash())
		if err != nil {
			if err.Error() != "not found" {
				log.Fatal(err)
			}
			time.Sleep(1 * time.Second)
			continue
		}
		if r.Status != 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return storageAddress.String()
}
