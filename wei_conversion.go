package main

import (
	"math/big"
	"strconv"
)

func convertToWei(amount int64) *big.Int {
	tokenAmount := new(big.Int)
	tokenAmount.SetString(strconv.FormatInt(amount, 10)+"000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination
	return tokenAmount
}

func convertFromWei(tokenAmount *big.Int) string {
	// Convert the balance to a human-readable format
	readableBalance := new(big.Float).Quo(new(big.Float).SetInt(tokenAmount), big.NewFloat(1e18))
	return readableBalance.String()
}
