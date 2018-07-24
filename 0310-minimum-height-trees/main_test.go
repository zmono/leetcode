package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	edges  [][]int
	output []int
}

func TestFindMinHeightTrees(t *testing.T) {
	cases := []TestCase{
		{0, [][]int{}, nil},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{0, 1}}, []int{1, 0}},
		{3, [][]int{{0, 1}, {0, 2}}, []int{0}},
		{3, [][]int{{0, 1}, {1, 2}}, []int{1}},
		{4, [][]int{{1, 0}, {1, 2}, {1, 3}}, []int{1}},
		{6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}, []int{4, 3}},
	}
	for _, c := range cases {
		output := findMinHeightTrees(c.n, c.edges)
		if output == nil && c.output == nil {
			continue
		}
		if !reflect.DeepEqual(output, c.output) {
			t.Errorf("findMinHeightTrees(%d, %v) == %v != %v\n", c.n, c.edges, output, c.output)
		}
	}
}
