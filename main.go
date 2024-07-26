package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

const (
	_subnetURL             = "http://127.0.0.1:9650/ext/bc/mySubnet/rpc"
	_prefundedAddress      = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027" // for address 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC
	_contractAddressString = "0xa1"                                                             // replace with your contract address once defined
	_defaultStoredAmount   = 1000
)

func main() {
	ctx := context.Background()

	// connect to subnet
	ec, err := ethclient.Dial(_subnetURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// capture the prefunded account's key
	prefundedKey, err := crypto.HexToECDSA(_prefundedAddress)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// setup transaction signature
	opts := createTransactionOptions(ctx, ec, prefundedKey)

	// deploy contract if it's not already
	contractAddress := _contractAddressString
	deployed := contractDeployed(ec, contractAddress)
	if !deployed {
		contractAddress = deployContract(ctx, ec, opts)
		fmt.Println("New contract deployed at ", contractAddress)
	}

	// connect to the contract instance from the chain
	address := common.HexToAddress(contractAddress)
	instance, err := NewStorage(address, ec)
	if err != nil {
		log.Fatalf("Failed to load the contract: %v", err)
	}

	txn, err := instance.Store(opts, convertToWei(_defaultStoredAmount))
	if err != nil {
		log.Fatalf("Failed to store transaction: %v", err)
	}

	fmt.Println("Stored transaction at address: ", txn.Hash().Hex())
	receipt, err := bind.WaitMined(ctx, ec, txn)
	if err != nil {
		log.Fatalf("Failed to mined the transaction: %v", err)
	}
	fmt.Println("Transaction mined at address: ", receipt.ContractAddress.Hex())

	storedAmt, err := instance.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve stored amount: %v", err)
	}

	fmt.Printf("Stored amount equals: %s\n", convertFromWei(storedAmt))
}

func createTransactionOptions(ctx context.Context, ec ethclient.Client, privateKeyECDSA *ecdsa.PrivateKey) *bind.TransactOpts {
	// fetch networkid
	networkId, err := ec.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// derive 'c' address
	cAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// setup signer and transaction options.
	signer := types.LatestSignerForChainID(networkId)
	opts := &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, privateKeyECDSA)
		},
		From:     cAddress,
		Context:  ctx,
		GasLimit: uint64(3000000),
		GasPrice: gasPrice,
	}
	return opts
}
