// 10. Regular Expression Matching
// https://leetcode.com/problems/regular-expression-matching/description/
/*
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
*/

package main

import (
	"fmt"
)

type TestCase struct {
	Pattern, Input string
	Output         bool
}

type Regex struct {
	start                  *State
	frontier, nextFrontier map[*State]struct{}
}

type State struct {
	c          rune
	out1, out2 *State
}

func (s *State) String() string {
	out2 := "self"
	if s != s.out2 {
		out2 = fmt.Sprintf("%s", s.out2)
	}
	return fmt.Sprintf("{c: %s, out1: %s, out2: %s}", string(s.c), s.out1, out2)
}

var matchstate = &State{}

func NewRegex(p string) *Regex {
	r, runes, prev := &Regex{}, []rune(p), &State{}
	prestart := prev
	for i := 0; i < len(runes); {
		s := &State{c: runes[i]}
		if i+1 < len(runes) && runes[i+1] == '*' {
			i, s.out2 = i+1, s
		}
		i, prev.out1, prev = i+1, s, s
	}
	prev.out1 = matchstate
	r.start = prestart.out1
	return r
}

func (r *Regex) addState(s *State) {
	if s == nil {
		return
	}
	if s.out2 == s {
		r.addState(s.out1)
	}
	r.nextFrontier[s] = struct{}{}
}

func (r *Regex) next(c rune) {
	r.nextFrontier = make(map[*State]struct{})
	for s := range r.frontier {
		if s.c == c || s.c == '.' {
			r.addState(s.out1)
			r.addState(s.out2)
		}
	}
	r.frontier = r.nextFrontier
}

func (r *Regex) Match(s string) bool {
	r.nextFrontier = make(map[*State]struct{})
	r.addState(r.start)
	r.frontier = r.nextFrontier
	for _, c := range s {
		r.next(c)
	}
	_, ok := r.nextFrontier[matchstate]
	return ok
}

func isMatch(s string, p string) bool {
	r := NewRegex(p)
	return r.Match(s)
}

func main() {
	cases := []TestCase{
		{"", "", true},
		{"a", "a", true},
		{"a", "aa", false},
		{".", "a", true},
		{".", "aa", false},
		{".a", "cb", false},
		{"a*", "", true},
		{"a*", "a", true},
		{"a*", "aa", true},
		{"a*b", "b", true},
		{"a*b", "ab", true},
		{"a*b", "aab", true},
		{"a*b", "aaab", true},
		{"a*b", "aaa", false},
		{"c*a*b", "aab", true},
		{"abc*d", "abcd", true},
		{"abc*d", "abd", true},
		{"abc*d", "abcdd", false},
		{"a*c.b", "cdb", true},
		{".a*..", "xyz", true},
		{".a*..", "xayz", true},
		{"a*a*a*a*", "", true},
		{"a*a*a*a*", "a", true},
		{"a*a*a*a*", "aaaaaaaaaa", true},
		{"a*a*a*a*", "aaaaaaaaaab", false},
		{"mis*is*p*.", "mississippi", false},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := isMatch(c.Input, c.Pattern); output != c.Output {
			fmt.Printf("isMatch(%#v, %#v) == %#v != %#v\n", c.Input, c.Pattern, output, c.Output)
		}
		// fmt.Println()
	}
}
