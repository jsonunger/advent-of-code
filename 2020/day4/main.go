package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Passport map[string]string

func buildPassports(lines []string) []Passport {
	passports := []Passport{}
	currentPassport := Passport{}
	for _, line := range lines {
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = Passport{}
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			keyVal := strings.Split(field, ":")
			currentPassport[keyVal[0]] = keyVal[1]
		}
	}
	return append(passports, currentPassport)
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func part1(lines []string) interface{} {
	passports := buildPassports(lines)

	numValid := 0
	for _, passport := range passports {
		valid := true
		for _, field := range requiredFields {
			if _, found := passport[field]; !found {
				valid = false
				break
			}
		}
		if valid {
			numValid++
		}
	}

	return numValid
}

func validate(key, value string) bool {
	switch key {
	case "byr":
		year := utils.ConvertStringToInt(value)
		return year >= 1920 && year <= 2002
	case "iyr":
		year := utils.ConvertStringToInt(value)
		return year >= 2010 && year <= 2020
	case "eyr":
		year := utils.ConvertStringToInt(value)
		return year >= 2020 && year <= 2030
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			height := utils.ConvertStringToInt(strings.TrimSuffix(value, "cm"))
			return height >= 150 && height <= 193
		} else if strings.HasSuffix(value, "in") {
			height := utils.ConvertStringToInt(strings.TrimSuffix(value, "in"))
			return height >= 59 && height <= 76
		}
	case "hcl":
		match, err := regexp.MatchString(`^#[0-9a-f]{6}$`, value)
		utils.PanicWithMsg(err, "matching pid regexp")
		return match
	case "ecl":
		match, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, value)
		utils.PanicWithMsg(err, "matching pid regexp")
		return match
	case "pid":
		match, err := regexp.MatchString(`^\d{9}$`, value)
		utils.PanicWithMsg(err, "matching pid regexp")
		return match
	}
	return false
}

func part2(lines []string) interface{} {
	passports := buildPassports(lines)

	numValid := 0
	for _, passport := range passports {
		valid := true
		for _, field := range requiredFields {
			if val, found := passport[field]; !found || !validate(field, val) {
				valid = false
				break
			}
		}
		if valid {
			numValid++
		}
	}

	return numValid
}

func main() {
	testCaseLines := utils.ReadLines("./2020/day4/test_case.txt")
	inputLines := utils.ReadLines("./2020/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
