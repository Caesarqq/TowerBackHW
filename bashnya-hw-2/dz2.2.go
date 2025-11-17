// package main

// import (
// 	"fmt"
// 	"strings"
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
// 			return
// 		} else {
// 			name++
// 		}
// 	}
// 	fmt.Printf("Итоговое число: %d\n", name)
// 	fmt.Printf("Текстом: %s\n", numberToWords(name))
// }

// func numberToWords(n int) string {

// 	ones := map[int]string{
// 		0: "", 1: "один", 2: "два", 3: "три", 4: "четыре",
// 		5: "пять", 6: "шесть", 7: "семь", 8: "восемь", 9: "девять",
// 	}

// 	teens := map[int]string{
// 		10: "десять", 11: "одиннадцать", 12: "двенадцать", 13: "тринадцать",
// 		14: "четырнадцать", 15: "пятнадцать", 16: "шестнадцать",
// 		17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать",
// 	}

// 	tens := map[int]string{
// 		2: "двадцать", 3: "тридцать", 4: "сорок", 5: "пятьдесят",
// 		6: "шестьдесят", 7: "семьдесят", 8: "восемьдесят", 9: "девяносто",
// 	}

// 	hundreds := map[int]string{
// 		1: "сто", 2: "двести", 3: "триста", 4: "четыреста",
// 		5: "пятьсот", 6: "шестьсот", 7: "семьсот", 8: "восемьсот", 9: "девятьсот",
// 	}

// 	if n == 0 {
// 		return "ноль"
// 	}

// 	var parts []string

// 	thousands := n / 1000
// 	if thousands > 0 {
// 		if thousands >= 10 && thousands < 20 {
// 			parts = append(parts, teens[thousands])
// 		} else {
// 			if thousands >= 20 {
// 				parts = append(parts, tens[thousands/10])
// 			}
// 			lastDigit := thousands % 10
// 			if lastDigit == 1 {
// 				parts = append(parts, "одна")
// 			} else if lastDigit == 2 {
// 				parts = append(parts, "две")
// 			} else if lastDigit > 0 {
// 				parts = append(parts, ones[lastDigit])
// 			}
// 		}

// 		lastDigit := thousands % 10
// 		lastTwo := thousands % 100
// 		if lastTwo >= 11 && lastTwo <= 14 {
// 			parts = append(parts, "тысяч")
// 		} else if lastDigit == 1 {
// 			parts = append(parts, "тысяча")
// 		} else if lastDigit >= 2 && lastDigit <= 4 {
// 			parts = append(parts, "тысячи")
// 		} else {
// 			parts = append(parts, "тысяч")
// 		}
// 	}

// 	n = n % 1000
// 	h := n / 100
// 	if h > 0 {
// 		parts = append(parts, hundreds[h])
// 	}

// 	n = n % 100
// 	if n >= 10 && n < 20 {
// 		parts = append(parts, teens[n])
// 	} else {
// 		if n >= 20 {
// 			parts = append(parts, tens[n/10])
// 		}
// 		if n%10 > 0 {
// 			parts = append(parts, ones[n%10])
// 		}
// 	}

// 	result := strings.Join(parts, " ")

// 	if len(result) > 0 {
// 		result = strings.ToUpper(string(result[0])) + result[1:]
// 	}

// 	return result
// }
