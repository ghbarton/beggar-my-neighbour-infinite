package main

import "strings"

func main() {}

func simulateGame(h1 []int, h2 []int) {
	deckSize := len(h1) + len(h2)
	discard := []int{}
	isH1Turn := true
	card := 0
	for len(h1) < deckSize || len(h2) < deckSize {
		if isH1Turn {
			card, h1, discard = placeCard(h1, discard)
			if card > 0 {
				playTrick(h1, h2, discard)
			}
			isH1Turn = false
		} else {
			card, h2, discard = placeCard(h2, discard)
			if card > 0 {
				playTrick(h2, h1, discard)
			}
			isH1Turn = true
		}
	}
}

func playTrick(initiator []int, player []int, discard []int) ([]int, []int, []int) {
	turns := discard[len(discard)-1]
	for i := 0; i < turns; i++ {

	}
}

// Hand[0] = top
// Discard[0] = bottom
func placeCard(hand []int, discard []int) (int, []int, []int) {
	card, hand := hand[0], hand[1:]
	discard = append(discard, card)
	return card, hand, discard
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
