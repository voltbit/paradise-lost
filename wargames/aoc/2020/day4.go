package main

import (
	"fmt"
	"regexp"
	"strings"
)

func getPassportData(input []string) []map[string]string {
	var passportData []map[string]string
	passport := make(map[string]string)
	for i, line := range input {
		if line == "" || i == len(input)-1 {
			passportData = append(passportData, passport)
			passport = make(map[string]string)
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			kv := strings.Split(field, ":")
			passport[kv[0]] = kv[1]
		}
	}
	// fmt.Printf("Input size: %d\n", len(passportData))
	// fmt.Println(passportData)
	return passportData
}

func checkValidData(p map[string]string) bool {
	if match, _ := regexp.MatchString("19[2-9][0-9]|200[0-2]", p["byr"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("201[0-9]|2020", p["iyr"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("202[0-9]|2030", p["eyr"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in)", p["hgt"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("#[0-9a-f]{6}", p["hcl"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", p["ecl"]); !match {
		return false
	}
	if match, _ := regexp.MatchString("[0-9]{9}", p["pid"]); !match {
		return false
	}
	return true
}

func processPassport(input []string) (int, int) {
	checkKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	passportData := getPassportData(input)
	validCount := 0
	validDataCount := 0
	for _, passport := range passportData {
		valid := true
		for _, checkKey := range checkKeys {
			if _, found := passport[checkKey]; !found {
				valid = false
				// fmt.Printf("Missing: %s in %v\n", checkKey, passport)
				break
			}
		}
		if valid {
			validCount++
			if checkValidData(passport) {
				validDataCount++
			}
		}
	}
	return validCount, validDataCount
}

func Day4() {
	fmt.Println("AoC day 4")
	testInput, err := GetStringList("input/day_4_example")
	if err != nil {
		panic(err)
	}
	fmt.Println(processPassport(testInput))

	input, err := GetStringList("input/day_4")
	if err != nil {
		panic(err)
	}
	fmt.Println(processPassport(input))
}
