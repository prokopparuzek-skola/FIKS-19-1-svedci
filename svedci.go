package main

import "fmt"

func split(friends [][]int) (ans []int) {
	ans = make([]int, len(friends)-1)
	var d1, d2 map[int]bool
	d1 = make(map[int]bool)
	d2 = make(map[int]bool)

	for i := 1; i < len(friends); i++ {
		d1[i] = true
	}

	for count := 0; count < 200; count++ {
		for i := 1; i < len(friends); i++ {
			f := friends[i]
			var cnt1, cnt2 int
			for _, neigh := range f {
				if d1[neigh] == true {
					cnt1++
				}
			}
			for _, neigh := range f {
				if d2[neigh] == true {
					cnt2++
				}
			}
			if cnt1 > len(f)/2 && cnt2 <= len(f)/2 {
				d2[i] = true
				d1[i] = false
			}
		}
		for i := 1; i < len(friends); i++ {
			f := friends[i]
			var cnt1, cnt2 int
			for _, neigh := range f {
				if d1[neigh] == true {
					cnt1++
				}
			}
			for _, neigh := range f {
				if d2[neigh] == true {
					cnt2++
				}
			}
			if cnt2 > len(f)/2 && cnt1 <= len(f)/2 {
				d1[i] = true
				d2[i] = false
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		t := i + 1
		if d1[t] == true {
			ans[i] = 1
		} else {
			ans[i] = 2
		}
	}
	return
}

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		var N, M int
		var friends [][]int
		fmt.Scanf("%d%d", &N, &M)
		friends = make([][]int, N+1)
		for j := 1; j <= N; j++ {
			friends[j] = make([]int, 0)
		}
		for j := 0; j < M; j++ {
			var from, to int
			fmt.Scanf("%d%d", &from, &to)
			friends[from] = append(friends[from], to)
			friends[to] = append(friends[to], from)
		}
		ans := split(friends)
		if len(ans) == 1 {
			fmt.Printf("%d\n", ans[0])
		} else {
			for j := 0; j < len(ans)-1; j++ {
				fmt.Printf("%d ", ans[j])
			}
			fmt.Printf("%d\n", ans[len(ans)-1])
		}
	}
}
