package tokens

import "fmt"

// CryptoCurrency with public access
type CryptoToken struct {
	Ticker    string
	MaxSupply int
}

// CryptoCurrency with private access
type cryptoToken struct {
	ticker    string
	maxSupply int
}

func PrintTokenData(token CryptoToken) {
	fmt.Println(token)
}
