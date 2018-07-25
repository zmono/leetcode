// 85. Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/

package main

import (
	"fmt"
)

func join(A, B [][2]int) ([][2]int, int) {
	switch {
	case A == nil && B == nil:
		return [][2]int{{1, 1}}, 1
	case A == nil:
		m := B[0][1]
		for _, b := range B {
			if b[1] > m {
				m = b[1]
			}
		}
		return [][2]int{{1, m + 1}}, m + 1
	case B == nil:
		m := A[0][0]
		for _, a := range A {
			if a[0] > m {
				m = a[0]
			}
		}
		return [][2]int{{m + 1, 1}}, m + 1
	}

	m := make(map[int]int)
	for _, a := range A {
		if a[1] == 1 {
			m[a[0]+1] = 1
			continue
		}
		for _, b := range B {
			if a[0]-b[0] < 0 && b[1]-a[1] < 0 && m[a[0]+1] < b[1]+1 {
				m[a[0]+1] = b[1] + 1
			}
		}
	}
	for _, b := range B {
		if b[0] == 1 {
			m[1] = b[1] + 1
		}
	}
	i, max, pairs := 0, 0, make([][2]int, len(m))
	for y, x := range m {
		i, pairs[i] = i+1, [2]int{y, x}
		if t := y * x; t > max {
			max = t
		}
	}
	return pairs, max
}

func MaximalRectangle1(matrix [][]byte) (max int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, memo := 0, make([][][2]int, len(matrix[0]))
	for i, row := range matrix {
		for j, b := range row {
			if b != '1' {
				memo[j] = nil
				continue
			}
			var upper, left [][2]int
			if i > 0 {
				upper = memo[j]
			}
			if j > 0 {
				left = memo[j-1]
			}
			memo[j], m = join(upper, left)
			if m > max {
				max = m
			}
		}
	}
	return
}

func MaximalRectangle(matrix [][]byte) (max int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	left, right, height := make([]int, n), make([]int, n), make([]int, n)
	for i := range right {
		right[i] = n
	}
	for i := 0; i < len(matrix); i++ {
		curLeft, curRight := 0, n
		for j := 0; j < n; j++ {
			if matrix[i][j] != '1' {
				height[j], left[j], curLeft = 0, 0, j+1
			} else {
				height[j]++
				if curLeft > left[j] {
					left[j] = curLeft
				}
			}

			k := n - j - 1
			if matrix[i][k] != '1' {
				right[k], curRight = n, k
			} else if curRight < right[k] {
				right[k] = curRight
			}
		}
		for j := 0; j < n; j++ {
			if r := (right[j] - left[j]) * height[j]; r > max {
				max = r
			}
		}
	}
	return max
}

func main() {
	fmt.Println("85. Maximal Rectangle")
}
