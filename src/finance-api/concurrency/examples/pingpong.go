package examples

import (
	"fmt"
	"time"
)

type Ball struct {
	hits    int
	players []string
}

func execute() *Ball {
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)

	time.Sleep(1 * time.Second)

	return <-table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table

		ball.hits++
		ball.players = append(ball.players, name)
		fmt.Print(name, ball.hits)
		time.Sleep(100 * time.Millisecond)

		table <- ball
	}
}
