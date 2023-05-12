package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var value int
	// fmt.Scan(&value)

	// if value%2 == 0 {
	// 	fmt.Println("The number", value, "is even")
	// } else {
	// 	fmt.Println("The number", value, "is odd")
	// }

	// var color string
	// fmt.Scan(&color)
	// if strings.Compare(color, "green") == 0 {
	// 	fmt.Println("Color is green")
	// } else if strings.Compare(color, "red") == 0 {
	// 	fmt.Println("Color is red")
	// } else {
	// 	fmt.Println("Unknow color")
	// }

	var (
		A int
		B int
		C int
	)
	fmt.Scan(&A, &B, &C)

	fmt.Print(strconv.Itoa(C) + strconv.Itoa(A) + strconv.Itoa(B))

}
