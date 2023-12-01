package main

import "fmt"

func main() {

	var a, b, sum int
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		if b <= 10 && b%8 == 0 && b <= 99 {
			sum += b
		}
	}
	fmt.Println(sum)

}
