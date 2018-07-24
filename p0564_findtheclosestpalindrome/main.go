// 564. Find the Closest Palindrome
// https://leetcode.com/problems/find-the-closest-palindrome/description/

package main

import (
	"fmt"
	"strconv"
)

type TestCase struct {
	Input, Output string
}

// https://stackoverflow.com/a/1754209
func reversed(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func atoi(n string) int64 {
	x, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func itoa(x int64) string {
	return fmt.Sprintf("%d", x)
}

func nearestPalindromicRec(start int64, n string) string {
	nInt, _ := strconv.ParseInt(n, 10, 64)
	var candidates []int64
	digits := make(map[rune]int)
	for _, d := range n {
		digits[d] += 1
	}
	var d rune
	for d = range digits {
		break
	}
	switch {
	case len(digits) == 1 && d == '9' && nInt > 10:
		return itoa(nInt + 2)
	case len(n)%2 == 0:
		half := n[:len(n)/2]
		candidates = append(candidates, atoi(half+reversed(half)))
		halfInt, _ := strconv.ParseInt(half, 10, 64)
		incrementedHalf := fmt.Sprintf("%d", halfInt+1)
		candidates = append(candidates, atoi(incrementedHalf+reversed(incrementedHalf)))
		decrementedHalf := fmt.Sprintf("%d", halfInt-1)
		if len(decrementedHalf) < len(half) || decrementedHalf == "0" {
			candidates = append(candidates, atoi(decrementedHalf+decrementedHalf+"9"))
		} else {
			candidates = append(candidates, atoi(decrementedHalf+reversed(decrementedHalf)))
		}
	default:
		half := n[:len(n)/2+1]
		candidates = append(candidates, atoi(half+reversed(half[:len(half)-1])))
		halfInt, _ := strconv.ParseInt(half, 10, 64)
		incrementedHalf := fmt.Sprintf("%d", halfInt+1)
		candidates = append(candidates, atoi(incrementedHalf+reversed(incrementedHalf[:len(incrementedHalf)-1])))
		decrementedHalf := fmt.Sprintf("%d", halfInt-1)
		if len(decrementedHalf) < len(half) {
			candidates = append(candidates, atoi(decrementedHalf+reversed(decrementedHalf)))
		} else {
			candidates = append(candidates, atoi(decrementedHalf+reversed(decrementedHalf[:len(decrementedHalf)-1])))
		}
	}
	bestDiff, best := int64(1<<63-1), int64(1<<63-1)
	for _, c := range candidates {
		d := abs(c - nInt)
		if d != 0 && (d < bestDiff || d == bestDiff && c < best) {
			bestDiff, best = d, c
		}
	}
	return itoa(best)
}

func nearestPalindromic(n string) string {
	x, _ := strconv.ParseInt(n, 10, 64)
	return nearestPalindromicRec(x, n)
}

// надо определить все возможные варианты построения палиндромов, а затем из полученных просто выбрать оптимальный
// abcd: abba, a(b+1)(b+1)a, a(b-1)(b-1)a,

// думаю, всё-таки надо использовать рекурсию с передачей стартового значения

func main() {
	cases := []TestCase{
		{"1", "0"},
		{"2", "1"},
		{"9", "8"},
		{"10", "9"},
		{"11", "9"},
		{"12", "11"},
		{"14", "11"},
		{"19", "22"},
		{"44", "33"},
		{"48", "44"},
		{"59", "55"},
		{"66", "55"},
		{"88", "77"},
		{"99", "101"},
		{"100", "99"},
		{"123", "121"},
		{"202", "212"},
		{"909", "919"},
		{"989", "979"},
		{"998", "999"},
		{"999", "1001"},
		{"1000", "999"},
		{"1001", "999"},
		{"1283", "1331"},
		{"2002", "1991"},
		{"9889", "9779"},
		{"10001", "9999"},
		{"10801", "10701"},
		{"10899", "10901"},
		{"10999", "11011"},
		{"11011", "11111"},
		{"100001", "99999"},
		{"108001", "107701"},
		{"938671", "938839"},
		{"749271017381326528", "749271017710172947"},
	}
	for _, c := range cases {
		if output := nearestPalindromic(c.Input); output != c.Output {
			fmt.Printf("nearestPalindromic(%#v) == %#v != %#v\n", c.Input, output, c.Output)
		}
	}
}
