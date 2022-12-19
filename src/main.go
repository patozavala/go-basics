package main

import "fmt"

func main() {
	// for-loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while-loop
	var counter int = 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// conditionals
	const btc = true
	const money = true
	const garbage = true
	const fiat = true

	if btc || money {
		fmt.Println("BTC is sound money")
	}

	if garbage && fiat {
		fmt.Println("FIAT money is garbage")
	}

	// Switch
	module := 5 % 2
	switch module {
	case 0:
		fmt.Println("is pair")
	default:
		fmt.Println("is odd")
	}

}
