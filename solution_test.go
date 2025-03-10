package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
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

func TestPlaceCard(t *testing.T) {
	testCases := []struct {
		name       string
		hand       []int
		discard    []int
		handExp    []int
		discardExp []int
		cardExp    int
	}{
		{
			name:       "Add card to discard",
			hand:       []int{0},
			discard:    []int{},
			discardExp: []int{0},
			handExp:    []int{},
			cardExp:    0,
		},
		{
			name:       "Check correct order",
			hand:       []int{3, 2, 1, 0},
			discard:    []int{},
			discardExp: []int{3},
			handExp:    []int{2, 1, 0},
			cardExp:    3,
		},
		{
			name:       "Check correct order",
			hand:       []int{3, 2, 1, 0},
			discard:    []int{0, 0, 0},
			discardExp: []int{0, 0, 0, 3},
			handExp:    []int{2, 1, 0},
			cardExp:    3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cardA, handA, discardA := placeCard(tc.hand, tc.discard)
			assert.Equal(t, tc.cardExp, cardA)
			assert.Equal(t, tc.discardExp, discardA)
			assert.Equal(t, tc.handExp, handA)
		})
	}
}
