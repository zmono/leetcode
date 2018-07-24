// 218. The Skyline Problem
// https://leetcode.com/problems/the-skyline-problem/description/

package main

import (
	"fmt"
	"reflect"
	"sort"
)

type TestCase struct {
	input  [][]int
	output [][]int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cut(a, b, h int, overlays map[[2]int]struct{}) [][3]int {
	var dents [][2]int
	min, max := a, b
	for s := range overlays {
		switch {
		case s[1] < a || b < s[0]: // s doesn't overlay [a, b] at all
			continue
		case s[0] <= a: // s overlays [a, b] on the left side
			if b <= s[1] { // s completely overlays [a, b]
				return nil
			}
			a, min = s[1], Min(min, s[0])
		case b <= s[1]: // s overlaps [a, b] on the right side
			b, max = s[0], Max(max, s[1])
		case len(dents) > 0 && dents[len(dents)-1][1] == s[0]:
			dents[len(dents)-1][1] = s[1]
		default: // s is somewhere between a and b
			dents = append(dents, s)
		}
		delete(overlays, s)
	}
	overlays[[2]int{min, max}] = struct{}{}
	if len(dents) == 0 {
		return [][3]int{{a, b, h}}
	}
	sort.Slice(dents, func(i, j int) bool { return dents[i][0] < dents[j][0] })
	output := [][3]int{{a, dents[0][0], h}}
	for i, d := range dents[1:] {
		output = append(output, [3]int{dents[i][1], d[0], h})
	}
	return append(output, [3]int{dents[len(dents)-1][1], b, h})
}

func joinSegments(segments [][3]int) [][3]int {
	i, joined := 0, [][3]int{segments[0]}
	for _, s := range segments[1:] {
		if joined[i][1] < s[0] {
			i, joined = i+1, append(joined, [3]int{joined[i][1], s[0], 0})
		}
		if joined[i][1] == s[0] && joined[i][2] == s[2] {
			joined[i][1] = s[1]
		} else {
			i, joined = i+1, append(joined, s)
		}
	}
	return joined
}

func getSkyline(buildings [][]int) (output [][]int) {
	if len(buildings) == 0 {
		return nil
	}
	rightmost := buildings[0][0]
	sort.SliceStable(buildings, func(i, j int) bool { return buildings[i][2] > buildings[j][2] })
	segments, overlays := [][3]int{}, make(map[[2]int]struct{})
	for _, b := range buildings {
		segments = append(segments, cut(b[0], b[1], b[2], overlays)...)
	}
	sort.Slice(segments, func(i, j int) bool { return segments[i][0] < segments[j][0] })
	for _, s := range joinSegments(segments) {
		rightmost, output = Max(rightmost, s[1]), append(output, []int{s[0], s[2]})
	}
	return append(output, []int{rightmost, 0})
}

func main() {
	// visited := map[[2]int]struct{}{
	// 	// [2]int{2, 4}: struct{}{},
	// 	// [2]int{3, 4}:  struct{}{},
	// 	// [2]int{1, 40}: struct{}{},
	// }
	// fmt.Println(cut(1, 2, 1, visited), visited)
	cases := []TestCase{
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{1, 2, 1}},
			[][]int{{1, 1}, {2, 0}},
		},
		{
			[][]int{{0, 3, 2}, {1, 2, 1}},
			[][]int{{0, 2}, {3, 0}},
		},
		{
			[][]int{{1, 2, 1}, {2, 3, 1}},
			[][]int{{1, 1}, {3, 0}},
		},
		{
			[][]int{{1, 3, 1}, {2, 4, 1}},
			[][]int{{1, 1}, {4, 0}},
		},
		{
			[][]int{{0, 3, 1}, {1, 2, 2}},
			[][]int{{0, 1}, {1, 2}, {2, 1}, {3, 0}},
		},
		{
			[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := getSkyline(c.input); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("getSkyline(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
