package hand

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tj/assert"
)

type CompareHandTestSuite struct {
	suite.Suite
}

func TestCompareHandTestSuite(t *testing.T) {
	suite.Run(t, new(CompareHandTestSuite))
}

func (s *CompareHandTestSuite) TestcompareHand() {

	tests := []struct {
		name   string
		hand1  string
		hand2  string
		result result
	}{
		{
			name:   "Tie 1",
			hand1:  "AAAQQ",
			hand2:  "QQAAA",
			result: Tie,
		},
		{
			name:   "Tie 2",
			hand1:  "53QQ2",
			hand2:  "Q53Q2",
			result: Tie,
		},
		{
			name:   "Tie 3",
			hand1:  "53888",
			hand2:  "88385",
			result: Tie,
		},
		{
			name:   "Tie 4",
			hand1:  "QQAAA",
			hand2:  "AAAQQ",
			result: Tie,
		},
		{
			name:   "Tie 5",
			hand1:  "Q53Q2",
			hand2:  "53QQ2",
			result: Tie,
		},
		{
			name:   "Tie 6",
			hand1:  "88385",
			hand2:  "53888",
			result: Tie,
		},
		{
			name:   "Hand 1 wins 1",
			hand1:  "AAAQQ",
			hand2:  "QQQAA",
			result: Win,
		},
		{
			name:   "Hand 1 wins 2",
			hand1:  "Q53Q4",
			hand2:  "53QQ2",
			result: Win,
		},
		{
			name:   "Hand 1 wins 3",
			hand1:  "53888",
			hand2:  "88375",
			result: Win,
		},
		{
			name:   "Hand 1 wins 4",
			hand1:  "33337",
			hand2:  "QQAAA",
			result: Win,
		},
		{
			name:   "Hand 1 wins 5",
			hand1:  "22333",
			hand2:  "AAA58",
			result: Win,
		},
		{
			name:   "Hand 1 wins 6",
			hand1:  "33389",
			hand2:  "AAKK4",
			result: Win,
		},
		{
			name:   "Hand 1 wins 7",
			hand1:  "44223",
			hand2:  "AA892",
			result: Win,
		},
		{
			name:   "Hand 1 wins 8",
			hand1:  "22456",
			hand2:  "AKQJT",
			result: Win,
		},
		{
			name:   "Hand 1 wins 9",
			hand1:  "99977",
			hand2:  "77799",
			result: Win,
		},
		{
			name:   "Hand 1 wins 10",
			hand1:  "99922",
			hand2:  "88866",
			result: Win,
		},
		{
			name:   "Hand 1 wins 11",
			hand1:  "9922A",
			hand2:  "9922K",
			result: Win,
		},
		{
			name:   "Hand 1 wins 12",
			hand1:  "99975",
			hand2:  "99965",
			result: Win,
		},
		{
			name:   "Hand 1 wins 13",
			hand1:  "99975",
			hand2:  "99974",
			result: Win,
		},
		{
			name:   "Hand 1 wins 14",
			hand1:  "99752",
			hand2:  "99652",
			result: Win,
		},
		{
			name:   "Hand 1 wins 15",
			hand1:  "99752",
			hand2:  "99742",
			result: Win,
		},
		{
			name:   "Hand 1 wins 16",
			hand1:  "99753",
			hand2:  "99752",
			result: Win,
		},
		{
			name:   "Hand 1 wins 17",
			hand1:  "88822",
			hand2:  "QQ777",
			result: Win,
		},
		{
			name:   "Hand 1 wins 18",
			hand1:  "99662",
			hand2:  "88776",
			result: Win,
		},
		{
			name:   "Hand 2 wins 1",
			hand1:  "QQQAA",
			hand2:  "AAAQQ",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 2",
			hand1:  "53QQ2",
			hand2:  "Q53Q4",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 3",
			hand1:  "88375",
			hand2:  "53888",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 4",
			hand1:  "QQAAA",
			hand2:  "33337",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 5",
			hand1:  "AAA58",
			hand2:  "22333",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 6",
			hand1:  "AAKK4",
			hand2:  "33389",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 7",
			hand1:  "AA892",
			hand2:  "44223",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 8",
			hand1:  "AKQJT",
			hand2:  "22456",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 9",
			hand1:  "77799",
			hand2:  "99977",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 10",
			hand1:  "88866",
			hand2:  "99922",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 11",
			hand1:  "9922K",
			hand2:  "9922A",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 12",
			hand1:  "99965",
			hand2:  "99975",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 13",
			hand1:  "99974",
			hand2:  "99975",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 14",
			hand1:  "99652",
			hand2:  "99752",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 15",
			hand1:  "99742",
			hand2:  "99752",
			result: Lose,
		},
		{
			name:   "Hand 2 wins 15",
			hand1:  "99752",
			hand2:  "99753",
			result: Lose,
		},
	}
	for _, test := range tests {
		s.T().Run(test.name, func(t2 *testing.T) {

			hand1 := NewHand(test.hand1)
			hand2 := NewHand(test.hand2)

			result := hand1.CompareHand(hand2)

			assert.Equal(s.T(), test.result, result)
		})
	}
}
