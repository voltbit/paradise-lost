package leetcode

import "fmt"

const url = "https://leetcode.com/explore/challenge/card/july-leetcoding-challenge/547/week-4-july-22nd-july-28th/3401/"

/*MinSortedArray2 Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
Find the minimum element. The array may contain duplicates.
*/
func MinSortedArray2(nums []int) int {
	fmt.Println(url)
	if len(nums) < 0 { // TODO: learn to throw errors
		return 0
	} else if len(nums) == 0 {
		return nums[0]
	}
	h := int(len(nums) / 2)
	if nums[h] <= nums[0] {
		return MinSortedArray2(nums[0:h])
	}
	return nums[0]
}
