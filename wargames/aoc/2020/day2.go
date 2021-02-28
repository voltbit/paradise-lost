package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PasswordData struct {
	first    int
	second   int
	letter   string
	password string
}

func GetPasswordData(name string) ([]PasswordData, error) {
	var result []PasswordData
	raw, err := GetStringList(name)
	if err != nil {
		return nil, err
	}
	for _, line := range raw {
		fields := strings.Split(line, ": ")
		nums := strings.Split(fields[0], " ")
		first, _ := strconv.Atoi(strings.Split(nums[0], "-")[0])
		second, _ := strconv.Atoi(strings.Split(nums[0], "-")[1])
		passData := PasswordData{first: first, second: second, letter: nums[1], password: fields[1]}
		result = append(result, passData)
	}
	return result, nil
}

func fixPasswords(data []PasswordData) (int, int) {
	valid_1 := 0
	valid_2 := 0
	for _, data := range data {
		c := strings.Count(data.password, data.letter)
		if c >= data.first && c <= data.second {
			valid_1++
		}
		if len(data.password) > data.second-1 {
			if (string(data.password[data.first-1]) == data.letter && string(data.password[data.second-1]) != data.letter) ||
				(string(data.password[data.first-1]) != data.letter && string(data.password[data.second-1]) == data.letter) {
				valid_2++
			}
		}
	}
	return valid_1, valid_2
}

func Day2() {
	testInput := []PasswordData{
		{1, 3, "a", "abcde"},
		{1, 3, "b", "cdefg"},
		{2, 9, "c", "ccccccccc"},
	}
	input, err := GetPasswordData("input/day_2")
	if err != nil {
		panic(err)
	}
	fmt.Println(fixPasswords(testInput))
	fmt.Println(fixPasswords(input))
}
