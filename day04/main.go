package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type IntSet map[int]struct{}

func Part1() int {
	lines := utils.LinesFromShell("awk -F':' '{ print $2 }' " + INPUT + " | sed -e 's/  / /g' -e  's/^ //g'")
	score := 0

	for _, line := range lines {
		matchingNumbers := 0
		card := make(IntSet)
		numbers := make(IntSet)

		split := strings.Split(line, " | ")
		cardStr := strings.Split(split[0], " ")
		myNumbersStr := strings.Split(split[1], " ")

		for _, numStr := range cardStr {
			numInt, _ := strconv.Atoi(numStr)
			card[numInt] = struct{}{}
		}
		for _, numStr := range myNumbersStr {
			numInt, _ := strconv.Atoi(numStr)
			numbers[numInt] = struct{}{}
		}

		for number := range numbers {
			if _, ok := card[number]; ok {
				matchingNumbers++
			}
		}
		if matchingNumbers > 0 {
			score += int(math.Pow(2, float64(matchingNumbers)-1))
		}
	}

	return score
}

func Part2() int {
	lines := utils.LinesFromShell("awk -F':' '{ print $2 }' " + INPUT + " | sed -e 's/  / /g' -e  's/^ //g'")
	queue := make([]int, 0, 20000)
	score := 0
	lastCardIndex := len(lines) - 1

	for i := 0; i < len(lines); i++ {
		queue = append(queue, i)
	}

	for len(queue) > 0 {
		score++
		currentCardIndex := queue[0]

		if currentCardIndex == lastCardIndex+1 {
			queue = queue[1:]
			continue
		}

		line := lines[currentCardIndex]
		newQueueItemOffset := 1

		card := make(IntSet)
		numbers := make(IntSet)

		split := strings.Split(line, " | ")
		cardStr := strings.Split(split[0], " ")
		myNumbersStr := strings.Split(split[1], " ")

		for _, numStr := range cardStr {
			numInt, _ := strconv.Atoi(numStr)
			card[numInt] = struct{}{}
		}
		for _, numStr := range myNumbersStr {
			numInt, _ := strconv.Atoi(numStr)
			numbers[numInt] = struct{}{}
		}

		for number := range numbers {
			if _, ok := card[number]; ok {
				queue = append(queue, currentCardIndex+newQueueItemOffset)
				newQueueItemOffset++
			}
		}

		queue = queue[1:]
	}

	return score
}

func main() {
	fmt.Println("[Part 1]", Part1(), 13)
	fmt.Println("[Part 2]", Part2(), 30)
}
