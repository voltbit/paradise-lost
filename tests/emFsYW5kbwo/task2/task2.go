package main

import "fmt"

func Solution(a []int, b []int) int {
	n := len(a)
	adj := make(map[int]int)
	for i := 0; i < n; i++ {
		adj[a[i]] = b[i]
	}
	reachedBy := make(map[int][]int)
	for i := 0; i < n; i++ {
		current := a[i]
		visited := make(map[int]bool)
		for {
			if !visited[current] {
				visited[current] = true
				if next, found := adj[current]; found {
					current = next
				} else {
					break
				}
				reachedBy[current] = append(reachedBy[current], i)
				if len(reachedBy[current]) == n {
					return current
				}
			} else {
				break
			}
		}
		visited = nil
	}
	for node, nodes := range reachedBy {
		if len(nodes) == len(a) {
			return node
		}
	}
	return -1
}

func main() {
	fmt.Println(Solution([]int{0, 1, 2, 4, 5}, []int{2, 3, 3, 3, 2}))
	fmt.Println(Solution([]int{1, 2, 3}, []int{0, 0, 0}))
	fmt.Println(Solution([]int{2, 3, 3, 4}, []int{1, 1, 0, 0}))
}
