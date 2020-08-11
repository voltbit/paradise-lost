// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	current, best := 0, 0
	for _, x := range numbers {
		if x > 0 {
			current += x
		} else {
			if current > best {
				best = current
			}
			current = 0
		}
	}
	return best
}

func main() {
	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
