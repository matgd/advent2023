package main

import (
	"fmt"
	"sort"
	"strings"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

var cardSymbols = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type Card struct {
	label    rune
	strength int
}

func (c Card) String() string {
	return fmt.Sprintf("%c", c.label)
}

type Hand struct {
	cards                  [5]Card
	bid                    int
	cardsCount             map[Card]int
	strength               int
	individualCardStrength int
}

func (h *Hand) assignIndividualCardStrengthValue() {
	s := ""
	for _, c := range h.cards {
		s += fmt.Sprintf("%02d", c.strength)
	}
	h.individualCardStrength = utils.ToInt(s)
}

func (h Hand) String() string {
	return fmt.Sprintf("%v [%v] {%v} %v (%v)", h.cards, h.strength, h.individualCardStrength, h.cardsCount, h.bid)
}

func (h Hand) fiveOfAKind() bool {
	for _, v := range h.cardsCount {
		if v == 5 {
			return true
		}
	}
	return false
}

func (h Hand) fourOfAKind() bool {
	for _, v := range h.cardsCount {
		if v == 4 {
			return true
		}
	}
	return false
}

func (h Hand) fullHouse() bool {
	three := false
	pair := false
	for _, v := range h.cardsCount {
		if v == 3 {
			three = true
		}
		if v == 2 {
			pair = true
		}
	}
	return three && pair
}

func (h Hand) threeOfAKind() bool {
	three := false
	pair := false
	for _, v := range h.cardsCount {
		if v == 3 {
			three = true
		}
		if v == 2 {
			pair = true
		}
	}
	return three && !pair
}

func (h Hand) twoPair() bool {
	pair := 0
	for _, v := range h.cardsCount {
		if v == 2 {
			pair++
		}
	}
	return pair == 2
}

func (h Hand) onePair() bool {
	pair := 0
	for _, v := range h.cardsCount {
		if v == 2 {
			pair++
			if pair > 1 {
				return false
			}
		}
	}
	return pair == 1
}

func (h Hand) highCard() *Card {
	var maxCard *Card = cardMap['2']
	for _, card := range h.cards {
		if card.strength > maxCard.strength {
			maxCard = &card
		}
	}
	return maxCard
}

func (h *Hand) calculateStrength() {
	if h.fiveOfAKind() {
		h.strength = 1 << 5
		return
	}
	if h.fourOfAKind() {
		h.strength = 1 << 4
		return
	}
	if h.fullHouse() {
		h.strength = 1 << 3
		return
	}
	if h.threeOfAKind() {
		h.strength = 1 << 2
		return
	}
	if h.twoPair() {
		h.strength = 1 << 1
		return
	}
	if h.onePair() {
		h.strength = 1
		return
	}
	h.strength = 0
}

func getAllCards() map[rune]*Card {
	cards := make(map[rune]*Card)
	for i, s := range cardSymbols {
		cards[s] = &Card{s, len(cardSymbols) - i}
	}
	return cards
}

var cardMap = getAllCards()

func getHands() []Hand {
	var hands []Hand
	input := utils.ReadLines(INPUT)
	for _, line := range input {
		hand := Hand{}
		split := strings.Split(line, " ")

		cards := [5]Card{}
		cardsCount := make(map[Card]int)
		for i, c := range split[0] {
			cards[i] = *cardMap[c]
			cardsCount[cards[i]]++
		}

		hand.cards = cards
		hand.bid = utils.ToInt(split[1])
		hand.cardsCount = cardsCount
		hand.calculateStrength()
		hand.assignIndividualCardStrengthValue()
		hands = append(hands, hand)
	}

	return hands
}

func Part1() int {
	hands := getHands()

	bestHandsDesc := map[string]*[]Hand{
		"fiveOfAKind":  {},
		"fourOfAKind":  {},
		"fullHouse":    {},
		"threeOfAKind": {},
		"twoPair":      {},
		"onePair":      {},
		"highCard":     {},
	}

	for _, hand := range hands {
		if hand.strength == 1<<5 {
			*bestHandsDesc["fiveOfAKind"] = append(*bestHandsDesc["fiveOfAKind"], hand)
		}
		if hand.strength == 1<<4 {
			*bestHandsDesc["fourOfAKind"] = append(*bestHandsDesc["fourOfAKind"], hand)
		}
		if hand.strength == 1<<3 {
			*bestHandsDesc["fullHouse"] = append(*bestHandsDesc["fullHouse"], hand)
		}
		if hand.strength == 1<<2 {
			*bestHandsDesc["threeOfAKind"] = append(*bestHandsDesc["threeOfAKind"], hand)
		}
		if hand.strength == 1<<1 {
			*bestHandsDesc["twoPair"] = append(*bestHandsDesc["twoPair"], hand)
		}
		if hand.strength == 1 {
			*bestHandsDesc["onePair"] = append(*bestHandsDesc["onePair"], hand)
		}
		if hand.strength == 0 {
			*bestHandsDesc["highCard"] = append(*bestHandsDesc["highCard"], hand)
		}
	}

	iterOrder := []string{"fiveOfAKind", "fourOfAKind", "fullHouse", "threeOfAKind", "twoPair", "onePair", "highCard"}
	bestToWorst := []Hand{}
	for _, s := range iterOrder {
		if len(*bestHandsDesc[s]) == 0 {
			continue
		}
		if len(*bestHandsDesc[s]) == 1 {
			bestToWorst = append(bestToWorst, (*bestHandsDesc[s])[0])
			continue
		}
		sort.Slice(*bestHandsDesc[s], func(i, j int) bool {
			return (*bestHandsDesc[s])[i].individualCardStrength > (*bestHandsDesc[s])[j].individualCardStrength
		})
		for _, h := range *bestHandsDesc[s] {
			bestToWorst = append(bestToWorst, h)
		}
	}

	winnings := 0
	for i, btw := range bestToWorst {
		multiplier := len(bestToWorst) - i
		fmt.Println(btw, multiplier)
		winnings += btw.bid * multiplier
	}

	return winnings
}

func Part2() int {
	return 0
}

func main() {
	p1 := Part1()
	p2 := Part2()
	fmt.Println("[Part 1]", p1 == 6440, p1)
	fmt.Println("[Part 2]", p2 == -1, p2)
}
