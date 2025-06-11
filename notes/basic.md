# Go 基础
### 安装 Go

Fedora 安装 Go：
```bash
sudo dnf install golang
go version
mkdir -p ~/code/go-proj
cd ~/code/go-proj
go mod init go-proj
go run main.go deck.go
```
示例代码:
```go
package main // Package declaration
import "fmt" // format, import a package
// Declare function
func main() {
	// var whatToSay string
	// whatToSay = "Hello World!"
	whatToSay := "Hello World!"
	saySomething(whatToSay)
}
func saySomething(whatToSay string) {
	fmt.Println(whatToSay)
}
```

### 变量

Types: bool, string, int, float64...  
静态（Static）类型：C, Go, Java  
动态（Dynamic）类型：JS, Python, Ruby

```go
var card string = "Ace of Spades"
var card = "Ace of Spades"
card := "Ace of Spades"
card = "Five of Diamonds"  // 重新赋值
```

### 数组和切片

数组和切片都是用于存储相同类型数据元素的集合（数据结构），每个元素必须是相同类型

- Array(Fixed length)
- Slice(An Array that can grow or shrink)

```go
func main() {
	cards := []string{"Ace of Diamonds", newCard()} // 不指定长度的切片
	cards = append(cards, "Six of Spades") // 追加元素、生成新的切片
	for i, card := range cards { //遍历切片
		fmt.Println(i, card)
	}
}
func newCard() string { // 函数返回一个字符串
	return "Five of diamonds"
}

```

### 自定义类型
自定义类型（Custom Type Declarations）：扩展基本类型，添加更多功能
```go
package main
// Create a new type of 'deck'
// which is a slice of strings
// 定义一个自定义类型，它的地层类型是 []string （字符串切片）、继承 []string 的所有特性
type deck []string
```

### 接收器（接收者）函数
接收器函数（Receiver Functions），格式：

```go
func (receiver ReceiverType) MethodName(parameters) (returnTypes) {
    // 函数体
}
```

示例代码：
```go
// deck.go
type deck []string
func (d deck) print() { // 允许 deck 类型的实例调用函数
	for i, card := range d { // 这里的 d 可以理解为”接收器参数“
		fmt.Println(i, card)
	}
}
//main.go
cards.print()
```

### 迭代
双重 for 循环生成新的切片
```go
//deck.go
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits { // 忽略用不到的索引/键
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards // 返回 deck 类型的切片
}

// main.go
func main() {
	cards := newDeck()
	cards.print()
}
```
多个形参、多个返回值
```go
// deck.go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//main.go
func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
```

Shuffle