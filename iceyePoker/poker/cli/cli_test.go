package cli

import (
	"testing"

	pokerErrors "iceyePoker/poker/errors"

	"github.com/stretchr/testify/suite"
	"github.com/tj/assert"
)

type CliTestSuite struct {
	suite.Suite
}

func TestCliTestSuite(t *testing.T) {
	suite.Run(t, new(CliTestSuite))
}

func (s *CliTestSuite) TestvalidateHand() {

	tests := []struct {
		name string
		hand string
		err  error
	}{
		{
			name: "Happy Path 1 - valid hand",

			hand: "AAA33",
		},
		{
			name: "Error - not enough cards",

			hand: "AAA5",

			err: pokerErrors.NewUserInputError(wrongNumberOfCards),
		},
		{
			name: "Error - too many cards",

			hand: "AAA337",

			err: pokerErrors.NewUserInputError(wrongNumberOfCards),
		},
		{
			name: "Error - invalid input",

			hand: "AAAAY",

			err: pokerErrors.NewUserInputError(invalidInput),
		},
		{
			name: "Error - same card 5 times",

			hand: "AAAAA",

			err: pokerErrors.NewUserInputError(fiveCardsSameValue),
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t2 *testing.T) {

			err := validateHand(test.hand)

			assert.Equal(s.T(), test.err, err)
		})
	}
}
