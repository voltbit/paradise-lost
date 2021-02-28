package main

import (
	"fmt"
	"strings"
)

var DAY string = "19"

func parseMonsterData(testData []string) (map[string][]string, []string) {
	rules := make(map[string][]string)
	delim := 0
	for index, rule := range testData {
		if len(strings.TrimSpace(rule)) == 0 {
			delim = index
			break
		}
		rawRule := strings.Split(rule, ": ")
		rules[rawRule[0]] = strings.Split(rawRule[1], " | ")
	}
	return rules, testData[delim:]
}

func monsterScan(rawRules []string) int {
	rules, data := parseMonsterData(rawRules)
	fmt.Println(rules, data)
	return -1
}

func Day19() {
	testData, err := GetStringList(fmt.Sprintf("input/day_%s_example", DAY))
	// data, err := GetStringList(fmt.Sprintf("input/day_%s", DAY))
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("AoC 2020 Day %s", DAY))
	fmt.Println(monsterScan(testData))
	// fmt.Println(monsterScan(data))
}
