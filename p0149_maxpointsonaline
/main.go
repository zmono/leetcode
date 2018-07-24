// 149. Max Points on a Line
// https://leetcode.com/problems/max-points-on-a-line/description/

package main

import (
	"fmt"
)

/*
Input: [[1,1],[2,2],[3,3]]
Output: 3
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4

Input: [[-1,0],[3,2],[5,3],[4,1],[2,3],[7,4]]
Output: 4
	  ^
	  |
	  |                    o
	  |     o        o
	  |        o
	  |           o
---o--+----------------------->
  -1  0  1  2  3  4  5  6  7
*/

type Point struct {
	X int
	Y int
}

func getSlope(p, p2 Point) Point {
	x, y := p2.X-p.X, p2.Y-p.Y
	switch {
	case y == 0:
		return Point{1, 0}
	case x == 0:
		return Point{0, 1}
	case y < 0:
		x, y = -x, -y
	}
	a, b := x, y
	if a < 0 {
		a = -x
	}
	for b > 0 {
		a, b = b, a%b
	}
	return Point{x / a, y / a}
}

func maxPoints(points []Point) (max int) {
	counts := make(map[Point]int)
	uniq := make([]Point, 0, len(points))
	for _, p := range points {
		counts[p] += 1
		if counts[p] == 1 {
			uniq = append(uniq, p)
		}
		if counts[p] > max {
			max = counts[p]
		}
	}
	slopes := make(map[[2]Point]int, len(points)*len(points)/2)
	for i, p := range uniq {
		for _, p2 := range uniq[i+1:] {
			slopes[[2]Point{p, getSlope(p, p2)}] += counts[p2]
		}
	}
	for pointSlope, count := range slopes {
		c := count + counts[pointSlope[0]]
		if c > max {
			max = c
		}
	}
	return max
}

func main() {
	fmt.Println("149. Max Points on a Line")
	fmt.Println(maxPoints([]Point{{-1, 0}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {7, 4}}))
	fmt.Println(maxPoints([]Point{{3, 2}, {5, 3}, {4, 1}, {2, 3}, {7, 4}}))
	fmt.Println(maxPoints([]Point{{1, 2}, {4, 2}}))
	fmt.Println(maxPoints([]Point{{1, 2}, {1, 2}}))
	fmt.Println(maxPoints([]Point{{1, 2}, {0, 0}, {1, 2}}))
	fmt.Println(maxPoints([]Point{{1, 2}}))
	fmt.Println(maxPoints([]Point{}))
	fmt.Println(maxPoints([]Point{{0, 0}, {94911151, 94911150}, {94911152, 94911151}}))
	fmt.Println(maxPoints([]Point{{0, 9}, {138, 429}, {115, 359}, {115, 359}, {-30, -102}, {230, 709}, {-150, -686}, {-135, -613}, {-60, -248}, {-161, -481}, {207, 639}, {23, 79}, {-230, -691}, {-115, -341}, {92, 289}, {60, 336}, {-105, -467}, {135, 701}, {-90, -394}, {-184, -551}, {150, 774}}))
	fmt.Println(maxPoints([]Point{{-4, 1}, {-7, 7}, {-1, 5}, {9, -25}}))
	fmt.Println(maxPoints([]Point{{2, 3}, {3, 3}, {-5, 3}}))
}
