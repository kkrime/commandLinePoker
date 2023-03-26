package main

import (
	"iceyePoker/poker"
	"log"
)

func main() {
	if err := poker.NewPokerGame().Play(); err != nil {
		log.Fatal(err)
	}
}
