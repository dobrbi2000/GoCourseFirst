package main

import "fmt"

type Empty interface {
}

type Student struct {
	name string
}

func Describer(pretendent interface{}) {
	fmt.Printf("Type - %T and value %v\n", pretendent, pretendent)
}

func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int)
	fmt.Println("Value", val, "OK?", ok)
}

func main() {
	strSample := "Hello World!"
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Jorh"})

	SemiGeneric(10)
	SemiGeneric(Student{"Vasya"})
	SemiGeneric("Hello world")

}
