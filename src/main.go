package main

import (
	"src/tokens"
)

func main() {

	cryptopunk := tokens.Nft{Name: "CryptoPunks"}
	clt := tokens.Ft{Name: "Cafinca Loyalty Token"}

	tokens.Mint(&cryptopunk)
	tokens.Mint(&clt)

}
