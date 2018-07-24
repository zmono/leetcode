// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/description/

package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	input  []int
	output int
}

func longestConsecutive(nums []int) (max int) {
	h := make(map[int]int)
	for _, n := range nums {
		if _, ok := h[n]; !ok {
			h[n] = h[n-1] + 1 + h[n+1]
			h[n-h[n-1]], h[n+h[n+1]] = h[n], h[n]
			if max < h[n] {
				max = h[n]
			}
		}
	}
	return
}

func main() {
	cases := []TestCase{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 3}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 4, 2, 5, 7, 8, 3, 6}, 8},
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{4, 100, 200, 1, 102, 202, 2, 101, 201, 3}, 4},
		{[]int{0, 3, 7, 2, 5, 4, 6, 1}, 8},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{-4, -1, 4, -5, 1, -6, 9, 0, 2, 7, -3, 8, -2, 5, 3}, 12},
		{[]int{-4, -1, 4, -5, 1, -6, 9, -6, 0, 2, 2, 7, 0, 9, -3, 8, 9, -2, -6, 5, 0, 3, 4, -2}, 12},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := longestConsecutive(c.input); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("longestConsecutive(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
