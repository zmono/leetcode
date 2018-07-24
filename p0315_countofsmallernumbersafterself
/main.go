// 315. Count of Smaller Numbers After Self
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/

package main

import (
	"fmt"
	"reflect"
)

type TestCase struct{ input, output []int }

type Node struct {
	Val, count  int
	Left, Right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("{%s [%d, %d] %s}", n.Left, n.Val, n.Count, n.Right)
}

func (n *Node) Push(x int) *Node {
	switch {
	case n == nil:
		return &Node{Val: x, count: 1}
	case x < n.Val:
		n.Left = n.Left.Push(x)
	case x > n.Val:
		n.Right = n.Right.Push(x)
	}
	n.count++
	return n
}

func (n *Node) Count() int {
	if n == nil {
		return 0
	}
	return n.count
}

func (n *Node) CountLess(x int) int {
	switch {
	case n == nil:
		return 0
	case x < n.Val:
		return n.Left.CountLess(x)
	case x > n.Val:
		return n.Count() - n.Right.Count() + n.Right.CountLess(x)
	default:
		return n.Left.Count()
	}
}

func countSmaller(nums []int) []int {
	var root *Node
	output := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = root.CountLess(nums[i])
		root = root.Push(nums[i])
	}
	return output
}

func main() {
	cases := []TestCase{
		{[]int{}, []int{}},
		{[]int{1}, []int{0}},
		{[]int{1, 1}, []int{0, 0}},
		{[]int{0, 1}, []int{0, 0}},
		{[]int{2, 1}, []int{1, 0}},
		{[]int{2, 2, 1}, []int{1, 1, 0}},
		{[]int{0, 2, 1}, []int{0, 1, 0}},
		{[]int{1, 2, 1}, []int{0, 1, 0}},
		{[]int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{7, 6, 5, 4, 3, 2, 1}, []int{6, 5, 4, 3, 2, 1, 0}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{0, 0, 0, 0, 0, 0, 0}},
		{[]int{8, 6, 5, 6, 1}, []int{4, 2, 1, 1, 0}},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := countSmaller(c.input); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("countSmaller(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
