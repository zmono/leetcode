// 239. Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/description/

package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	nums   []int
	k      int
	output []int
}

type Node struct {
	val        int
	prev, next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("{%d, %s}", n.val, n.next)
}

type Deque struct {
	front, back *Node
}

func (d *Deque) String() string {
	return fmt.Sprintf("%s", d.front)
}
func (d Deque) IsEmpty() bool { return d.front == nil }
func (d *Deque) Push(x int) {
	n := &Node{val: x, prev: d.back, next: nil}
	if d.front == nil {
		d.front = n
	}
	if d.back == nil {
		d.back = n
	} else {
		d.back.next, d.back = n, n
	}
}
func (d *Deque) Peek() int     { return d.front.val }
func (d *Deque) PeekBack() int { return d.back.val }
func (d *Deque) Pop() int {
	f := d.front
	d.front = d.front.next
	if d.front == nil {
		d.back = nil
	} else {
		d.front.prev = nil
	}
	return f.val
}
func (d *Deque) PopBack() int {
	b := d.back
	d.back = d.back.prev
	if d.back == nil {
		d.front = nil
	} else {
		d.back.next = nil
	}
	return b.val
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	output := make([]int, len(nums)-k+1)
	d := Deque{}
	for i := 0; i < len(nums); i++ {
		if !d.IsEmpty() && d.Peek() < i-k+1 {
			d.Pop()
		}
		for !d.IsEmpty() && nums[i] >= nums[d.PeekBack()] {
			d.PopBack()
		}
		d.Push(i)
		if i-k+1 >= 0 {
			output[i-k+1] = nums[d.Peek()]
		}
	}
	return output
}

func main() {
	cases := []TestCase{
		{nums: []int{}, k: 0, output: nil},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, output: []int{3, 3, 5, 5, 6, 7}},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := maxSlidingWindow(c.nums, c.k); !reflect.DeepEqual(output, c.output) {
			fmt.Printf("maxSlidingWindow(%v, %v) == %v != %v\n", c.nums, c.k, output, c.output)
		}
	}
}
