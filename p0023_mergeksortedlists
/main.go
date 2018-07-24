// 23. Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/description/

package main

import (
	"fmt"
	"math"
)

type TestCase struct {
	input  []*ListNode
	output *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("{%d %s}", n.Val, n.Next)
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := &ListNode{}
	start := n
	for {
		min, minListIndex := math.MaxInt64, -1
		for i, l := range lists {
			if l != nil && l.Val < min {
				min, minListIndex = l.Val, i
			}
		}
		if minListIndex == -1 {
			break
		}
		n.Next = lists[minListIndex]
		n = n.Next
		n.Val = min
		lists[minListIndex] = lists[minListIndex].Next
	}
	return start.Next
}

func main() {
	cases := []TestCase{
		{[]*ListNode{&ListNode{4, nil}, &ListNode{8, nil}}, &ListNode{4, &ListNode{8, nil}}},
	}
	for _, c := range cases {
		fmt.Println(c)
		if output := mergeKLists(c.input); output != c.output {
			fmt.Printf("mergeKLists(%v) == %v != %v\n", c.input, output, c.output)
		}
	}
}
