package examples

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	if wallet == nil || wallet.money == nil || len(wallet.money) == 0 {
		return 0
	}

	var total float32
	for _, money := range wallet.money {
		total += money.value
	}

	return total
}

func (wallet *Wallet) String() string {
	if wallet == nil || wallet.money == nil || len(wallet.money) == 0 {
		return ""
	}

	var builder *strings.Builder = &strings.Builder{}
	builder.Grow(len(wallet.money) * (8 + 16 + 8))
	for _, money := range wallet.money {
		builder.WriteString(money.id.String())
		builder.WriteString(":")
		builder.WriteString(strconv.FormatFloat(float64(money.value), 'f', 6, 64))
		builder.WriteString(";\n")
	}

	return builder.String()
}

func mine(n int, workers []string) (total float32) {
	walletChan := make(chan *Wallet)

	rand.New(rand.NewSource(int64(time.Now().Second())))

	for _, work := range workers {
		//if only have one worker, then there will be not switch between goroutines, so the worker only mine once (and wait until timeout)
		go worker(work, walletChan)
	}

	walletChan <- new(Wallet)

	time.Sleep(time.Second * time.Duration(1))

	wallet := <-walletChan

	println(wallet.String())
	total = wallet.GetTotal()
	return
}

func worker(name string, walletChan chan *Wallet) {
	for {
		w := <-walletChan

		time.Sleep(time.Millisecond * time.Duration(200))
		value := rand.Int31n(10000)

		fmt.Println("Entered")
		if value < 1000 {
			walletChan <- w
			continue
		}

		money := Money{
			createdBy: name,
			value:     float32(value) / 100,
			id:        uuid.New(),
		}
		fmt.Printf("Mined: %f by: %s\n", money.value, name)
		w.money = append(w.money, money)

		walletChan <- w
	}
}
