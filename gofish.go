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
	defer fmt.Printf("\n------------------\n\n")
	fmt.Printf("\n------------------\n\n")

	game.dealStartingCards()

	// game logic
	if game.playerTurn() {

	}

	if game.opponentTurn() {

	}
}

func (game *Game) playerTurn() bool {
	for true {
		fmt.Printf("What would you like to fish for? ")
		gofish := enterString()
		bait, _ := strconv.ParseInt(gofish, 10, 64)
		temp := []Card{}

		// check opponent cards, if match, give to player, remove from opponent, return true
		for i := 0; i < len(game.opponentCards); i++ {
			fish := int64(game.opponentCards[i].value)
			fmt.Print(bait, fish, ". \n")
			if bait == fish {
				// add to card to player hand
				fmt.Printf("You caught a fish! \n")
				game.playerCards = append(game.playerCards, game.opponentCards[i])
				displayCards(game.playerCards)
				fmt.Printf("\n")

				// remove from opponent hand
				for j := 0; j < len(game.opponentCards); j++ {
					if bait != int64(game.opponentCards[j].value) {
						temp = append(temp, game.opponentCards[j])
					} else {

					}
				}
				game.opponentCards = temp
				displayCards(game.opponentCards)
				return true
			}
		}

		fmt.Printf("Go fish! \n")
		card := game.deck.deal(1)[0]
		game.playerCards = append(game.playerCards, card)
		displayCards(game.playerCards)
		displayCards(game.opponentCards)
		return false
	}
	return true
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

func (game *Game) opponentTurn() bool {
	// if has multi cards, ask for those
	// otherwise randomly choose 1 in hand to ask for
	// if success, cont. if fail, choose card from deck and end turn
	return true
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
	game.play()
}
