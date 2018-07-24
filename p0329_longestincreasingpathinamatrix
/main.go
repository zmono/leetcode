// 329. Longest Increasing Path in a Matrix
// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/

package main

import (
	"fmt"
	"reflect"
	"sort"
)

type TestCase struct {
	input  [][]int
	output int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestIncreasingPath(matrix [][]int) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}
	uniq := []int{}
	indexes := make(map[int][][2]int)
	lengths := make([][]int, N)
	for i, row := range matrix {
		lengths[i] = make([]int, M)
		for j, x := range row {
			if _, ok := indexes[x]; !ok {
				uniq = append(uniq, x)
			}
			indexes[x] = append(indexes[x], [2]int{i, j})
			lengths[i][j] = 1
		}
	}
	sort.Ints(uniq)
	max := 1
	for i := len(uniq) - 2; i >= 0; i-- {
		for _, n := range indexes[uniq[i]] {
			m, x, y := 1, n[0], n[1]
			l, n := lengths[x][y], matrix[x][y]
			for _, c := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				cx, cy := x+c[0], y+c[1]
				valid := 0 <= cx && cx < N && 0 <= cy && cy < M
				if valid && matrix[cx][cy] > n {
					m = Max(m, l+lengths[cx][cy])
				}
			}
			lengths[x][y] = m
			max = Max(m, max)
		}
	}
	return max
}

func main() {
	cases := []TestCase{
		{[][]int{}, 0},
		{[][]int{{}}, 0},
		{[][]int{{1}}, 1},
		{[][]int{
			{1},
			{1}}, 1},
		{[][]int{{9, 9, 4}}, 2},
		{[][]int{{3, 3, 14, 2, 17, 12, 5}}, 3},
		{[][]int{
			{9, 9, 4},
			{6, 6, 8},
			{2, 1, 1}}, 4},
		{[][]int{
			{3, 4, 5},
			{3, 2, 6},
			{2, 2, 1}}, 4},
		{[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1}}, 1},
		{[][]int{
			{7, 7, 5},
			{2, 4, 6},
			{8, 2, 0}}, 4},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := longestIncreasingPath(c.input); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("longestIncreasingPath(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
