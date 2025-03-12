package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {}

func simulateGame(h1 []int, h2 []int) int {
	deckSize := len(h1) + len(h2)
	discard := []int{}
	isH1Turn := true
	turns := 0
	var err error
	for len(h1) < deckSize || len(h2) < deckSize {
		turns++
		if isH1Turn {
			h1, discard, err = placeCard(h1, discard)
			if err != nil {
				return turns
			}
			if discard[len(discard)-1] > 0 {
				h1, h2, discard, err = playTrick(h1, h2, discard)
				if err != nil {
					return turns
				}
			}
			isH1Turn = false
		} else {
			h2, discard, err = placeCard(h2, discard)
			if err != nil {
				return turns
			}
			if discard[len(discard)-1] > 0 {
				h2, h1, discard, err = playTrick(h2, h1, discard)
				if err != nil {
					return turns
				}
			}
			isH1Turn = true
		}
	}
	return 0
}
func runGame(h1 []int, h2 []int) (int, int) {
	//deckSize := len(h1) + len(h2)
	discard := []int{}
	isH1Turn := true
	turns := 0
	tricks := 0
	var err error
	for len(h1) > 0 || len(h2) > 0 {
		cardsToPlay := 1
		battle := false
		for cardsToPlay > 0 {
			if isH1Turn {
				h1, discard, err = placeCard(h1, discard)
				if err != nil {
					return turns, tricks
					//break
				}
			} else {
				h2, discard, err = placeCard(h2, discard)
				if err != nil {
					return turns, tricks
					//break
				}
			}
			turns++
			if discard[len(discard)-1] == 0 {
				if battle {
					cardsToPlay--
				} else {
					isH1Turn = !isH1Turn
				}
			} else {
				battle = true
				cardsToPlay = discard[len(discard)-1]
				isH1Turn = !isH1Turn
			}
		}
		tricks++
		battle = false
		if isH1Turn {
			h2, discard = giveWinnerCards(h2, discard)
		} else {
			h1, discard = giveWinnerCards(h1, discard)
		}

		fmt.Printf("%v%v%v\n", h1, h2, discard)
		isH1Turn = !isH1Turn
	}
	return turns, tricks
}

func playTrick(initiator []int, player []int, discard []int) ([]int, []int, []int, error) {
	var err error
	turns := discard[len(discard)-1] // top card
	for i := 0; i < turns; i++ {
		player, discard, err = placeCard(player, discard)
		if err != nil {
			return initiator, player, discard, errors.New("player lost")
		}
		if discard[len(discard)-1] > 0 {
			player, initiator, discard, err = playTrick(player, initiator, discard)
			if err != nil {
				return initiator, player, discard, errors.New("player lost")
			}
			return initiator, player, discard, nil
		}
	}
	initiator, discard = giveWinnerCards(initiator, discard)
	return initiator, player, discard, nil
}

func giveWinnerCards(initiator []int, discard []int) ([]int, []int) {
	return append(initiator, discard...), []int{}
}

// Hand[0] = top
// Discard[0] = bottom
func placeCard(hand []int, discard []int) ([]int, []int, error) {
	if len(hand) == 0 {
		return hand, discard, errors.New("hand lost")
	}
	card, hand := hand[0], hand[1:]
	discard = append(discard, card)
	return hand, discard, nil
}

func convertStringIntoGameArray(x string) []int {
	cards := strings.Split(x, "")
	result := make([]int, len(cards))
	for i, card := range cards {
		switch card {
		case "A":
			result[i] = 4
		case "K":
			result[i] = 3
		case "Q":
			result[i] = 2
		case "J":
			result[i] = 1
		default:
			result[i] = 0
		}
	}
	return result
}
