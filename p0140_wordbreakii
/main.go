// 140. Word Break II
// https://leetcode.com/problems/word-break-ii/description/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type TestCase struct {
	s      string
	dict   []string
	output []string
}

func areEqual(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)
	switch {
	case a == nil && b == nil:
		return true
	case a == nil || b == nil:
		return false
	case len(a) != len(b):
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type Node struct {
	word     string
	children map[rune]*Node
}

const EMPTY = ""

func (n *Node) String() string {
	word := "{"
	if n.word != EMPTY {
		word = fmt.Sprintf("{word: %s", n.word)
	}
	var children []string
	for c, n := range n.children {
		children = append(children, fmt.Sprintf("%s: %s", string(c), n))
	}
	childrenStr := strings.Join(children, ", ")
	if len(children) > 0 {
		if n.word != EMPTY {
			word = word + ", "
		}
		childrenStr = childrenStr + "}"
	} else {
		childrenStr = "}"
	}
	return word + childrenStr
}

func wordBreakRec(cache map[string][]string, root *Node, s string) (sentences []string) {
	if sentences, ok := cache[s]; ok {
		return sentences
	}
	defer func() { cache[s] = sentences }()
	node, ok := root, false
	for i, c := range s {
		node, ok = node.children[c]
		switch {
		case !ok:
			return
		case node.word == EMPTY:
		case s[i+1:] == EMPTY:
			sentences = append(sentences, node.word)
		default:
			for _, subsentence := range wordBreakRec(cache, root, s[i+1:]) {
				sentences = append(sentences, node.word+" "+subsentence)
			}
		}
	}
	return
}

func wordBreak(s string, words []string) []string {
	root := &Node{children: make(map[rune]*Node)}
	for _, w := range words {
		n := root
		for _, c := range w {
			if _, ok := n.children[c]; !ok {
				n.children[c] = &Node{children: make(map[rune]*Node)}
			}
			n = n.children[c]
		}
		n.word = w
	}
	return wordBreakRec(make(map[string][]string), root, s)
}

func main() {
	cases := []TestCase{
		{"abcd", []string{"a", "b", "abc", "cd"}, []string{"a b cd"}},
		{"a", []string{"a", "b"}, []string{"a"}},
		{"sand", []string{"and", "sand"}, []string{"sand"}},
		{"sand", []string{"san"}, nil},
		{"ab", []string{"a", "b", "ab", "ba"}, []string{"a b", "ab"}},
		{"aba", []string{"a", "b", "ab", "ba"}, []string{"a b a", "a ba", "ab a"}},
		{"catsand", []string{"cat", "cats", "and", "sand"}, []string{"cat sand", "cats and"}},
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cat sand dog", "cats and dog"}},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, []string{"pine apple pen apple",
			"pine applepen apple",
			"pineapple pen apple"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, nil},
		{"aaaaaaa", []string{"aaa", "aaaa"}, []string{"aaaa aaa", "aaa aaaa"}},
		{"aaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, []string{"a a a a a a", "aa a a a a", "a aa a a a", "aaa a a a", "a a aa a a", "aa aa a a", "a aaa a a", "aaaa a a", "a a a aa a", "aa a aa a", "a aa aa a", "aaa aa a", "a a aaa a", "aa aaa a", "a aaaa a", "aaaaa a", "a a a a aa", "aa a a aa", "a aa a aa", "aaa a aa", "a a aa aa", "aa aa aa", "a aaa aa", "aaaa aa", "a a a aaa", "aa a aaa", "a aa aaa", "aaa aaa", "a a aaaa", "aa aaaa", "a aaaaa"}},
		// {"aaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, nil},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, nil},
	}
	for _, c := range cases {
		fmt.Println(c)
		// wordBreak(c.s, c.dict)
		// fmt.Println("done")
		if output := wordBreak(c.s, c.dict); !areEqual(output, c.output) {
			fmt.Printf("wordBreak(%#v, %#v) == %#v != %#v\n", c.s, c.dict, output, c.output)
		}
	}
}
