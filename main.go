package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} // 不指定长度的切片
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
