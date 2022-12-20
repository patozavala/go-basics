package tokens

import (
	"fmt"
)

type IToken interface {
	Mint()
}

type Nft struct {
	Supply int
	Name   string
}

type Ft struct {
	Supply int
	Name   string
}

func (token *Nft) Mint() {
	var message string
	if token.Supply == 0 {
		token.Supply++
		message = fmt.Sprintf("%s. Supply : %d", token.Name, token.Supply)
	} else {
		message = fmt.Sprintln("NFT already exists")
	}
	fmt.Println(message)
}

func (token *Ft) Mint() {
	var message string
	token.Supply++
	message = fmt.Sprintf("%s. Supply : %d", token.Name, token.Supply)
	fmt.Println(message)
}

func Mint(token IToken) {
	token.Mint()
}
