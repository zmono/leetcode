// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/description/

// Было лень думать, содрал решение из форума, переписал на Go.
// Моя первая идея приблизительно похожа, но далеко.

package main

import (
	"fmt"
	"math"
)

type TestCase struct {
	s, t, output string
}

func minWindow(s string, t string) string {
	bytes := make(map[byte]int)
	for _, c := range t {
		bytes[byte(c)]++
	}
	counter, begin, end, d, head := len(t), 0, 0, math.MaxInt64, 0
	for end < len(s) {
		if bytes[s[end]] > 0 {
			counter--
		}
		bytes[s[end]]--
		end++
		for counter == 0 {
			if end-begin < d {
				head, d = begin, end-begin
			}
			if bytes[s[begin]] == 0 {
				counter++
			}
			bytes[s[begin]]++
			begin++
		}
	}
	if d == math.MaxInt64 {
		return ""
	}
	return s[head : head+d]
}

func main() {
	cases := []TestCase{
		{"ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := minWindow(c.s, c.t); output != c.output {
			fmt.Printf("minWindow(%v, %v) == %v != %v\n", c.s, c.t, output, c.output)
		}
	}
}
