package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		var N, M int
		var friends [][]int
		fmt.Scanf("%d%d", &N, &M)
		friends = make([][]int, N)
		for j := 0; j < N; j++ {
			friends[j] = make([]int, 0)
		}
		for j := 0; j < N; j++ {
			var from, to int
			fmt.Scanf("%d%d", &from, &to)
			friends[from] = append(friends[from], to)
			friends[to] = append(friends[to], from)
		}
	}
}
