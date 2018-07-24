// 124. Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

package main

import (
	"fmt"
)

type TestCase struct {
	input  []*Int
	output int
}

type Int struct{ i int }

func buildTree(nums []*Int, level uint, i int) *TreeNode {
	idx := 1<<level + i - 1
	if idx >= len(nums) {
		return nil
	}
	n := nums[idx]
	if n == nil {
		return nil
	}
	return &TreeNode{
		Val:   n.i,
		Left:  buildTree(nums, level+1, 2*i),
		Right: buildTree(nums, level+1, 2*i+1),
	}
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("{%d, %s, %s}", n.Val, n.Left, n.Right)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxPathSumRec(t *TreeNode) (int, int) {
	switch {
	case t == nil:
		return 0, 0
	case t.Left == nil && t.Right == nil:
		return t.Val, t.Val
	case t.Right == nil:
		l, _ := maxPathSumRec(t.Left)
		return t.Val + max(0, l), max(t.Val, l)
	case t.Left == nil:
		r, _ := maxPathSumRec(t.Right)
		return t.Val + max(0, r), max(t.Val, r)
	default:
		l, ln := maxPathSumRec(t.Left)
		r, rn := maxPathSumRec(t.Right)
		n := max(l+t.Val+r, max(max(l, r), max(ln, rn)))
		return t.Val + max(0, max(l, r)), n
	}
}

func maxPathSum(root *TreeNode) int {
	m, n := maxPathSumRec(root)
	return max(m, n)
}

func main() {
	cases := []TestCase{
		{[]*Int{{0}}, 0},
		{[]*Int{{-1}}, -1},
		{[]*Int{{10}, {-2}, {-3}}, 10},
		{[]*Int{{1}, {2}, {3}}, 6},
		{[]*Int{{0}, {0}, {0}}, 0},
		{[]*Int{{-7}, {-1}, {-1}}, -1},
		{[]*Int{{-7}, {10}, {10}}, 13},
		{[]*Int{{-10}, {9}, {20}, nil, nil, {15}, {7}}, 42},
		{[]*Int{{1}, {-2}, {-3}, {1}, {3}, {-2}, nil, {-1}}, 3},
		{[]*Int{{-6}, nil, {3}, nil, nil, {2}}, 5},
		{[]*Int{{-4},
			{-5}, {4},
			{5}, nil, {3}, nil,
			{-2}, nil, nil, nil, {6}, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, {2}, {-2}}, 15},
	}
	for _, c := range cases {
		// fmt.Println(c)
		root := buildTree(c.input, 0, 0)
		if output := maxPathSum(root); output != c.output {
			fmt.Printf("maxPathSum(%s) == %v != %v\n", root, output, c.output)
		}
	}
}
