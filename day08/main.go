package main

import (
	"fmt"
	"regexp"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type Cycle struct {
	indexQ      []int
	currentItem int
}

func (c *Cycle) next() int {
	res := c.indexQ[c.currentItem]
	c.currentItem++
	if c.currentItem >= len(c.indexQ) {
		c.currentItem = 0
	}
	return res
}

func parseInput() (Cycle, map[string][2]string) {
	lines := utils.ReadLines(INPUT)
	cycle := Cycle{}
	elMap := make(map[string][2]string)

	for i, line := range lines {
		if i == 0 {
			for _, char := range line {
				if char == 'L' {
					cycle.indexQ = append(cycle.indexQ, 0)
				}
				if char == 'R' {
					cycle.indexQ = append(cycle.indexQ, 1)
				}
			}
		}
		if i >= 2 {
			re := regexp.MustCompile(`[a-zA-Z0-9]{3}`).FindAllString(line, -1)
			elMap[re[0]] = [2]string{re[1], re[2]}
		}
	}
	return cycle, elMap
}

func solve() int {
	cycle, elementMap := parseInput()
	steps := 0
	nextKey := "AAA"
	for true {
		if nextKey == "ZZZ" {
			break
		}
		steps++
		nextKey = elementMap[nextKey][cycle.next()]
	}

	return steps
}

func solve2() []int {
	cycle, elementMap := parseInput()
	nextKeys := []string{}
	steps := []int{}
	for k := range elementMap {
		if k[2] == 'A' {
			nextKeys = append(nextKeys, k)
			steps = append(steps, 0)
		}
	}

	minStepsToFind := []int{}
	for range nextKeys {
		minStepsToFind = append(minStepsToFind, 0)
	}
	fmt.Println(nextKeys)
	fmt.Println(steps)

	for true {
		allTrue := true
		for i, nk := range nextKeys {
			if nk[2] != 'Z' {
				allTrue = false
				break
			} else {
				allTrue = true
				minStepsToFind[i] = steps[i]
				if i == 2 {
					fmt.Println(nextKeys)
					fmt.Println(minStepsToFind)

				}
				steps[i] = 0
			}
		}
		if allTrue {
			break
		}
		nextIndex := cycle.next()
		for i, nk := range nextKeys {
			nextKeys[i] = elementMap[nk][nextIndex]
			steps[i]++
		}

		toReturn := true
		for v := range minStepsToFind {
			if v != 0 {
				toReturn = false
				break
			}
		}
		if toReturn {
			return minStepsToFind
		}
	}

	fmt.Println(nextKeys)
	fmt.Println(steps)
	return steps
}

func main() {
	fmt.Println("1-->", solve())
	fmt.Println("X 2-->", solve2()) // didn't work
}
