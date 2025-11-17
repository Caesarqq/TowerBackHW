package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Стек пустой")
		return -1
	}
	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []int{}
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Размер стека:", stack.Size())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Проверка пустоты", stack.IsEmpty())
	stack.Clear()
	fmt.Println("Проверка пустоты", stack.IsEmpty())
}
