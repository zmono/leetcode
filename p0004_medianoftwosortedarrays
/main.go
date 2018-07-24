// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
*/

package main

import (
	"fmt"
)

type TestCase struct {
	Nums1, Nums2 []int
	Output       float64
}

func bisectLeft(n int, nums []int) int {
	i := len(nums) / 2
	switch {
	case len(nums) == 0:
		return 0
	case nums[i] == n:
		return i
	case nums[i] > n:
		return bisectLeft(n, nums[:i])
	default:
		return bisectLeft(n, nums[i+1:]) + i + 1
	}
}

func takeNth(n int, nums1, nums2 []int) (nums []int, k int, ok bool) {
	switch {
	case len(nums1) == 0 && len(nums2) == 0:
		return nil, 0, false
	case len(nums1) == 0:
		return takeNth(n, nums2, nums1)
	case len(nums2) == 0:
		return nums1, n, true
	}

	i := len(nums1) / 2
	j := bisectLeft(nums1[i], nums2)
	switch {
	case i+j == n:
		return nums1, i, true
	case i+j < n:
		return takeNth(n-(i+1), nums1[i+1:], nums2)
	default:
		return takeNth(n, nums1[:i], nums2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	i1, i2 := l/2, l/2-1
	nums1_, n1, ok := takeNth(i1, nums1, nums2)
	switch {
	case !ok:
		return 0
	case l%2 == 1:
		return float64(nums1_[n1])
	default:
		if nums2_, n2, ok := takeNth(i2, nums1, nums2); ok {
			return float64(nums1_[n1]+nums2_[n2]) / 2
		}
	}
	return 0
}

func main() {
	// fmt.Println(takeNth(0, []int{2}, []int{}))
	// fmt.Println(takeNth(0, []int{}, []int{2}))
	// fmt.Println(takeNth(1, []int{}, []int{2, 5}))
	// fmt.Println(takeNth(3, []int{1, 7}, []int{2, 5}))
	// fmt.Println(takeNth(0, []int{1, 3}, []int{2}))
	// fmt.Println(takeNth(2, []int{1, 2}, []int{3, 4}))
	// fmt.Println(takeNth(7, []int{1, 4, 7}, []int{5, 6}))
	// fmt.Println(takeNth(2, []int{5, 6}, []int{1, 4, 7}))
	// fmt.Println(takeNth(5, []int{1, 6, 7}, []int{2, 3, 5}))
	cases := []TestCase{
		{[]int{}, []int{}, 0.0},
		{[]int{}, []int{2}, 2.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{1}, []int{1}, 1.0},
		{[]int{1, 2}, []int{2, 3}, 2.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 1, 7}, []int{2, 3, 6}, 2.5},
		{[]int{1, 3, 5, 7, 9, 11, 13}, []int{2, 6, 14}, 6.5},
		{[]int{1, 3, 5, 7, 9, 11, 13}, []int{2, 4, 6, 8, 10, 14}, 7},
		{[]int{-10, -1, 1, 0}, []int{}, 0},
	}
	for _, c := range cases {
		// fmt.Println(c)
		if output := findMedianSortedArrays(c.Nums1, c.Nums2); output != c.Output {
			fmt.Printf("findMedianSortedArrays(%#v, %#v) == %#v != %#v\n", c.Nums1, c.Nums2, output, c.Output)
		}
		// fmt.Println()
	}
}
