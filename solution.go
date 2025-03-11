package main

import "strings"

func main() {}

func simulateGame(h1 []int, h2 []int) int {
	deckSize := len(h1) + len(h2)
	discard := []int{}
	isH1Turn := true
	for len(h1) < deckSize || len(h2) < deckSize {
		if isH1Turn {
			if len(h1) == 0 {
				return 1
			}
			h1, discard = placeCard(h1, discard)
			if discard[len(discard)-1] > 0 {
				h1, h2, discard = playTrick(h1, h2, discard)
			}
			isH1Turn = false
		} else {
			if len(h2) == 0 {
				return 2
			}
			h2, discard = placeCard(h2, discard)
			if discard[len(discard)-1] > 0 {
				h2, h1, discard = playTrick(h2, h1, discard)
			}
			isH1Turn = true
		}
	}
	return 0
}

func playTrick(initiator []int, player []int, discard []int) ([]int, []int, []int) {
	turns := discard[len(discard)-1] // top card
	for i := 0; i < turns; i++ {
		player, discard = placeCard(player, discard)
		if discard[len(discard)-1] > 0 {
			player, initiator, discard = playTrick(player, initiator, discard)
			return initiator, player, discard
		}
	}
	initiator, discard = giveWinnerCards(initiator, discard)
	return initiator, player, discard
}

func giveWinnerCards(initiator []int, discard []int) ([]int, []int) {
	return append(initiator, discard...), []int{}
}

// Hand[0] = top
// Discard[0] = bottom
func placeCard(hand []int, discard []int) ([]int, []int) {
	card, hand := hand[0], hand[1:]
	discard = append(discard, card)
	return hand, discard
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
