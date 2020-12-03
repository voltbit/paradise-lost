package main

import "fmt"

func checkSlope(mapLines []string, right int, down int) int {
	width := 1
	treeCount := 0
	for i, line := range mapLines {
		if i%down != 0 {
			continue
		}
		index := width%len(line) - 1
		if index == -1 {
			index = len(line) - 1
		}
		if string(line[index]) == "#" {
			treeCount++
		}
		width += right
	}
	return treeCount
}

func solve(mapLines []string) int {
	slopes := []struct {
		right int
		down  int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	slopeProduct := 1
	for _, slope := range slopes {
		slopeTrees := checkSlope(mapLines, slope.right, slope.down)
		fmt.Printf("slope (%d %d): %d\n", slope.right, slope.down, slopeTrees)
		slopeProduct *= slopeTrees
	}
	return slopeProduct
}

func Day3() {
	testData, err := GetStringList("input/day_3_example")
	data, err := GetStringList("input/day_3")
	if err != nil {
		panic(err)
	}
	fmt.Println("AoC 2020 Day 3")
	fmt.Println(solve(testData))
	fmt.Println(solve(data))
}
