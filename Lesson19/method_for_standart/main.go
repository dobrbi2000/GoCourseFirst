package main

import "fmt"

type MyInt int

func (mi *MyInt) IsEven() bool {
	if *mi%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MyInt(400)
	num2 := MyInt(401)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())

}
