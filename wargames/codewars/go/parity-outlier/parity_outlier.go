package main

import (
	"fmt"
	"math"
)

func FindOutlier_fail(integers []int) int {
	var flag int
	mask := integers[0]%2 | integers[1]%2<<1 | integers[2]%2<<2
	fmt.Println("Mask:", mask)
	if mask == 0 {
		flag = 0
	} else if mask == 7 {
		flag = 1
	}
	for _, x := range integers {
		x++
	}
	return flag
}

func FindOutlier(integers []int) int {
	odd, even := 0, 0
	last_odd, last_even := 0, 0
	for i, x := range integers {
		if x%2 == 0 {
			even++
			last_even = i
		} else {
			odd++
			last_odd = i
		}
		if i > 1 {
			if odd == 1 {
				return integers[last_odd]
			} else if even == 1 {
				return integers[last_even]
			}
		}
	}
	return integers[0]
}

func main() {
	fmt.Println(FindOutlier([]int{0, -2, 9, 100, 4, 11, 2602, 36}))
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(FindOutlier([]int{math.MaxInt32, 0, 1}))
	// fmt.Println(FindOutlier_fail([]int{-3, -7, 9, 100, 4, 11, 2602, 36}))
	// fmt.Println(FindOutlier_fail([]int{-2, -4, 0, 100, 4, 11, 2602, 36}))
	// fmt.Println(FindOutlier_fail([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	// fmt.Println(FindOutlier_fail([]int{160, 3, 1719, 19, 11, 13, -21}))
}
