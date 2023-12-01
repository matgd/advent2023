package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type Round struct {
	red   int
	green int
	blue  int
}

func (r Round) roundIsPossible() bool {
	return r.red <= 12 && r.green <= 13 && r.blue <= 14
}

type Game struct {
	id     int
	rounds []Round
}

func (g Game) gameIsPossible() bool {
	for _, round := range g.rounds {
		if !round.roundIsPossible() {
			return false
		}
	}

	return true
}

func parseInput(filePath string) []Game {
	games := []Game{}
	lines := utils.ReadLines(filePath)

	for i, line := range lines {
		game := Game{id: i + 1}
		noGamesText := strings.Split(line, ": ")[1]
		roundsSplit := strings.Split(noGamesText, "; ")

		rounds := []Round{}
		for _, roundText := range roundsSplit {
			round := Round{}

			rRed := regexp.MustCompile(`(\d+) red`)
			rGreen := regexp.MustCompile(`(\d+) green`)
			rBlue := regexp.MustCompile(`(\d+) blue`)

			redMatch := rRed.FindStringSubmatch(roundText)
			greenMatch := rGreen.FindStringSubmatch(roundText)
			blueMatch := rBlue.FindStringSubmatch(roundText)

			if redMatch == nil {
				round.red = 0
			} else {
				red, _ := strconv.Atoi(redMatch[1])
				round.red = red
			}

			if greenMatch == nil {
				round.green = 0
			} else {
				green, _ := strconv.Atoi(greenMatch[1])
				round.green = green
			}

			if blueMatch == nil {
				round.blue = 0
			} else {
				blue, _ := strconv.Atoi(blueMatch[1])
				round.blue = blue
			}

			rounds = append(rounds, round)
		}
		game.rounds = rounds
		games = append(games, game)
	}

	return games
}

func Part1(filePath string) int {
	games := parseInput(filePath)
	sum := 0
	for _, game := range games {
		if game.gameIsPossible() {
			sum += game.id
		}
	}
	return sum
}

func Part2(filePath string) int {
	games := parseInput(filePath)
	sum := 0
	for _, game := range games {
		minRed, minGreen, minBlue := 0, 0, 0
		for _, round := range game.rounds {
			minRed = max(minRed, round.red)
			minGreen = max(minGreen, round.green)
			minBlue = max(minBlue, round.blue)
		}
		power := minRed * minGreen * minBlue
		sum += power
	}

	return sum
}

func main() {
	fmt.Println("[Part 1]", Part1(INPUT), 8)
	fmt.Println("[Part 2]", Part2(INPUT), 2286)
}
