// package poker is used to facilitate a game of poker
package poker

import (
	"errors"
	"fmt"
	"iceyePoker/poker/cli"
	"iceyePoker/poker/console"
	pokerErrors "iceyePoker/poker/errors"
	"iceyePoker/poker/hand"
	"sync"
)

// poker interface
type Poker interface {
	Play() error
}

type poker struct {
	cli cli.Cli
}

// NewPokerGame is used to create a new instance of Poker
func NewPokerGame() Poker {
	cli := cli.NewCli()
	return newPokerGame(cli)
}

// Added for testing
func newPokerGame(cli cli.Cli) *poker {
	return &poker{
		cli: cli,
	}
}

// Play starts a game of poker
func (p *poker) Play() error {
	printWelcomeMessage()

	for {
		// play
		result, err := p.play()
		// result, err := p.play_()
		if err != nil {
			// check if user input error
			var userInputError *pokerErrors.UserInputError
			if errors.As(err, &userInputError) {
				console.Error(userInputError.Error())
				continue
			} else {
				return err
			}
		}

		// print out result
		printResult(result)

		// ask user if they would like to play again?
		playAgain, err := p.cli.AskBoolQuestion("Play another game?")
		if err != nil {
			return err
		}

		// user does not want to play again
		if !playAgain {
			printExitMessage()
			return nil
		}
	}
}

// play interacts with the cli to get poker hands and determins the winner
// NOTE: named returns are used here becuase otherwise,
// it would not be obvious from (string, error) that the string is the result
// NOTE: I could have just made a 'result' type, but this way, I can showcase more
// of my golang knowledge :)
func (p *poker) play() (result string, err error) {

	var hand1 *hand.Hand
	var hand2 *hand.Hand
	var usreInput string

	// print empty line
	fmt.Println()

	usreInput, err = p.cli.ReadHand(enterFirstHand)
	if err != nil {
		return
	}

	// use a Channel to detect when the first hand finished being constructed
	// NOTE: because there is a possibility that an error might occur during
	// ReadHand() below (line 108), there is a chance, that this function can return before
	// createHand1Chann is read. Therefore, to prevent a hanging goroutine and
	// a memory leak, createHand1Chann is a buffered channel with a length of 1
	// and NOT a unbuffered channel
	createHand1Chann := make(chan struct{}, 1)

	// construct hand1 on a seperate goroutine while the user enters hand2
	// to increase performance
	// NOTE: we must pass in userInput, beucase we use it again later in this fuction,
	// otherwise its value might change before the goroutine below runs
	go func(input string) {
		hand1 = hand.NewHand(input)
		createHand1Chann <- struct{}{}
	}(usreInput)

	usreInput, err = p.cli.ReadHand(enterSecondHand)
	if err != nil {
		return
	}

	hand2 = hand.NewHand(usreInput)

	// make sure hand1 has finished being created
	<-createHand1Chann

	// hand 1 vs hand 2
	res := hand1.CompareHand(hand2)
	if res == hand.Win {
		result = hand1Wins
	} else if res == hand.Lose {
		result = hand2Wins
	} else {
		result = tie
	}
	return
}

// play_ is the same as play above, but in this implementaotin,
// a Wait Group is used instead of a Channel
func (p *poker) play_() (result string, err error) {

	var hand1 *hand.Hand
	var hand2 *hand.Hand
	// use a Wait Group to detect when the first hand has finished been constructed
	var wg sync.WaitGroup
	var usreInput string

	// print empty line
	fmt.Println()

	usreInput, err = p.cli.ReadHand(enterFirstHand)
	if err != nil {
		return
	}

	// construct hand1 on a seperate goroutine while the user enters hand2
	// to increase performance
	// NOTE: we must pass in userInput, beucase we use it again later in this fuction,
	// otherwise its value might change before the goroutine below runs
	wg.Add(1)
	go func(input string) {
		defer wg.Done()
		hand1 = hand.NewHand(input)
	}(usreInput)

	usreInput, err = p.cli.ReadHand(enterSecondHand)
	if err != nil {
		return
	}

	hand2 = hand.NewHand(usreInput)

	// make sure hand1 has finished being created
	wg.Wait()

	// hand 1 vs hand 2
	res := hand1.CompareHand(hand2)
	if res == hand.Win {
		result = hand1Wins
	} else if res == hand.Lose {
		result = hand2Wins
	} else {
		result = tie
	}
	return
}

// printWelcomeMessage prints a welcome message
func printWelcomeMessage() {
	fmt.Println()
	console.BringWhite("Welcome To ICEYEPoker, The Best Command Line Poker Game Known To Humanity!!")
}

// printExitMessage prints a exit messgae
func printExitMessage() {
	fmt.Println()
	console.BringWhite("Thanks for playing! Come back soon :)")
}

// printResult prints the result
func printResult(result string) {
	fmt.Println()
	console.Yellow(result)
	fmt.Println()
}
