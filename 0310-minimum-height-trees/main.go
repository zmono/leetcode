// 310. Minimum Height Trees
// https://leetcode.com/problems/minimum-height-trees/description/

package main

func longestPath(adj map[int][]int, i int) []int {
	longest, adji := []int{}, adj[i]
	adj[i] = nil
	for _, n := range adji {
		if l := adj[n]; l == nil {
			continue
		}
		if p := longestPath(adj, n); len(p) > len(longest) {
			longest = p
		}
	}
	adj[i] = adji
	return append(longest, i)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	}
	adj := make(map[int][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	p := longestPath(adj, 0)
	p = longestPath(adj, p[0])
	if len(p)%2 == 1 {
		return []int{p[len(p)/2]}
	}
	return []int{p[len(p)/2], p[len(p)/2-1]}
}
