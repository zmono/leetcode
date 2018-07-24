// 295. Find Median from Data Stream
// https://leetcode.com/problems/find-median-from-data-stream/description/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Peek() int          { return h[0] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old, last := *h, len(*h)-1
	*h = old[:last]
	return old[last]
}

type MedianFinder struct{ less, greater IntHeap }

func Constructor() MedianFinder { return MedianFinder{} }

func (f *MedianFinder) AddNum(n int) {
	a, b := &f.greater, &f.less
	if len(f.greater) != 0 && n <= f.greater[0] {
		a, b, n = b, a, -n
	}
	heap.Push(a, n)
	if len(*a)-len(*b) > 1 {
		heap.Push(b, -heap.Pop(a).(int))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	ll, gl := len(f.less), len(f.greater)
	switch {
	case ll > gl:
		return float64(-f.less[0])
	case ll < gl:
		return float64(f.greater[0])
	case ll == 0:
		return 0
	default:
		return float64(f.greater[0]-f.less[0]) / 2
	}
}

func main() {
	f := Constructor()
	f.AddNum(1)
	f.AddNum(2)
	f.AddNum(3)
	f.AddNum(4)
	f.AddNum(5)
	fmt.Println(f.FindMedian())
	fmt.Println(f.less, f.greater)
}
