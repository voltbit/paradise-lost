package main

import (
	"fmt"
	"strconv"
)

func GrayCode(n int) []int {
	start := (1 << n) - 1
	res := []int{start}
	right := true
	for {
		if start == 1 {
			res = append([]int{0}, res...)
			break
		}
		if right {
			res = append(res, start-1)
			right = false
		} else {
			res = append([]int{start - 1}, res...)
			right = true
		}
		start--
	}
	fmt.Printf("Result:\n%+v\n", res)
	for _, i := range res {
		fmt.Printf("%s ", strconv.FormatInt(int64(i), 2))
	}
	return res
}
