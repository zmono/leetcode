// 25. Reverse Nodes in k-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/description/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("{%d %s}", n.Val, n.Next)
}

func reverse(head, last *ListNode) (prev *ListNode) {
	n, sentinel := head, last.Next
	for n != sentinel {
		prev, n, n.Next = n, n.Next, prev
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}

	preHead := &ListNode{Next: head}
	for i, n, prev := 0, head, preHead; n != nil; {
		if i < k-1 {
			i, n = i+1, n.Next
			continue
		}
		x := n.Next
		i, prev.Next, prev, head.Next, n, head = 0, reverse(head, n), head, x, x, x
	}
	return preHead.Next
}
