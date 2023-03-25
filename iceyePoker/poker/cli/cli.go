// package cli is used to manage interactions between user and the command line
package cli

import (
	"bufio"
	"fmt"
	"iceyePoker/poker/console"
	pokerErrors "iceyePoker/poker/errors"
	"os"
	"strings"
)

// validCardMap stores all the valid card values
var validCardMap = map[rune]bool{
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'T': true,
	'J': true,
	'Q': true,
	'K': true,
	'A': true,
}

// Cli interface
type Cli interface {
	ReadHand(message string) (string, error)
	AskBoolQuestion(message string) (bool, error)
}

type cli struct {
	reader *bufio.Reader
}

// NewCli returns a new Cli instance
func NewCli() Cli {
	return &cli{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ReadHand reads an input (the hand) from the cli
// The parameter message is printed before the user
// inputs the hand
func (c *cli) ReadHand(message string) (string, error) {
	// print message
	console.Green(message)
	hand, err := c.readInput()
	if err != nil {
		return "", err
	}
	// strip off '\n'
	hand = hand[:len(hand)-1]
	// make all letters capitals
	hand = strings.ToUpper(hand)
	// validate input
	err = validateHand(hand)
	if err != nil {
		return "", err
	}
	return hand, nil
}

// validateHand vaidates the hand
func validateHand(hand string) error {
	if len(hand) != 5 {
		return pokerErrors.NewUserInputError(wrongNumberOfCards)
	}

	cardMap := make(map[rune]bool)
	for _, c := range hand {
		if !validCardMap[c] {
			return pokerErrors.NewUserInputError(invalidInput)
		}
		cardMap[c] = true
	}

	if len(cardMap) == 1 {
		return pokerErrors.NewUserInputError(fiveCardsSameValue)
	}

	return nil
}

// AskBoolQuestion asks a yes or no question to the user
// The parameter question is displayed before the user inputs their answer
func (c *cli) AskBoolQuestion(question string) (bool, error) {
	for {
		// ask question
		console.Blue(question)
		fmt.Println("please enter either 'y' or 'n'")
		answer, err := c.readInput()
		if err != nil {
			return false, err
		}

		if answer[0] == 'n' {
			return false, nil
		} else if answer[0] == 'y' {
			return true, nil
		} else {
			// if invalid response ask again, until we get a valid response
			console.Error("Invalid response")
			fmt.Println()
			continue
		}
	}
}

// readInput reads an input from the user
// NOTE: this is a small function, and the go compiler should
// automatically detect this and compile it inline to improve performance
func (c *cli) readInput() (string, error) {
	// prompt
	fmt.Print("--> ")
	// read from stdIn
	return c.reader.ReadString('\n')
}
