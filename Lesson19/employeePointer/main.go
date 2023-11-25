package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

//работа с копией

func (e Employee) setName(newName string) {
	e.name = newName
}

//работа со значением

func (e *Employee) setSalary(newSalary int) {
	e.salary = newSalary
}

func main() {

	e := Employee{"Alex", 3000}
	fmt.Println("Before setting parameters:", e)

	e.setName("Bob")
	fmt.Println("After setName call", e)
	(&e).setSalary(4500)
	fmt.Println("After setSalary call:", e)
}
