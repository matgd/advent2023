package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type Race struct {
	time           int
	recordDistance int
}

func (r Race) clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > r.time {
		return r.time
	}
	return v
}

func parseInput(path string, part int) []Race {
	var input []string
	if part == 1 {
		input = strings.Split(utils.LinesFromShell("cut -d: -f2 " + path + " | xargs")[0], " ")
	} else {
		input = strings.Split(utils.LinesFromShell("cut -d: -f2 " + path + " | tr '\n' '|' | xargs | sed -e 's/ //g' -e 's/.$//'")[0], "|")
	}

	races := []Race{}

	for i := 0; i < len(input)/2; i++ {
		races = append(races, Race{
			time:           utils.ToInt(input[i]),
			recordDistance: utils.ToInt(input[len(input)/2+i]),
		})
		fmt.Println(races)
	}

	return races
}

func Part1() int {
	races := parseInput(INPUT, 1)

	winningRanges := [][2]int{}

	for _, race := range races {

		leftBoundTime, rightBoundTime := 0, race.time
		minimumStartup := leftBoundTime
		maximumStartup := rightBoundTime

		// Binary search

		// Minimum -> go left
		tooHigh := false
		tooLow := false
		for true {
			currentStartupTime := race.clamp(int(math.Round((float64(rightBoundTime) + float64(leftBoundTime)) / 2)))
			distanceTraveled := (race.time - currentStartupTime) * currentStartupTime
			if distanceTraveled > race.recordDistance {
				tooHigh = true
				tooLow = false
				rightBoundTime = currentStartupTime
			}
			if distanceTraveled < race.recordDistance {
				tooHigh = false
				tooLow = true
				leftBoundTime = currentStartupTime
			}
			if distanceTraveled == race.recordDistance {
				// Need to break record, not be equal to it hence +1
				leftBoundTime = currentStartupTime + 1
				break
			}
			if leftBoundTime == rightBoundTime-1 {
				if tooLow {
					leftBoundTime++
				} else if tooHigh {
					leftBoundTime--
				}
			}
			if leftBoundTime == rightBoundTime {
				break
			}
		}
		minimumStartup = leftBoundTime

		leftBoundTime, rightBoundTime = minimumStartup+1, race.time
		tooHigh = false
		tooLow = false

		// Maximum -> go right
		for true {
			currentStartupTime := race.clamp(int(math.Round((float64(rightBoundTime) + float64(leftBoundTime)) / 2)))
			distanceTraveled := (race.time - currentStartupTime) * currentStartupTime
			if distanceTraveled > race.recordDistance {
				tooHigh = false
				tooLow = true
				leftBoundTime = currentStartupTime
			}
			if distanceTraveled < race.recordDistance {
				tooHigh = true
				tooLow = false
				rightBoundTime = currentStartupTime
			}
			if distanceTraveled == race.recordDistance {
				// Need to break record, not be equal to it hence -1
				rightBoundTime = currentStartupTime - 1
				break
			}

			if leftBoundTime == rightBoundTime {
				break
			}
			if leftBoundTime == rightBoundTime-1 {
				if tooLow {
					rightBoundTime++
				} else if tooHigh {
					rightBoundTime--
				}
			}
		}
		maximumStartup = rightBoundTime

		winningRanges = append(winningRanges, [2]int{minimumStartup, maximumStartup})
	}

	multiplied := 1
	for _, wr := range winningRanges {
		multiplied *= wr[1] - wr[0] + 1
	}
	fmt.Println(winningRanges)
	return multiplied
}

func Part2() int {
	races := parseInput(INPUT, 2)

	timeStart := time.Now()
	winningRange := [2]int{}

	for _, race := range races {
		leftBoundTime, rightBoundTime := 0, race.time
		minimumStartup := leftBoundTime
		maximumStartup := rightBoundTime

		// Binary search

		// Minimum -> go left
		tooHigh := false
		tooLow := false
		for true {
			currentStartupTime := race.clamp(int(math.Round((float64(rightBoundTime) + float64(leftBoundTime)) / 2)))
			distanceTraveled := (race.time - currentStartupTime) * currentStartupTime
			if distanceTraveled > race.recordDistance {
				tooHigh = true
				tooLow = false
				rightBoundTime = currentStartupTime
			}
			if distanceTraveled < race.recordDistance {
				tooHigh = false
				tooLow = true
				leftBoundTime = currentStartupTime
			}
			if distanceTraveled == race.recordDistance {
				// Need to break record, not be equal to it hence +1
				leftBoundTime = currentStartupTime + 1
				break
			}
			if leftBoundTime == rightBoundTime-1 {
				if tooLow {
					leftBoundTime++
				} else if tooHigh {
					leftBoundTime--
				}
			}
			if leftBoundTime == rightBoundTime {
				break
			}
		}
		minimumStartup = leftBoundTime

		leftBoundTime, rightBoundTime = minimumStartup+1, race.time
		tooHigh = false
		tooLow = false

		// Maximum -> go right
		for true {
			currentStartupTime := race.clamp(int(math.Round((float64(rightBoundTime) + float64(leftBoundTime)) / 2)))
			distanceTraveled := (race.time - currentStartupTime) * currentStartupTime
			if distanceTraveled > race.recordDistance {
				tooHigh = false
				tooLow = true
				leftBoundTime = currentStartupTime
			}
			if distanceTraveled < race.recordDistance {
				tooHigh = true
				tooLow = false
				rightBoundTime = currentStartupTime
			}
			if distanceTraveled == race.recordDistance {
				// Need to break record, not be equal to it hence -1
				rightBoundTime = currentStartupTime - 1
				break
			}

			if leftBoundTime == rightBoundTime {
				break
			}
			if leftBoundTime == rightBoundTime-1 {
				if tooLow {
					rightBoundTime++
				} else if tooHigh {
					rightBoundTime--
				}
			}
		}
		maximumStartup = rightBoundTime

		winningRange = [2]int{minimumStartup, maximumStartup}
	}

	fmt.Println("Time taken:", time.Since(timeStart))
	return winningRange[1] - winningRange[0] + 1
}

func main() {
	p1 := Part1()
	p2 := Part2()
	fmt.Println("[Part 1]", p1 == 288, p1)
	fmt.Println("[Part 2]", p2 == 71503, p2)
}
