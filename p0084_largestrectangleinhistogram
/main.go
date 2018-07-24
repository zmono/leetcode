// 84. Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

package main

import (
	"fmt"
)

type TestCase struct {
	input  []int
	output int
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights)+1)
	heights = append(heights, 0)
	max, l := 0, -1
	for i := len(heights) - 1; i >= 0; i-- {
		for l >= 0 && heights[stack[l]] >= heights[i] {
			stack = stack[:l]
			l--
		}
		stack = append(stack, i)
		l++
		for j, k := range stack[1:] {
			a := (stack[j] - i) * heights[k]
			if a > max {
				max = a
			}
		}
	}
	return max
}

func main() {
	cases := []TestCase{
		{[]int{1, 1}, 2},
		{[]int{2, 1, 2}, 3},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 0, 4, 4, 0, 0, 0, 5, 0, 0, 1}, 8},
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 1, 5, 6, 3, 3}, 12},
		{[]int{2, 3, 4, 2, 0, 2, 3, 4, 2}, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 30},
		{[]int{3, 6, 5, 7, 4, 8, 1, 0}, 20},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := largestRectangleArea(c.input); output != c.output {
			fmt.Printf("largestRectangleArea(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
