package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // Custom Type

// Receiver Function
func (d deck) print() { // 所有deck类型的变量都可以调用这个函数
	for i, card := range d { // 这里的d可以理解为”接受器参数“
		fmt.Println(i, card)
	}
}
