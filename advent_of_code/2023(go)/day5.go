package adventofcode

import (
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type matchRelationship struct {
	from  int64
	end   int64
	steps int64
}

type seedsLocation struct {
	// seems int is enough, no need for int64 in part 2
	from int64
	end  int64
}

var (
	seedToSoil              = map[int64]matchRelationship{}
	soilToFertilizer        = map[int64]matchRelationship{}
	fertilizerToWater       = map[int64]matchRelationship{}
	waterToLight            = map[int64]matchRelationship{}
	lighToTemperature       = map[int64]matchRelationship{}
	temperatureToHumidility = map[int64]matchRelationship{}
	humidilityToLocation    = map[int64]matchRelationship{}

	allMap = map[string]map[int64]matchRelationship{
		"seed-to-soil":            seedToSoil,
		"soil-to-fertilizer":      soilToFertilizer,
		"fertilizer-to-water":     fertilizerToWater,
		"water-to-light":          waterToLight,
		"light-to-temperature":    lighToTemperature,
		"temperature-to-humidity": temperatureToHumidility,
		"humidity-to-location":    humidilityToLocation,
	}
)

func findNextResource(match map[int64]matchRelationship, from int64) int64 {

	for start, resources := range match {
		// if where it starts not in the range, then no need to retry
		if start > from {
			continue
		}
		if start+resources.steps < from {
			continue
		}

		return resources.end + (from - resources.from)
	}

	return from
}

func day5Part1(input []string) int {
	// get pre-calculated relationship
	processInput(input)

	// get all seeds
	minimumLocation := math.MaxInt64
	for _, seedStr := range strings.Fields(strings.Split(input[0], ":")[1]) {
		seed, _ := strconv.ParseInt(seedStr, 10, 64)
		// get soil
		soil := findNextResource(seedToSoil, seed)
		fertilizer := findNextResource(soilToFertilizer, soil)
		water := findNextResource(fertilizerToWater, fertilizer)
		light := findNextResource(waterToLight, water)
		temperature := findNextResource(lighToTemperature, light)
		humidity := findNextResource(temperatureToHumidility, temperature)
		location := findNextResource(humidilityToLocation, humidity)
		minimumLocation = min(int(location), minimumLocation)
	}

	return minimumLocation
}

func day5Part2(input []string) int {
	// get pre-calculated relationship
	// processInput(input)

	// get all seeds
	minimumLocation := math.MaxInt64

	seedsList := strings.Fields(strings.Split(input[0], ":")[1])

	sls := []seedsLocation{}
	for i := 0; i < len(seedsList); i += 2 {
		start, _ := strconv.ParseInt(seedsList[i], 10, 64)
		gap, _ := strconv.ParseInt(seedsList[i+1], 10, 64)

		// first location range
		if len(sls) == 0 {
			sls = append(sls, seedsLocation{
				from: start,
				end:  start + gap,
			})
			continue
		}
		// later find whether there is overlap
		findOverlap(&sls, start, start+gap)
	}
	for _, seeds := range sls {
		for seedInt := seeds.from; seedInt <= seeds.end; seedInt++ {
			// get soil
			seed := int64(seedInt)
			soil := findNextResource(seedToSoil, seed)
			fertilizer := findNextResource(soilToFertilizer, soil)
			water := findNextResource(fertilizerToWater, fertilizer)
			light := findNextResource(waterToLight, water)
			temperature := findNextResource(lighToTemperature, light)
			humidity := findNextResource(temperatureToHumidility, temperature)
			location := findNextResource(humidilityToLocation, humidity)
			minimumLocation = min(int(location), minimumLocation)
		}
	}

	return minimumLocation
}

func findOverlap(sls *[]seedsLocation, from, end int64) {
	isFound := false
	index := 0
	newSls := seedsLocation{from: from, end: end}

	for index = range *sls {
		// scenario 1, no overlap at all
		if (*sls)[index].end < from || (*sls)[index].from > end {
			continue
		} else {
			// generate new sls
			newSls.from = min(from, (*sls)[index].from)
			newSls.end = max(end, (*sls)[index].end)
			isFound = true
			break
		}
	}

	// delete previous one if already found
	if isFound {
		slices.Delete(*sls, index, index)
	}
	*sls = append(*sls, newSls)
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
			dest, _ := strconv.ParseInt(numbers[0], 10, 64)
			source, _ := strconv.ParseInt(numbers[1], 10, 64)
			steps, _ := strconv.ParseInt(numbers[2], 10, 64)

			selectedMap := allMap[filterPrefix]
			matchRelationship := matchRelationship{
				from:  source,
				end:   dest,
				steps: steps,
			}
			selectedMap[source] = matchRelationship

		}
	}
}
