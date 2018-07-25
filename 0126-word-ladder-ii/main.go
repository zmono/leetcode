// 126. Word Ladder II
// https://leetcode.com/problems/word-ladder-ii/description/

package main

import (
	"fmt"
)

/*
My solution (findLadders2) works but is really slow.

It iterates over the whole list of words with isTransformation func to find transformations.

But the "rigth" solution uses a constant (26) way to find transformations.

findLadders is a Golang version of this solution:
https://leetcode.com/problems/word-ladder-ii/discuss/40493/The-fastest-golang-solution-using-bfs-(275ms-beats-100.00-of-golang-submissions)
*/

func isTransformation(a, b string) bool {
	d := 0
	for i := 0; i < len(a) && d < 2; i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d == 1
}

func findLadders2(b, e string, words []string) (ladders [][]string) {
	visited := make(map[string]bool)
	t := make(map[string]map[string]bool)
	t[b] = make(map[string]bool)
	for _, w := range words {
		t[w] = make(map[string]bool)
	}
	q, min := [][]string{{b}}, len(words)+1
	for len(q) > 0 {
		ladder, current := q[0], q[0][len(q[0])-1]
		q = q[1:]
		if len(t[current]) == 0 {
			for w := range t {
				if isTransformation(current, w) {
					t[current][w] = true
				}
			}
		}
		for w := range t[current] {
			if visited[w] {
				continue
			}
			ladder := append(append([]string{}, ladder...), w)
			if w == e && len(ladder) <= min {
				min, ladders = len(ladder), append(ladders, ladder)
			} else {
				q = append(q, ladder)
			}
		}
		visited[current] = true
	}
	return
}

func findLadders(b, e string, words []string) [][]string {
	nonvisited, frontier := make(map[string]struct{}), make(map[string]struct{})
	nonvisited[b], frontier[b] = struct{}{}, struct{}{}
	for _, w := range words {
		nonvisited[w] = struct{}{}
	}
	trace := make(map[string][]string)
	trace[b] = make([]string, 0)
	_, ok := frontier[e]
	for len(frontier) != 0 && !ok {
		for w := range frontier {
			delete(nonvisited, w)
		}
		nextFrontier := make(map[string]struct{})
		for w := range frontier {
			for i := range w {
				for _, c := range "abcdefghijklmnopqrstuvwxyz" {
					candidate := w[:i] + string(c) + w[i+1:]
					if _, ok := nonvisited[candidate]; ok {
						trace[candidate] = append(trace[candidate], w)
						nextFrontier[candidate] = struct{}{}
					}
				}
			}
		}
		frontier = nextFrontier
		_, ok = frontier[e]
	}

	var results [][]string
	if len(frontier) != 0 {
		backtrace(&results, trace, []string{}, e)
	}
	return results
}

func backtrace(results *[][]string, trace map[string][]string, path []string, word string) {
	if len(trace[word]) == 0 {
		*results = append(*results, append([]string{word}, path...))
	}
	for _, prev := range trace[word] {
		backtrace(results, trace, append([]string{word}, path...), prev)
	}
}

func main() {
	fmt.Println(findLadders("a", "c", []string{"a", "b", "c"}))
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(
		findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}),
	)
}
