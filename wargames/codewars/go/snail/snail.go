// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1
package main

import (
	"fmt"
)

func Snail(m [][]int) []int {
	n := len(m)
	result := make([]int, n*n)
	fmt.Println(n)
	if n == 1 {
		if len(m[0]) == 0 {
			return []int{}
		} else {
			result[0] = m[0][0]
			return result
		}
	}
	it := 0
	count := 0
	for count < n*n {
		// top line
		for i := it; i < n-it; i++ {
			// fmt.Print(m[it][i], " ")
			result[count] = m[it][i]
			count++
		}
		// fmt.Println()
		// right column
		for i := it + 1; i < n-it; i++ {
			// fmt.Print(m[i][n-it-1], " ")
			result[count] = m[i][n-it-1]
			count++
		}
		// fmt.Println()
		// bottom line
		for i := n - it - 2; i >= it; i-- {
			// fmt.Print(m[n-it-1][i], " ")
			result[count] = m[n-it-1][i]
			count++
		}
		// fmt.Println()
		// left column
		for i := n - it - 2; i > it; i-- {
			// fmt.Print(m[i][it], " ")
			result[count] = m[i][it]
			count++
		}
		// fmt.Println()
		it++
	}
	return result
}

func makeMap(n int) [][]int {
	count := 0
	newMap := make([][]int, n)
	for i := 0; i < n; i++ {
		newMap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			newMap[i][j] = count
			count++
			fmt.Print(newMap[i][j], " ")
		}
		fmt.Println()
	}
	return newMap
}

func main() {
	// fmt.Println("Solving 'Snail'")
	fmt.Println(Snail([][]int{{}}))
	fmt.Println(Snail([][]int{{1}}))
	fmt.Println(Snail(makeMap(3)))
	fmt.Println(Snail(makeMap(2)))
	fmt.Println(Snail(makeMap(0)))
	fmt.Println(Snail(makeMap(14)))
}
