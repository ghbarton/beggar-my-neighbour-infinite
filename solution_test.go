package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateGame(t *testing.T) {
	testCases := []struct {
		name  string
		hand1 string
		hand2 string
		exp   int
	}{
		{
			name:  "game 1",
			hand1: "Q--J----K--K-J---Q---A---A",
			hand2: "---K-K--JA-QA--J-----Q----",
			exp:   5104,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := simulateGame(convertStringIntoGameArray(tc.hand1), convertStringIntoGameArray(tc.hand2))
			assert.Equal(t, tc.exp, audit)
		})
	}
}

func TestRunGame(t *testing.T) {
	testCases := []struct {
		name      string
		hand1     string
		hand2     string
		expTurns  int
		expTricks int
	}{
		{
			name:      "game 1",
			hand1:     "Q--J----K--K-J---Q---A---A",
			hand2:     "---K-K--JA-QA--J-----Q----",
			expTurns:  5105,
			expTricks: 713,
		},
		{
			name:      "game 1",
			hand1:     "---JQ---K-A----A-J-K---QK-",
			hand2:     "-J-----------AJQA----K---Q",
			expTurns:  5790,
			expTricks: 805,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			auditTurns, auditTricks := runGame(convertStringIntoGameArray(tc.hand1), convertStringIntoGameArray(tc.hand2))
			assert.Equal(t, tc.expTurns, auditTurns)
			assert.Equal(t, tc.expTricks, auditTricks)
		})
	}
}

func TestValueMapping(t *testing.T) {
	testCases := []struct {
		name string
		data string
		exp  []int
	}{
		{
			name: "values map correctly",
			data: "AKQJ-",
			exp:  []int{4, 3, 2, 1, 0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := convertStringIntoGameArray(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}

func TestGiveWinnerCards(t *testing.T) {
	testCases := []struct {
		name    string
		init    []int
		discard []int
		expInit []int
		expDisc []int
	}{
		{
			name:    "Give winner cards",
			init:    []int{0},
			discard: []int{1},
			expInit: []int{0, 1},
			expDisc: []int{},
		},
		{
			name:    "Give winner cards",
			init:    []int{},
			discard: []int{1},
			expInit: []int{1},
			expDisc: []int{},
		},
		{
			name:    "Give winner cards",
			init:    []int{0},
			discard: []int{1, 0, 0, 0, 0},
			expInit: []int{0, 1, 0, 0, 0, 0},
			expDisc: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initA, discardA := giveWinnerCards(tc.init, tc.discard)
			assert.Equal(t, tc.expInit, initA)
			assert.Equal(t, tc.expDisc, discardA)
		})
	}
}

func TestPlayTrick(t *testing.T) {
	testCases := []struct {
		name     string
		init     []int
		player   []int
		discard  []int
		expInit  []int
		expPlay  []int
		expDisc  []int
		errorExp error
	}{
		{
			name:     "Play trick",
			init:     []int{},
			player:   []int{0},
			discard:  []int{1},
			expInit:  []int{1, 0},
			expPlay:  []int{},
			expDisc:  []int{},
			errorExp: nil,
		},
		{
			name:     "Initiator looses",
			init:     []int{0, 0, 0, 0},
			player:   []int{1},
			discard:  []int{1},
			expInit:  []int{0, 0, 0},
			expPlay:  []int{1, 1, 0},
			expDisc:  []int{},
			errorExp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initA, playerA, discardA, err := playTrick(tc.init, tc.player, tc.discard)
			assert.Equal(t, tc.expInit, initA)
			assert.Equal(t, tc.expPlay, playerA)
			assert.Equal(t, tc.expDisc, discardA)
			assert.Equal(t, tc.errorExp, err)
		})
	}
}

func TestPlaceCard(t *testing.T) {
	testCases := []struct {
		name       string
		hand       []int
		discard    []int
		handExp    []int
		discardExp []int
		cardExp    int
		errorExp   error
	}{
		{
			name:       "Add card to discard",
			hand:       []int{0},
			discard:    []int{},
			discardExp: []int{0},
			handExp:    []int{},
			cardExp:    0,
			errorExp:   nil,
		},
		{
			name:       "Check correct order",
			hand:       []int{3, 2, 1, 0},
			discard:    []int{},
			discardExp: []int{3},
			handExp:    []int{2, 1, 0},
			cardExp:    3,
			errorExp:   nil,
		},
		{
			name:       "Check correct order",
			hand:       []int{3, 2, 1, 0},
			discard:    []int{0, 0, 0},
			discardExp: []int{0, 0, 0, 3},
			handExp:    []int{2, 1, 0},
			cardExp:    3,
			errorExp:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handA, discardA, err := placeCard(tc.hand, tc.discard)
			assert.Equal(t, tc.cardExp, discardA[len(discardA)-1])
			assert.Equal(t, tc.discardExp, discardA)
			assert.Equal(t, tc.handExp, handA)
			assert.Equal(t, tc.errorExp, err)
		})
	}
}
