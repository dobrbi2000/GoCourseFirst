package main

import "fmt"

func main() {

	fruit := make([]string, 6)

	fruit[0] = "banana"
	fruit[1] = "cucumber" //"apple"
	fruit[2] = "grape"
	fruit[3] = "watermelon"
	fruit[4] = "pineapple"
	fruit[5] = ""

	fmt.Println("Before append in fruit", fruit)

	xs := fruit[1:3:3] //"cucumber" "grape"
	xs[0] = "apple"    // "apple" "grape" "mango"

	xs = append(xs, "mango")

	fmt.Println("After append in fruit", fruit)
	fmt.Println(xs)
}
