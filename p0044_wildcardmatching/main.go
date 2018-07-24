// 44. Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/description/

/*
At first I wrote a simple recursion that obviously exceeded time limit.

Then I wrote a non-recursive solution (main2.go).

I tried to optimize it, read and tried to understand https://swtch.com/~rsc/regexp/regexp1.html.

At last I wrote some code but it ended up slower than my initial non-recursive solution.

Though I'm not sure I understood and implemented everything correctly.

Then I surrendered and checked out LeetCode discussion forum, took a solution from there and rewrote it in Golang and it run 12ms.

https://leetcode.com/problems/wildcard-matching/discuss/17810/Linear-runtime-and-constant-space-solution

It's not extendable though, I don't see a clear way to add more regex features.

But it's good enough for this specific problem.
*/

package main

import (
	"fmt"
)

type TestCase struct {
	Pattern, Input string
	Output         bool
}

func isMatch(str, pattern string) bool {
	s, p, match, starIdx := 0, 0, 0, -1
	for s < len(str) {
		switch {
		case p < len(pattern) && (pattern[p] == '?' || pattern[p] == str[s]):
			s++
			p++
		case p < len(pattern) && pattern[p] == '*':
			starIdx, match = p, s
			p++
		case starIdx != -1:
			match++
			p, s = starIdx+1, match
		default:
			return false
		}
	}
	for p < len(pattern) && pattern[p] == '*' {
		p++
	}
	return p == len(pattern)
}

func main() {
	cases := []TestCase{
		{"", "", true},
		{"a", "a", true},
		{"a", "aa", false},
		{"?", "a", true},
		{"?", "aa", false},
		{"?a", "cb", false},
		{"*", "", true},
		{"*", "a", true},
		{"*a", "ba", true},
		{"*a", "cba", true},
		{"a*", "abc", true},
		{"abc*d", "abcd", true},
		{"*a*b", "adceb", true},
		{"*c?b", "cdcb", false},
		{"a*c?b", "acdcb", false},
		{"?*??", "abc", true},
		{"*?*?*?*?", "abcd", true},
		{"*?*", "c", true},
		{"a*", "a", true},
		{"a**", "a", true},
		{"a*******b", "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", false},
		{
			"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb",
			"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
			false,
		},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := isMatch(c.Input, c.Pattern); output != c.Output {
			fmt.Printf("isMatch(%#v, %#v) == %#v != %#v\n", c.Input, c.Pattern, output, c.Output)
		}
		// fmt.Println()
	}
}
