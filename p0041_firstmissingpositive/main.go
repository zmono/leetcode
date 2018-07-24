// 41. First Missing Positive
// https://leetcode.com/problems/first-missing-positive/description/

package main

import (
	"fmt"
)

type TestCase struct {
	input  []int
	output int
}

func sortBit(nums []int, bit uint) int {
	mask := 1 << bit
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i]&mask == 0 {
			i++
		}
		for j >= 0 && nums[j]&mask != 0 {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return i
}

func sort(nums []int, bit uint) {
	if len(nums) < 2 {
		return
	}
	i := sortBit(nums, bit)
	if bit > 0 {
		sort(nums[:i], bit-1)
		sort(nums[i:], bit-1)
	}
}

func firstMissingPositive(nums []int) int {
	sort(nums, 64)
	res := 1
	for _, n := range nums {
		if n == res {
			res++
		} else if n > res {
			return res
		}
	}
	return res
}

func main() {
	// nums := []int{1, 1}
	// sort(nums, 64)
	// fmt.Println(nums)
	cases := []TestCase{
		{[]int{0}, 1},
		{[]int{1}, 2},
		{[]int{1, 1}, 2},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{-1}, 1},
		{[]int{-1, -1, -1}, 1},
		{[]int{3, 2, 1}, 4},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := firstMissingPositive(c.input); output != c.output {
			fmt.Printf("firstMissingPositive(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
