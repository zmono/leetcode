// 2. Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/description/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	of, next := 0, &ListNode{0, nil}
	r := next
	for l1 != nil || l2 != nil || of > 0 {
		r1, r2 := 0, 0
		if l1 != nil {
			r1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			r2, l2 = l2.Val, l2.Next
		}
		r := r1 + r2 + of
		of = 0
		if r > 9 {
			r = r - 10
			of = 1
		}
		next.Next = &ListNode{r, nil}
		next = next.Next
	}
	return r.Next
}

func main() {
	l1 := &ListNode{8, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{9, nil}}}
	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
