package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count == 1 {
			fmt.Print(a[i], " ")
		}
	}
}
