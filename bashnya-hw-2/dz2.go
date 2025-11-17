// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name int
// 	fmt.Scan(&name)
// 	if name >= 12307 {
// 		fmt.Println("Число не соответствует изначальному условию, попробуйте ещё раз")
// 		return
// 	}
// 	for name < 12307 {
// 		if name < 0 {
// 			name *= -1
// 		}
// 		if name%7 == 0 {
// 			name *= 39
// 		}
// 		if name%9 == 0 {
// 			name *= 13
// 			name += 1
// 			continue
// 		} else {
// 			name += 2
// 			name *= 3
// 		}

// 		if name%13 == 0 && name%9 == 0 {
// 			fmt.Println("service error")
// 			break
// 		} else {
// 			name++
// 		}
// 	}
// 	fmt.Printf("Итоговое число: %d\n", name)
// }
