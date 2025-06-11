package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // Custom Type

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // 没有使用索引，使用下划线代替
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value) // append()返回一个新切片，然后赋值给原来的切片，相当于覆盖
		}
	}
	return cards // 返回 deck 类型的切片
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Receiver Function
func (d deck) print() { // 所有deck类型的变量都可以调用这个函数
	for i, card := range d { // 这里的d可以理解为”接受器参数“
		fmt.Println(i, card)
	}
}
