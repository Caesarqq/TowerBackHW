// package main

// import (
// 	"fmt"
// )

// // Функция для преобразования числа в текстовое представление
// func numberToWords(n int) string {
// 	if n == 0 {
// 		return "ноль"
// 	}

// 	// Словари для чисел
// 	units := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
// 	teens := []string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
// 	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
// 	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
// 	thousands := []string{"", "тысяча", "тысячи", "тысяч"}
// 	millions := []string{"", "миллион", "миллиона", "миллионов"}

// 	var result string

// 	// Обработка миллионов
// 	if n >= 1000000 {
// 		mil := n / 1000000
// 		if mil == 1 {
// 			result += "один миллион "
// 		} else if mil == 2 {
// 			result += "два миллиона "
// 		} else {
// 			result += numberToWordsHundreds(mil, units, teens, tens, hundreds) + " "
// 			if mil%10 == 1 && mil%100 != 11 {
// 				result += "миллион "
// 			} else if mil%10 >= 2 && mil%10 <= 4 && (mil%100 < 10 || mil%100 > 20) {
// 				result += "миллиона "
// 			} else {
// 				result += "миллионов "
// 			}
// 		}
// 		n %= 1000000
// 	}

// 	// Обработка тысяч
// 	if n >= 1000 {
// 		thou := n / 1000
// 		if thou == 1 {
// 			result += "одна тысяча "
// 		} else if thou == 2 {
// 			result += "две тысячи "
// 		} else {
// 			result += numberToWordsHundreds(thou, units, teens, tens, hundreds) + " "
// 			if thou%10 == 1 && thou%100 != 11 {
// 				result += "тысяча "
// 			} else if thou%10 >= 2 && thou%10 <= 4 && (thou%100 < 10 || thou%100 > 20) {
// 				result += "тысячи "
// 			} else {
// 				result += "тысяч "
// 			}
// 		}
// 		n %= 1000
// 	}

// 	// Обработка сотен, десятков и единиц
// 	result += numberToWordsHundreds(n, units, teens, tens, hundreds)

// 	return result
// }

// // Вспомогательная функция для обработки чисел до 999
// func numberToWordsHundreds(n int, units, teens, tens, hundreds []string) string {
// 	var result string

// 	if n >= 100 {
// 		result += hundreds[n/100] + " "
// 		n %= 100
// 	}

// 	if n >= 20 {
// 		result += tens[n/10] + " "
// 		n %= 10
// 		if n > 0 {
// 			result += units[n] + " "
// 		}
// 	} else if n >= 10 {
// 		result += teens[n-10] + " "
// 	} else if n > 0 {
// 		result += units[n] + " "
// 	}

// 	return result
// }

// func main() {
// 	var name int
// 	fmt.Scan(&name)
// 	if name > 12306 {
// 		fmt.Println("Число не соответствует изначальному условию, попробуйте ещё раз")
// 		return
// 	}
// 	for name != 6239143 { // Изменено условие для достижения 6239143
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
// 	// Вывод числа и его текстового представления
// 	fmt.Printf("Итоговое число: %d\n", name)
// 	fmt.Printf("Прописью: %s\n", numberToWords(name))
// }
