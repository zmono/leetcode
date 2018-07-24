// 72. Edit Distance
// https://leetcode.com/problems/edit-distance/description/

package main

import (
	"fmt"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://bitbucket.org/clearer/iosifovich/overview
func minDistance(a, b string) int {
	var i int
	for i = 0; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
	}
	a, b = a[i:], b[i:]
	if len(a) > len(b) {
		a, b = b, a
	}
	buf := make([]int, len(b)+1)
	for i := range buf {
		buf[i] = i
	}
	for i := 0; i < len(a); i++ {
		t := buf[0]
		buf[0]++
		for j := 0; j < len(buf)-1; j++ {
			cost := 0
			if a[i] != b[j] {
				cost = 1
			}
			//          insert  delete       substitute
			t = Min(Min(buf[j], buf[j+1])+1, t+cost)
			buf[j+1], t = t, buf[j+1]
		}
	}
	return buf[len(buf)-1]
}

func main() {
	fmt.Println()
}
