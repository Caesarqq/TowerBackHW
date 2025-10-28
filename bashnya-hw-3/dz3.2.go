package main

import "fmt"

type Deque struct {
	items []int
}

func (d *Deque) PushFront(value int) {
	d.items = append([]int{value}, d.items...)
}

func (d *Deque) PushBack(value int) {
	d.items = append(d.items, value)
}

func (d *Deque) PopFront() int {
	if d.IsEmpty() {
		fmt.Println("Дек пустой")
		return -1
	}
	value := d.items[0]
	d.items = d.items[1:]
	return value
}

func (d *Deque) PopBack() int {
	if d.IsEmpty() {
		fmt.Println("Дек пустой")
		return -1
	}
	lastIndex := len(d.items) - 1
	value := d.items[lastIndex]
	d.items = d.items[:lastIndex]
	return value
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}

func (d *Deque) Clear() {
	d.items = []int{}
}

func main() {
	deque := Deque{}
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)
	fmt.Println("Размер дека:", deque.Size())
	fmt.Println("PopFront:", deque.PopFront())
	fmt.Println("PopBack:", deque.PopBack())
	fmt.Println("PopBack:", deque.PopBack())
	fmt.Println("Проверка пустоты", deque.IsEmpty())
}
