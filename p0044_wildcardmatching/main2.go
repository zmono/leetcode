// 44. Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/description/

package main

import (
	"fmt"
)

type TestCase struct {
	Pattern, Input string
	Output         bool
}

type Regex struct {
	p        []rune
	frontier map[int]int
	match    bool
}

func (r *Regex) next(c rune) {
	r.match = false
	for n, count := range r.frontier {
		if count == 0 {
			continue
		}
		r.frontier[n]--
		if !(r.p[n] == '*' || r.p[n] == '?' || r.p[n] == c) {
			continue
		}
		if r.p[n] == '*' {
			r.frontier[n]++
		}
		if len(r.p) <= n+1 {
			r.match = true
			continue
		}
		r.frontier[n+1]++
		if r.p[n+1] == '*' {
			r.match = r.match || len(r.p) <= n+2
		} else if r.p[n] == '*' && (r.p[n+1] == '?' || r.p[n+1] == c) {
			if len(r.p) <= n+2 {
				r.match = true
			} else {
				r.frontier[n+2]++
				r.match = r.match || r.p[n+2] == '*' && len(r.p) <= n+3
			}
		}
	}
}

func cleanPattern(p string) string {
	var clean []rune
	var prev rune
	for _, c := range p {
		if !(prev == '*' && c == '*') {
			prev, clean = c, append(clean, c)
		}
	}
	return string(clean)
}

func isMatch(s, p string) bool {
	p = cleanPattern(p)
	if p == "" {
		return s == ""
	}
	regex := Regex{
		p:        []rune(p),
		frontier: map[int]int{0: 1},
		match:    p == "*",
	}
	for _, r := range s {
		regex.next(r)
	}
	return regex.match
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
		{"*", "aa", true},
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
		if output := isMatch(c.Input, c.Pattern); output != c.Output {
			fmt.Printf("isMatch(%#v, %#v) == %#v != %#v\n", c.Input, c.Pattern, output, c.Output)
		}
	}
}
