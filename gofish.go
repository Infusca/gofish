package main

import (
	"bufio"
	"fmt"
	// "math"
	"math/rand"
	"os"
	// "reflect"
	"strconv"
	"strings"
)

type Card struct {
	value int
	suit  int // 0-spade, 1-heart, 2-diamond, 3-club
}

func (card Card) getString() string {
	var suit string
	var value string

	switch card.suit {
	case 0:
		suit = "♠"
	case 1:
		suit = "♥"
	case 2:
		suit = "♦"
	case 3:
		suit = "♣"
	}

	switch card.value {
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	case 1:
		value = "A"
	default:
		value = fmt.Sprintf("%d", card.value)
	}

	return value + suit
}

type Deck struct {
	cards []Card
}

func (d *Deck) deal(num uint) []Card {
	dealtCards := []Card{}

	for i := uint(0); i < num; i++ {
		card := d.cards[0]
		dealtCards = append(dealtCards, card)
		d.cards = d.cards[1:] // remove first element (card) from slice (deck)
	}

	return dealtCards
}

func (d *Deck) create() {
	for i := 1; i <= 13; i++ {
		for j := 0; j < 4; j++ {
			card := Card{value: i, suit: j}
			d.cards = append(d.cards, card)
		}
	}
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

type Game struct {
	deck          Deck
	playerCards   []Card
	opponentCards []Card
	playerBooks   []Card
	opponentBooks []Card
}

func (game *Game) dealStartingCards() {
	game.deck.create()
	game.deck.shuffle()
	game.playerCards = game.deck.deal(7)
	game.opponentCards = game.deck.deal(7)

	fmt.Printf("Player has been dealt: \n")
	displayCards(game.playerCards)
	fmt.Printf("Opponent has been dealt: \n")
	displayCards(game.opponentCards)
}

func (game *Game) play() {
	defer fmt.Printf("\n end-of-play ------------------\n\n")
	fmt.Printf("\n start-of-play ------------------\n\n")

	// game logic
	if game.playerTurn() {

	}
	if game.opponentTurn() {

	}

}

// func (game *Game) searchCards(bait int) int {
// 	// bait, _ = strconv.ParseInt(bait, 10, 64)
// 	temp := []Card{}
// 	ask := bait
// 	removeCard := 0
// 	for i := 0; i < len(game.playerCards); i++ {
// 		// fish := int64(game.playerCards[i].value)

// 		// if bait == fish {
// 		// 	// add to card to opponent hand
// 		// 	game.opponentCards = append(game.opponentCards, game.playerCards[i])
// 		// 	// displayCards(game.playerCards)

// 		// 	// remove from player hand
// 		// 	for j := 0; j < len(game.playerCards); j++ {
// 		// 		card := game.playerCards[j].value
// 		// 		if int64(card) == bait && removeCard == 0 {
// 		// 			removeCard += 1
// 		// 		} else {
// 		// 			temp = append(temp, game.playerCards[j])
// 		// 		}
// 		// 	}
// 		// 	game.playerCards = temp

// 		// 	// continue playing if there are still cards in deck
// 		// 	if len(game.playerCards)+len(game.opponentCards) != 52 {
// 		// 		game.play()
// 		// 	} else {
// 		// 		// return false
// 		// 	}

// 		// }
// 	}
// 	fmt.Print(temp, ask, removeCard)
// 	return 5
// }

// func (game *Game) removeCardFromDeck(cards) {
// 	temp := []Card{}
// 	removeCard := 0

// 	for j := 0; j < len(cards); j++ {
// 		card := cards[j].value
// 		if int64(card) == bait && removeCard == 0 {
// 			removeCard += 1
// 		} else {
// 			temp = append(temp, game.opponentCards[j])
// 		}
// 	}
// 	game.opponentCards = temp
// }

func (game *Game) checkForBook(hand []Card) {
	count := 0
	for i := 0; i < len(game.playerCards); i++ {
		// card := game.playerCards[i]
		count++
	}
}

func (game *Game) playerTurn() bool {
	for true {
		fmt.Printf("What would you like to fish for? ")
		gofish := enterString()
		bait, _ := strconv.ParseInt(gofish, 10, 64)
		temp := []Card{}
		removeCard := 0

		// check opponent cards, if match, give to player, remove from opponent, return true
		for i := 0; i < len(game.opponentCards); i++ {
			fish := int64(game.opponentCards[i].value)

			if bait == fish {
				// add to card to player hand
				fmt.Printf("You caught a fish! \n")
				game.playerCards = append(game.playerCards, game.opponentCards[i])
				displayCards(game.playerCards)

				// remove from opponent hand
				for j := 0; j < len(game.opponentCards); j++ {
					card := game.opponentCards[j].value
					if int64(card) == bait && removeCard == 0 {
						removeCard += 1
					} else {
						temp = append(temp, game.opponentCards[j])
					}
				}
				game.opponentCards = temp
				displayCards(game.opponentCards)

				// check hand for book, remove from hand, add to player books
				game.checkForBook(game.playerCards)

				// continue playing if there are still cards in deck
				if len(game.playerCards)+len(game.opponentCards) != 52 {
					game.play()
				} else {
					return false
				}

			}
		}

		fmt.Printf("Go fish! \n")
		card := game.deck.deal(1)[0]
		game.playerCards = append(game.playerCards, card)
		displayCards(game.playerCards)
		displayCards(game.opponentCards)
		return false
	}
	return false
}

func (game *Game) opponentTurn() bool {
	// has4 := []Card{}
	// has3 := []Card{}
	// has2 := []Card{}
	// has1 := []Card{}
	// ask := []Card{}

	// determine how many of each card value opponent has
	// for i := 0; i < len(game.opponentCards); i++ {
	// 	fmt.Print("Opponent's turn \n --------------- \n")
	// 	for j := 0; j < len(has1); j++ {
	// 		if has1[j].value != game.opponentCards[i].value {
	// 			has1 = append(has1, game.opponentCards[i])
	// 		} else {
	// 			for k := 0; k < len(has2); k++ {
	// 				if has2[k].value != game.opponentCards[i].value {
	// 					has2 = append(has2, game.opponentCards[i])
	// 				} else {
	// 					for l := 0; l < len(has3); l++ {
	// 						if has3[l].value != game.opponentCards[i].value {
	// 							has3 = append(has3, game.opponentCards[i])
	// 						} else {
	// 							fmt.Printf("Book collected!")
	// 							has4 = append(has4, game.opponentCards[i])
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// ask randomly for value that has most of
	// if len(has3) > 0 {
	// 	rand.Shuffle(len(has3), func(i, j int) { has3[i], has3[j] = has3[j], has3[i] })
	// 	// ask := has3[0].value
	// } else if len(has2) > 0 {
	// 	rand.Shuffle(len(has2), func(i, j int) { has2[i], has2[j] = has2[j], has2[i] })
	// 	// ask := has2[0].value
	// } else {
	// 	rand.Shuffle(len(has1), func(i, j int) { has1[i], has1[j] = has1[j], has1[i] })
	// 	// ask := has1[0].value
	// }

	// fmt.Print(has3, "\n", has2, "\n", has1, "\n")

	// rand.Shuffle(len(game.opponentCards), func(i, j int) {
	// 	game.opponentCards[i], game.opponentCards[j] = game.opponentCards[j], game.opponentCards[i]
	// })
	// ask := game.opponentCards[0].value

	// fmt.Print("Opponent is asking for ", ask, "\n")

	// game.searchCards(ask)

	// hasBook(game.opponentCards)

	fmt.Printf("opponent's turn \n")
	if len(game.playerCards)+len(game.opponentCards) != 52 {
		game.play()
	} else {
		return false
	}
	return false
}

// check if player has book, remove book from hand, add to books
func hasBook(cards []Card) bool {
	return false
}

func enterString() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\r')
	if err != nil {
		fmt.Println("error. retry", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\r")
	input = strings.TrimSuffix(input, "\n")
	return input
}

func displayCards(cards []Card) {
	displayStr := ""

	for i, card := range cards {
		displayStr += card.getString()
		if i < len(cards)-1 {
			displayStr += " "
		}
	}
	fmt.Printf("%v\n", displayStr)
}

func main() {
	game := Game{}
	game.dealStartingCards()
	game.play()
}
