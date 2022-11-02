package main

import (
	"fmt"
)

// Khai báo interface này có nghĩa là: bất kì kiểu dữ liệu nào, miễn là nó có method Speak() string thì đều được gọi là implement interface Animal
type Animal interface {
	Speak() string
}

type Dog struct {
	nickname string
	age      int
}

func (d Dog) Speak() string {
	return "Gâu gâu"
}

type Duck struct{}

func (d Duck) Speak() string {
	return "Cạp cạp"
}

func main() {
	animals := []Animal{Dog{"lulu", 2}, Duck{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	describe(100)
	describe("Hello Gopher")
	str := struct {
		name string
	}{
		name: "Alice",
	}
	describe(str)

	// Ứng dụng empty interface
	product := make(map[string]interface{})
	product["name"] = "Iphone 14 Promax"
	product["price"] = 32000000

	var i interface{} = "Gopher"
	fmt.Printf("type: %T, value: %v\n", i, i)

	s := i.(string)
	fmt.Printf("type: %T, value: %v\n", s, s)

	f, ok := i.(float64)
	fmt.Printf("type: %T, value: %v, ok: %t\n", f, f, ok)

}

// Empty interface: interface{}
func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}
