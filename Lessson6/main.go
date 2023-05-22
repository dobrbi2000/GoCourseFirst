package main

import (
	"fmt"
)

func main() {
	//break
	for i := 0; i <= 15; i++ {
		if i > 10 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Printf("After for loop with BREAK\n")

	// continue
	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 14 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Printf("After for loop with CONTINUE\n")

	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
