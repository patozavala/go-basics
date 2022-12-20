package main

// we assume that there are many transactions that enter to a mempool to be validated

// disclaimer: this is not a realistic case, but it is a fun one :)

import (
	"fmt"
	"src/txs"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	c1 := make(chan *txs.Mempool, 1)
	c2 := make(chan *txs.Mempool, 1)

	tx1 := txs.Transaction{
		Token:  "BTC(SATS)",
		Amount: 1000000,
	}
	tx2 := txs.Transaction{
		Token:  "BTC(SATS)",
		Amount: 500000,
	}

	tx3 := txs.Transaction{
		Token:  "BTC(SATS)",
		Amount: 700000,
	}

	tx4 := txs.Transaction{
		Token:  "BTC(SATS)",
		Amount: 1100000,
	}

	var mempool txs.Mempool

	go txs.AddToMempool(tx1, &mempool, c1, &wg)
	wg.Add(1)
	fmt.Println("mempool", <-c1)

	go txs.AddToMempool(tx2, &mempool, c2, &wg)
	wg.Add(1)
	fmt.Println("mempool", <-c2)

	go txs.AddToMempool(tx3, &mempool, c1, &wg)
	wg.Add(1)
	fmt.Println("mempool", <-c1)

	go txs.AddToMempool(tx4, &mempool, c2, &wg)
	wg.Add(1)
	fmt.Println("mempool", <-c2)

	wg.Wait()

}
