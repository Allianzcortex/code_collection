package adventofcode

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	matchRegexp, _          = regexp.Compile(`(.*?)-to-(.*?) map:`)
	seedToSoil              = map[int]int{}
	soilToFertilizer        = map[int]int{}
	fertilizerToWater       = map[int]int{}
	waterToLight            = map[int]int{}
	lighToTemperature       = map[int]int{}
	temperatureToHumidility = map[int]int{}
	humidilityToLocation    = map[int]int{}

	allMap = map[string]map[int]int{
		"seed-to-soil":            seedToSoil,
		"soil-to-fertilizer":      soilToFertilizer,
		"fertilizer-to-water":     fertilizerToWater,
		"water-to-light":          waterToLight,
		"light-to-temperature":    lighToTemperature,
		"temperature-to-humidity": temperatureToHumidility,
		"humidity-to-location":    humidilityToLocation,
	}
)

func day5Part1NaiveSolution(input []string) int {
	// get pre-calculated relationship
	processInput(input)

	// get all seeds
	minimumLocation := math.MaxInt32
	for _, seedStr := range strings.Fields(strings.Split(input[0], ":")[1]) {
		seed, _ := strconv.Atoi(seedStr)
		// get soil
		soil := seed
		if expectedSoil, isFound := seedToSoil[seed]; isFound {
			soil = expectedSoil
		}
		fertilizer := soil
		if expectedFertilizer, isFound := soilToFertilizer[soil]; isFound {
			fertilizer = expectedFertilizer
		}
		water := fertilizer
		if expectedWater, isFound := fertilizerToWater[fertilizer]; isFound {
			water = expectedWater
		}
		light := water
		if expectedLight, isFound := waterToLight[water]; isFound {
			light = expectedLight
		}
		temperature := light
		if expectedTemperature, isFound := lighToTemperature[light]; isFound {
			temperature = expectedTemperature
		}
		humidity := temperature
		if expectedHumidity, isFound := temperatureToHumidility[temperature]; isFound {
			humidity = expectedHumidity
		}
		location := humidity
		if expectedLocation, isFound := humidilityToLocation[humidity]; isFound {
			location = expectedLocation
		}

		minimumLocation = min(location, minimumLocation)
	}

	return minimumLocation
}

func processInput(input []string) {

	var (
		filterPrefix string
	)
	for _, line := range input {
		if line == "" || strings.HasPrefix(line, "seeds") {
			continue
		}
		//
		if strings.Index(line, "map") != -1 {
			filterPrefix = strings.Split(line, " ")[0]
		} else {
			// find map
			numbers := strings.Split(line, " ")
			dest, _ := strconv.Atoi(numbers[0])
			source, _ := strconv.Atoi(numbers[1])
			step, _ := strconv.Atoi(numbers[2])

			selectedMap := allMap[filterPrefix]
			for i := 0; i < step; i++ {
				selectedMap[source+i] = dest + i
			}
		}
	}
}
