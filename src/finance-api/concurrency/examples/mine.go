package examples

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Money struct {
	id        uuid.UUID
	value     float32
	createdBy string
}

type Wallet struct {
	money []Money
}

func (wallet *Wallet) GetTotal() float32 {
	if wallet.money == nil || len(wallet.money) == 0 {
		return 0
	}

	var total float32
	for _, money := range wallet.money {
		total += money.value
	}

	return total
}

func mine(n int) (total float32) {
	walletChan := make(chan *Wallet)

	go worker("Sam", walletChan)
	go worker("Mike", walletChan)
	go worker("Alice", walletChan)

	walletChan <- new(Wallet)

	time.Sleep(time.Second * time.Duration(1))

	wallet := <-walletChan

	total = wallet.GetTotal()
	return
}

func worker(name string, walletChan chan *Wallet) {
	for {
		w := <-walletChan

		time.Sleep(time.Millisecond * time.Duration(200))
		value := rand.Int31n(10000)

		if value < 5000 {
			walletChan <- w
			continue
		}

		money := Money{
			createdBy: name,
			value:     float32(value) / 100,
			id:        uuid.New(),
		}

		println("Mined:", money.value, " by:", name)
		w.money = append(w.money, money)

		walletChan <- w
	}
}
