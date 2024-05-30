package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// To do:
// fix logic for opponent card selection
// what happens if cards in hand == 0 (player & opponent)

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

	fmt.Printf("\nPlayer has been dealt: \n")
	displayCards(game.playerCards)
	fmt.Printf("Opponent has been dealt: \n")
	displayCards(game.opponentCards)
}

func (game *Game) play() {
	defer fmt.Printf("\n end-of-play ------------------\n\n")
	fmt.Printf("\n start-of-play ------------------\n\n")

	if game.playerTurn() {
	}

	if game.opponentTurn() {
	}

	fmt.Print(game.isEndGame(game.playerBooks, game.opponentBooks))
}

func (game *Game) isEndGame(playerHand []Card, opponentHand []Card) string {
	var winner string
	if len(playerHand)+len(opponentHand) == 52 {
		if len(playerHand) > len(opponentHand) {
			winner = "You win!"
			// fmt.Print("You win!")
		} else if len(opponentHand) > len(playerHand) {
			winner = "Opponent wins!"
			// fmt.Print("Opponent wins!")
		} else if len(playerHand) == len(opponentHand) {
			winner = "You tied!"
			// fmt.Print("Tie!")
		}
	}
	return winner
}

func (game *Game) checkForBook(cards []Card, isPlayer bool) {
	counts := make(map[int]int)
	var newHand []Card
	var book []Card

	for _, card := range cards {
		counts[card.value]++
	}

	for _, card := range cards {
		if counts[card.value] == 4 {
			book = append(book, card)
		} else {
			newHand = append(newHand, card)
		}
	}

	if isPlayer {
		game.playerCards = newHand
		for i := 0; i < len(book); i++ {
			game.playerBooks = append(game.playerBooks, book[i])
		}
	} else {
		game.opponentCards = newHand
		for i := 0; i < len(book); i++ {
			game.opponentBooks = append(game.opponentBooks, book[i])
		}
	}

}

func (game *Game) chooseCard() []Card {

	rand.Shuffle(len(game.opponentCards), func(i, j int) {
		game.opponentCards[i], game.opponentCards[j] = game.opponentCards[j], game.opponentCards[i]
	})

	game.opponentCards = game.opponentCards[0:]
	return game.opponentCards
}

func (game *Game) playerTurn() bool {
	for true {
		game.checkForBook(game.playerCards, true)
		fmt.Print("Your hand: ")
		displayCards(game.playerCards)
		fmt.Print("Your books: ")
		displayCards(game.playerBooks)
		fmt.Print("\n")

		fmt.Printf("What would you like to fish for? ")
		gofish := enterString()
		bait, _ := strconv.ParseInt(gofish, 10, 64)
		temp := []Card{}
		removeCard := 0

		time.Sleep(1 * time.Second)
		fmt.Print("\n\n")

		// check opponent cards, if match, give to player, remove from opponent, return true
		for i := 0; i < len(game.opponentCards); i++ {
			fish := int64(game.opponentCards[i].value)

			if bait == fish {
				// add to card to player hand
				game.playerCards = append(game.playerCards, game.opponentCards[i])
				lastCard := game.playerCards[len(game.playerCards)-1].getString()
				fmt.Printf("You caught a fish! \n")
				fmt.Printf("You picked a %v. \n", lastCard)

				fmt.Printf("\n")

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

				time.Sleep(1 * time.Second)

				// continue playing if there are still cards in deck
				if len(game.playerCards)+len(game.opponentCards) != 52 {
					game.playerTurn()
				} else {
					return false
				}

			}
		}

		card := game.deck.deal(1)[0]
		game.playerCards = append(game.playerCards, card)
		lastCard := game.playerCards[len(game.playerCards)-1].getString()
		fmt.Printf("Go fish! ")
		fmt.Printf("You picked a %v. \n", lastCard)

		fmt.Printf("\n")

		time.Sleep(1 * time.Second)

		fmt.Printf("\n")
		fmt.Printf("Opponents turn!\n")
		fmt.Printf("-----------------------\n\n")

		game.opponentTurn()
	}
	return false
}

func (game *Game) opponentTurn() bool {
	time.Sleep(1 * time.Second)
	game.checkForBook(game.opponentCards, false)

	// temporary random selection
	request := game.chooseCard()[0].value
	fmt.Print("opponent asks for ", request, "\n\n")
	time.Sleep(1 * time.Second)
	hasCard := true

	removeCard := 0
	temp := []Card{}
	for i := 0; i < len(game.playerCards); i++ {
		fish := game.playerCards[i].value

		if request == fish {
			game.opponentCards = append(game.opponentCards, game.playerCards[i])
			fmt.Printf("You gave %v to opponent.\n", game.playerCards[i].getString())
			hasCard = true

			// remove card from player hand
			for j := 0; j < len(game.playerCards); j++ {
				card := game.playerCards[j].value
				if int(card) == int(fish) && removeCard == 0 {
					removeCard += 1
				} else {
					temp = append(temp, game.playerCards[j])
				}
			}
			game.playerCards = temp

			break
		} else {
			hasCard = false
		}
	}

	if hasCard == false {
		fmt.Printf("You do not have that card. Go fish!\n")
		fmt.Print("Opponent draws a card. \n\n")
		card := game.deck.deal(1)[0]
		game.opponentCards = append(game.opponentCards, card)

		time.Sleep(1 * time.Second)
		fmt.Printf("--------------------------\n\n")
	}

	if len(game.playerCards)+len(game.opponentCards) != 52 {
		if hasCard == true {
			game.opponentTurn()
		} else {
			game.playerTurn()
		}
	} else {
		return false
	}
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
