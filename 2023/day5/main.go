package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

type Map struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func (m Map) getValue(sourceNum int) (int, bool) {
	if sourceNum < m.sourceStart || sourceNum >= m.sourceStart+m.rangeLength {
		return sourceNum, false
	}
	destLen := sourceNum - m.sourceStart
	return m.destinationStart + destLen, true
}

func getMappingData(lines []string) []Map {
	maps := []Map{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		values := strings.Fields(line)
		destStart, _ := strconv.Atoi(values[0])
		sourceStart, _ := strconv.Atoi(values[1])
		rangeLen, _ := strconv.Atoi(values[2])
		maps = append(maps, Map{destStart, sourceStart, rangeLen})
	}

	return maps
}

func getMapperValue(sourceNum int, mappers []Map) int {
	for _, mapper := range mappers {
		if destNum, ok := mapper.getValue(sourceNum); ok {
			return destNum
		}
	}
	return sourceNum
}

func part1(lines []string) interface{} {
	seedsStr := lines[0]
	seedStrs := strings.Fields(seedsStr[5:])
	seedData := make(map[int]map[string]int)

	for _, seed := range seedStrs {
		seedNum, _ := strconv.Atoi(seed)
		seedData[seedNum] = map[string]int{}
	}

	delete(seedData, 0)

	fmt.Println("Initial Seed Data generated")

	seedToSoilMapIdx := slices.Index(lines, "seed-to-soil map:")
	soilToFertilizerMapIdx := slices.Index(lines, "soil-to-fertilizer map:")
	fertilizerToWaterMapIdx := slices.Index(lines, "fertilizer-to-water map:")
	waterToLightMapIdx := slices.Index(lines, "water-to-light map:")
	lightToTemperatureMapIdx := slices.Index(lines, "light-to-temperature map:")
	temperatureToHumidityMapIdx := slices.Index(lines, "temperature-to-humidity map:")
	humidityToLocationMapIdx := slices.Index(lines, "humidity-to-location map:")

	fmt.Println("Indexes determined")

	soilMap := getMappingData(lines[seedToSoilMapIdx+1 : soilToFertilizerMapIdx])
	for seedNum := range seedData {
		seedData[seedNum]["soil"] = getMapperValue(seedNum, soilMap)
	}

	fmt.Println("Soil complete")

	fertilizerMap := getMappingData(lines[soilToFertilizerMapIdx+1 : fertilizerToWaterMapIdx])
	for seedNum, data := range seedData {
		seedData[seedNum]["fertilizer"] = getMapperValue(data["soil"], fertilizerMap)
	}

	fmt.Println("Fertilizer complete")

	waterMap := getMappingData(lines[fertilizerToWaterMapIdx+1 : waterToLightMapIdx])
	for seedNum, data := range seedData {
		seedData[seedNum]["water"] = getMapperValue(data["fertilizer"], waterMap)
	}

	fmt.Println("Water complete")

	lightMap := getMappingData(lines[waterToLightMapIdx+1 : lightToTemperatureMapIdx])
	for seedNum, data := range seedData {
		seedData[seedNum]["light"] = getMapperValue(data["water"], lightMap)
	}

	fmt.Println("Light complete")

	temperatureMap := getMappingData(lines[lightToTemperatureMapIdx+1 : temperatureToHumidityMapIdx])
	for seedNum, data := range seedData {
		seedData[seedNum]["temperature"] = getMapperValue(data["light"], temperatureMap)
	}

	fmt.Println("Temperature complete")

	humidityMap := getMappingData(lines[temperatureToHumidityMapIdx+1 : humidityToLocationMapIdx])
	for seedNum, data := range seedData {
		seedData[seedNum]["humidity"] = getMapperValue(data["temperature"], humidityMap)
	}

	fmt.Println("Humidity complete")

	locationMap := getMappingData(lines[humidityToLocationMapIdx+1:])
	for seedNum, data := range seedData {
		seedData[seedNum]["location"] = getMapperValue(data["humidity"], locationMap)
	}

	fmt.Println("Location complete")

	min := math.MaxInt

	for seedNum, data := range seedData {
		if slices.Contains(seedStrs, strconv.Itoa(seedNum)) {
			if data["location"] < min {
				min = data["location"]
			}
		}
	}

	return min
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2023/day5/test_case.txt")
	input := utils.ReadLines("./2023/day5/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
