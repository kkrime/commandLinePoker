package main

import (
	"iceyePoker/poker"
	"log"
)

const (
	validHandString = "2,3,4,5,6,7,8,9,T,J,Q,K,A"
)

func main() {

	if err := poker.NewPokerGame().Play(); err != nil {
		log.Fatal(err)
	}
}
