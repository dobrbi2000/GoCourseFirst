package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

//метод

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Positi on:", e.position)
	fmt.Printf("Salary : %d %s\n", e.salary, e.currency)

}

func main() {

	emp1 := Employee{
		name:     "Bob",
		position: "Golang developer",
		salary:   3000,
		currency: "EUR",
	}

	emp1.DisplayInfo()

}
