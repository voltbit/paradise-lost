package main

import "fmt"

func solve(in []int) (int, int) {
	sol_1 := 0
	sol_2 := 0
	numMap := buildMaps(in)
	for _, num := range in {
		if numMap[2020-num] == 1 {
			sol_1 = num * (2020 - num)
			break
		}
	}
	for _, num := range in {
		for _, num2 := range in {
			diff := 2020 - num - num2
			if diff >= 0 && numMap[diff] == 1 {
				sol_2 = num * num2 * diff
				return sol_1, sol_2
			}
		}
	}
	return -1, -1
}

func buildMaps(in []int) []int {
	numMap := make([]int, 2021)
	for _, num := range in {
		numMap[num] = 1
	}
	return numMap
}

func main() {
	testInput := []int{1721, 979, 366, 299, 675, 1456}
	input, err := GetIntList("input/day_1")
	if err != nil {
		panic(err)
	}
	fmt.Println(solve(testInput))
	fmt.Println(solve(input))
}
