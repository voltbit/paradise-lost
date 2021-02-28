package main

import (
	"fmt"
	"math"
)

func trailing_zeros_n_fact(n int) int {
	zeros := 0
	log5 := math.Log(5)
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			zeros++
		} else if i%5 == 0 {
			v := math.Log(float64(i)) / log5
			if v == float64(int(v)) {
				zeros += int(v)
			} else {
				zeros++
			}
		}
	}
	return zeros
}

func main() {
	fmt.Printf("Check [%d]: %d\n", 6, trailing_zeros_n_fact(6))
	fmt.Printf("Check [%d]: %d\n", 12, trailing_zeros_n_fact(12))
	fmt.Printf("Check [%d]: %d\n", 30, trailing_zeros_n_fact(30))
	fmt.Printf("Check [%d]: %d\n", 1000, trailing_zeros_n_fact(1000))
}
