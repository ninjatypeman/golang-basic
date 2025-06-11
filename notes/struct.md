# Go语言结构体(Struct)介绍

结构体(struct)是Go语言中用于定义和存储复杂数据类型的基本构建块，它允许你将不同类型的数据组合在一起形成一个新的复合类型。

## 基本定义

```go
// 基本结构体定义
type Person struct {
    Name string
    Age  int
}
```

## 结构体特点

1. **字段集合**：由一组不同类型的字段(field)组成，每个字段都有名称和类型
2. **值类型**：结构体是值类型，不是引用类型
3. **内存布局**：字段在内存中是连续存储的

## 声明和使用

```go
// 声明方式1: 直接声明
var p1 Person
p1.Name = "Alice"
p1.Age = 25

// 声明方式2: 使用结构体字面量
p2 := Person{"Bob", 30}

// 声明方式3: 指定字段名赋值
p3 := Person{
    Name: "Charlie",
    Age: 35,
}
```

## 匿名结构体

不需要预先定义类型，直接使用的结构体：

```go
anonymous := struct {
    Title string
    Pages int
}{
    Title: "Go Programming",
    Pages: 350,
}
```

## 结构体嵌套

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Age     int
    Address Address // 嵌套结构体
}
```

## 匿名字段(嵌入结构体)

```go
type Employee struct {
    Name string
    Address // 匿名字段(嵌入)
}

// 可以直接访问Address的字段
e := Employee{Name: "Dave", Address: Address{"NY", "US"}}
fmt.Println(e.City) // 直接访问
```

## 结构体方法

可以为结构体定义方法：

```go
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s\n", p.Name)
}

// 使用
p := Person{"Eve", 28}
p.Greet()
```

## 指针接收器

```go
func (p *Person) Birthday() {
    p.Age++ // 修改原结构体的值
}
```

## 结构体标签(Struct Tags)

```go
type User struct {
    ID   int    `json:"id" db:"user_id"`
    Name string `json:"name" db:"user_name"`
}
```

结构体标签常用于序列化、ORM映射等场景。

Go语言的结构体提供了组织数据的基础设施，是面向对象编程的重要部分，虽然Go没有类的概念，但通过结构体和方法的组合可以实现类似的效果。


```go
type student struct{
	name string
	age int
	number string
	height float64
	weight float64
}
```