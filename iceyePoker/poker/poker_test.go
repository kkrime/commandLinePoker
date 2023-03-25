package poker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tj/assert"

	"iceyePoker/poker/cli"
	cliMock "iceyePoker/poker/cli/mocks"
)

type PokerTestSuite struct {
	suite.Suite
}

func TestPokerTestSuite(t *testing.T) {
	suite.Run(t, new(PokerTestSuite))
}

func (s *PokerTestSuite) Testplay() {

	tests := []struct {
		name   string
		cli    cli.Cli
		result string
		err    error
	}{
		{
			name: "Happy Path 1 - hand 1 wins",

			// MOCK
			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA33",
						nil,
					).
					Once()

				return cli

			}(),

			result: hand1Wins,
		},
		{
			name: "Happy Path 2 - hand 2 wins",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA66",
						nil,
					).
					Once()

				return cli

			}(),

			result: hand2Wins,
		},
		{
			name: "Happy Path 3 - tie",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				return cli

			}(),

			result: tie,
		},
		{
			name: "Error reading first hand from cli",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"",
						fmt.Errorf("error reading first hand"),
					).
					Once()

				return cli

			}(),

			err: fmt.Errorf("error reading first hand"),
		},
		{
			name: "Error reading second hand from cli",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"",
						fmt.Errorf("error reading second hand"),
					).
					Once()

				return cli

			}(),

			err: fmt.Errorf("error reading second hand"),
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t2 *testing.T) {

			pokerGame := newPokerGame(test.cli)

			result, err := pokerGame.play()

			assert.Equal(s.T(), test.result, result)
			assert.Equal(s.T(), test.err, err)
		})
	}
}

func (s *PokerTestSuite) Testplay_() {

	tests := []struct {
		name   string
		cli    cli.Cli
		result string
		err    error
	}{
		{
			name: "Happy Path 1 - hand 1 wins",

			// MOCK
			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA33",
						nil,
					).
					Once()

				return cli

			}(),

			result: hand1Wins,
		},
		{
			name: "Happy Path 2 - hand 2 wins",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA66",
						nil,
					).
					Once()

				return cli

			}(),

			result: hand2Wins,
		},
		{
			name: "Happy Path 3 - tie",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				return cli

			}(),

			result: tie,
		},
		{
			name: "Error reading first hand from cli",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"",
						fmt.Errorf("error reading first hand"),
					).
					Once()

				return cli

			}(),

			err: fmt.Errorf("error reading first hand"),
		},
		{
			name: "Error reading second hand from cli",

			cli: func() cli.Cli {
				cli := cliMock.NewCli(s.T())

				cli.On(
					"ReadHand",
					enterFirstHand,
				).
					Return(
						"AAA44",
						nil,
					).
					Once()

				cli.On(
					"ReadHand",
					enterSecondHand,
				).
					Return(
						"",
						fmt.Errorf("error reading second hand"),
					).
					Once()

				return cli

			}(),

			err: fmt.Errorf("error reading second hand"),
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t2 *testing.T) {

			pokerGame := newPokerGame(test.cli)

			result, err := pokerGame.play_()

			assert.Equal(s.T(), test.result, result)
			assert.Equal(s.T(), test.err, err)
		})
	}
}
