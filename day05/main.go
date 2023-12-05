package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"matgd.github.com/advent2023/utils"
)

const INPUT = "input.txt"

type Mappings struct {
	seeds                 []int
	seedToSoil            []RawMapping
	soilToFertilizer      []RawMapping
	fertilizerToWater     []RawMapping
	waterToLight          []RawMapping
	lightToTemperature    []RawMapping
	temperatureToHumidity []RawMapping
	humidityToLocation    []RawMapping
}

func (m Mappings) String() string {
	return fmt.Sprintf("seeds: %v\nseedToSoil: %v\nsoilToFertilizer: %v\nfertilizerToWater: %v\nwaterToLight: %v\nlightToTemperature: %v\ntemperatureToHumidity: %v\nhumidityToLocation: %v\n",
		m.seeds, m.seedToSoil, m.soilToFertilizer, m.fertilizerToWater, m.waterToLight, m.lightToTemperature, m.temperatureToHumidity, m.humidityToLocation)
}

type RawMapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

const (
	START                  = ""
	SEED_2_SOIL            = "seed-to-soil"
	SOIL_2_FERTILIZER      = "soil-to-fertilizer"
	FERTILIZER_2_WATER     = "fertilizer-to-water"
	WATER_2_LIGHT          = "water-to-light"
	LIGHT_2_TEMPERATURE    = "light-to-temperature"
	TEMPERATURE_2_HUMIDITY = "temperature-to-humidity"
	HUMIDITY_2_LOCATION    = "humidity-to-location"
)

func parseInput(path string) Mappings {
	lines := utils.ReadLines(path)
	m := Mappings{}
	readMode := START

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds: ") {
			seeds := strings.Split(line, "seeds: ")[1]
			for _, seed := range strings.Split(seeds, " ") {
				seedInt, _ := strconv.Atoi(seed)
				m.seeds = append(m.seeds, seedInt)
			}
			continue
		}

		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "seed-to-soil") {
			readMode = SEED_2_SOIL
			continue
		}
		if strings.HasPrefix(line, "soil-to-fertilizer") {
			readMode = SOIL_2_FERTILIZER
			continue
		}
		if strings.HasPrefix(line, "fertilizer-to-water") {
			readMode = FERTILIZER_2_WATER
			continue
		}
		if strings.HasPrefix(line, "water-to-light") {
			readMode = WATER_2_LIGHT
			continue
		}
		if strings.HasPrefix(line, "light-to-temperature") {
			readMode = LIGHT_2_TEMPERATURE
			continue
		}
		if strings.HasPrefix(line, "temperature-to-humidity") {
			readMode = TEMPERATURE_2_HUMIDITY
			continue
		}
		if strings.HasPrefix(line, "humidity-to-location") {
			readMode = HUMIDITY_2_LOCATION
			continue
		}

		splitRawMap := strings.Split(line, " ")

		destinationRangeStart, _ := strconv.Atoi(splitRawMap[0])
		sourceRangeStart, _ := strconv.Atoi(splitRawMap[1])
		rangeLength, _ := strconv.Atoi(splitRawMap[2])

		if readMode == SEED_2_SOIL {
			m.seedToSoil = append(m.seedToSoil, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == SOIL_2_FERTILIZER {
			m.soilToFertilizer = append(m.soilToFertilizer, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == FERTILIZER_2_WATER {
			m.fertilizerToWater = append(m.fertilizerToWater, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == WATER_2_LIGHT {
			m.waterToLight = append(m.waterToLight, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == LIGHT_2_TEMPERATURE {
			m.lightToTemperature = append(m.lightToTemperature, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == TEMPERATURE_2_HUMIDITY {
			m.temperatureToHumidity = append(m.temperatureToHumidity, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		if readMode == HUMIDITY_2_LOCATION {
			m.humidityToLocation = append(m.humidityToLocation, RawMapping{destinationRangeStart, sourceRangeStart, rangeLength})
		}
	}

	return m
}

func mappingToFunc(mapping RawMapping) func(int) (int, bool) {
	return func(source int) (int, bool) {
		if source < mapping.sourceRangeStart {
			return source, false
		}
		if source >= mapping.sourceRangeStart+mapping.rangeLength {
			return source, false
		}
		sourceOffest := source - mapping.sourceRangeStart
		return mapping.destinationRangeStart + sourceOffest, true
	}
}

func Part1() int {
	mappings := parseInput(INPUT)
	minLocation := math.Inf(1)

	for _, seed := range mappings.seeds {
		soil := seed
		for _, seedToSoil := range mappings.seedToSoil {
			if returnedSoil, ok := mappingToFunc(seedToSoil)(seed); ok {
				soil = returnedSoil
				break
			}
		}
		ferilizer := soil
		for _, soilToFertilizer := range mappings.soilToFertilizer {
			if returnedFertilizer, ok := mappingToFunc(soilToFertilizer)(soil); ok {
				ferilizer = returnedFertilizer
				break
			}
		}
		water := ferilizer
		for _, fertilizerToWater := range mappings.fertilizerToWater {
			if returnedWater, ok := mappingToFunc(fertilizerToWater)(ferilizer); ok {
				water = returnedWater
				break
			}
		}
		light := water
		for _, waterToLight := range mappings.waterToLight {
			if returnedLight, ok := mappingToFunc(waterToLight)(water); ok {
				light = returnedLight
				break
			}
		}
		temperature := light
		for _, lightToTemperature := range mappings.lightToTemperature {
			if returnedTemperature, ok := mappingToFunc(lightToTemperature)(light); ok {
				temperature = returnedTemperature
				break
			}
		}
		humidity := temperature
		for _, temperatureToHumidity := range mappings.temperatureToHumidity {
			if returnedHumidity, ok := mappingToFunc(temperatureToHumidity)(temperature); ok {
				humidity = returnedHumidity
				break
			}
		}
		location := humidity
		for _, humidityToLocation := range mappings.humidityToLocation {
			if returnedLocation, ok := mappingToFunc(humidityToLocation)(humidity); ok {
				location = returnedLocation
				break
			}
		}
		minLocation = math.Min(minLocation, float64(location))
	}

	return int(minLocation)
}

// func Part2() int {
// 	return 0
// }

func main() {
	fmt.Println("[Part 1]", Part1() == 111627841)
	// fmt.Println("[Part 2]", Part2(), -1)
}
