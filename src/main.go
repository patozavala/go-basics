package main

import (
	"src/tokens"
)

func main() {

	var bitcoin tokens.CryptoToken

	bitcoin.Ticker = "BTC"
	bitcoin.MaxSupply = 21000000
	tokens.PrintTokenData(bitcoin)

	cardano := tokens.CryptoToken{
		Ticker:    "ADA",
		MaxSupply: 45000000000,
	}

	tokens.PrintTokenData(cardano)

}
