package main

import "fmt"

func setBit(n int64, i int, value bool) int64 {
	if value {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	var x int64 = 0
	x = setBit(x, 2, true)
	x = setBit(x, 0, true)
	x = setBit(x, 2, false)
	fmt.Println(x, x)
}
