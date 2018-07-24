// 32. Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/description/

package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	input  string
	output int
}

func longestValidParentheses(s string) (max int) {
	G := []int{}
	L := func() int { return len(G) }
	for _, p := range s {
		if p == '(' {
			G = append(G, 0)
			continue
		}
		if L() == 0 {
			continue
		}
		if !(G[L()-1] == 0 || L() > 1 && G[L()-2] == 0) {
			G = nil
			continue
		}

		G[L()-1] += 2
		once := G[L()-1] > 2
		for once || L() > 1 && G[L()-2] > 0 {
			once = false
			G[L()-2] += G[L()-1]
			G = G[:L()-1]
		}
		if G[L()-1] > max {
			max = G[L()-1]
		}
	}
	return max
}

func main() {
	cases := []TestCase{
		{"(()(()()))))", 10},
		{"", 0},
		{")", 0},
		{"(", 0},
		{"()", 2},
		{"()))", 2},
		{"((()", 2},
		{"(()", 2},
		{")()())", 4},
		{"(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((", 0},
		{"()(()", 2},
		{"()(()()", 4},
		{"()(()()(()()()", 6},
		{")()())()()(", 4},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := longestValidParentheses(c.input); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("longestValidParentheses(\"%s\") == %v != %v\n", c.input, output, c.output)
		}
	}
}
