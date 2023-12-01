package main

import (
	"fmt"
	"strconv"
	"strings"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type Number struct {
	numberCharacters []string
	yStart, yEnd     int
	xStart, xEnd     int
	gearsCoords      [][2]int
}

func (n Number) surroundedBySpecial(input []string) bool {
	for y := n.yStart - 1; y <= n.yEnd+1; y++ {
		for x := n.xStart - 1; x <= n.xEnd+1; x++ {

			if y > 0 && y < len(input) && x > 0 && x < len(input[y]) {
				if isSpecialChar(rune(input[y][x])) {
					return true
				}
			}
		}
	}

	return false
}

func (n Number) surroundedByGear(input []string) *[2]int {
	for y := n.yStart - 1; y <= n.yEnd+1; y++ {
		for x := n.xStart - 1; x <= n.xEnd+1; x++ {

			if y > 0 && y < len(input) && x > 0 && x < len(input[y]) {
				if input[y][x] == '*' {
					return &[2]int{y, x}
				}
			}
		}
	}

	return nil
}

func (n Number) number() int {
	v, _ := strconv.Atoi(strings.Join(n.numberCharacters, ""))
	return v
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSpecialChar(r rune) bool {
	return r != '.' && !isNumber(r)
}

func Part1() int {
	lines := utils.ReadLines(INPUT)
	sum := 0

	for il, line := range lines {
		numbers := []Number{}
		var n *Number
		numberStarted := false

		for ic, char := range line {
			if isNumber(char) {
				if !numberStarted {
					n = &Number{}
					numberStarted = true
					n.yStart = il
					n.xStart = ic
				}
				n.numberCharacters = append(n.numberCharacters, string(char))
				if ic == len(line)-1 {
					n.yEnd = il
					n.xEnd = ic
					numbers = append(numbers, *n)
				}

			} else {
				if numberStarted {
					numberStarted = false
					n.yEnd = il
					n.xEnd = ic - 1
					numbers = append(numbers, *n)
				}
			}
		}

		for _, n := range numbers {
			if n.surroundedBySpecial(lines) {
				sum += n.number()
			}
		}
	}

	return sum
}

func Part2() int {
	lines := utils.ReadLines(INPUT)
	gearsCoords := make(map[[2]int][]Number)
	sum := 0

	numbers := []Number{}
	for il, line := range lines {
		var n *Number
		numberStarted := false

		for ic, char := range line {
			if isNumber(char) {
				if !numberStarted {
					n = &Number{}
					numberStarted = true
					n.yStart = il
					n.xStart = ic
				}
				n.numberCharacters = append(n.numberCharacters, string(char))
				if ic == len(line)-1 {
					n.yEnd = il
					n.xEnd = ic
					numbers = append(numbers, *n)
				}

			} else {
				if numberStarted {
					numberStarted = false
					n.yEnd = il
					n.xEnd = ic - 1
					numbers = append(numbers, *n)
				}
			}
		}

	}

	for _, n := range numbers {
		if gearCoords := n.surroundedByGear(lines); gearCoords != nil {
			gearsCoords[*gearCoords] = append(gearsCoords[*gearCoords], n)
		}
	}

	for _, v := range gearsCoords {
		if len(v) == 2 {
			sum += v[0].number() * v[1].number()
		}
	}

	return sum
}

func main() {
	fmt.Println("[Part 1]", Part1(), 4361)
	fmt.Println("[Part 2]", Part2(), 467835)
}
