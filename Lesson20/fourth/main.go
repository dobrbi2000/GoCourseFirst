package main

import "fmt"

func DoSomething(pretendent interface{}) {
	switch pretendent.(type) { //пытаемся извлечь нижележащий тип
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hi")
	DoSomething(func(a, b int) int { return a + b })

}
