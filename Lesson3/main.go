package main

import "fmt"

func main() {

	var (
		age  int
		name string
	)
	fmt.Scan(&age, &name)

	fmt.Printf("Name is: %s\n age is: %d\n", name, age)

}
