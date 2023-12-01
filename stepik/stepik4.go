package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	res := 0

	for i := a; i <= b; i++ {
		res += i

	}
	fmt.Println(res)
}
