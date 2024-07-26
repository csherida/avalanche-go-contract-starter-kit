# avalanche-go-contract-starter-kit
Avalanche subnet example in Go with simple smart contract.

## Overview
This repo is a simple example of working with a smart contract with an Avalanche sub in Go.  Avalance is a derivative of ethereum but has extra headers in its protocols which means we should use the Avalanche client vs the standard geth client.

## Prerequisites
You will want to have the latest version of Go installed.  I built the initial version of this repo with go 1.22.2.

### Avalanche Subnet
This repo assumes you have created a local Avalanche subnet.  Follow these [directions](https://docs.avax.network/build/subnet/hello-subnet) and ensure a few things:  
* You call your subnet "mySubnet"
* Use 99999 as your ChainId
* Use TOKEN as your token name
* Select "Airdrop 1 million tokens to the default ewoq address (do not use in production)"

The ewoq address is well-known which means if you use it in production, someone will likely steal your tokens.  Only use this option for local subnets.

I already created and compiled a simple contract in the code.  If you wish to build your own smart contract with Go, you will need the solidity compiler and abigen utility.
### Solidity Compiler
Download the solidity compiler (assuming you have a Mac):
```
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
```
If you have a different machine or want an alternative installation, visit https://docs.soliditylang.org/en/latest/installing-solidity.html.

Once installed, you execute the following command from the project root for the sample contract in this repo:
```
solc --abi contracts/Storage.sol --bin -o build
```
This will create a folder called build and place the Storage.abi file and Storage.bin bytecode in it.

### Abigen Utility
This is a must-have utility that generates a go file to interact with your smart contract and the blockchain protocol.  It generates all the helper functions you need to deploy and access your Smart Contract functions.

To install, execute the following command:
```
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Alternatively, you can download and compile the code yourself:
```
git clone git@github.com:ava-labs/coreth.git
cd coreth/
go build -o abigen cmd/abigen/*.go
cp abigen ~/bin
```
Once installed, you run the following command for our sample contract:
```
abigen --abi build/Storage.abi --pkg main --type Storage --bin build/Storage.bin --out Storage.go
```
Providing the --bin argument will tell the utility to generate a deploy function (e.g. DeployStorage) that you can use to deploy the contract to the Avalanche subnet.

You can find out more about Abigen at https://geth.ethereum.org/docs/tools/abigen.