package txs

import "sync"

type Transaction struct {
	Token     string
	Amount    int
	Sender    string
	Receiver  string
	Timestamp int
}

type Mempool struct {
	Transactions []Transaction
}

func AddToMempool(tx Transaction, mp *Mempool, c chan<- *Mempool, wg *sync.WaitGroup) {
	defer wg.Done()
	mp.Transactions = append(mp.Transactions, tx)
	c <- mp
}
