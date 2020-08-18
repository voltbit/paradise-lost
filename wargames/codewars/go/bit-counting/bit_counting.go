package main

import "fmt"

func CountBits(x uint) int {
	count := 0
	for x > 0 {
		if x&1 == 1 {
			count += 1
		}
		x = x >> 1
	}
	return count
}

func main() {
	fmt.Println(CountBits(1023))
}
