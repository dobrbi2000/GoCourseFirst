package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	lenght, width int
}

// функция
func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

// метод
func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.lenght + r.width)
}

func main() {

	c := Circle{11.5}

	fmt.Println("Call function:", Perimeter(c))
	fmt.Println("Call method:", c.Perimeter())

	r := Rectangle{10, 20}

	fmt.Println("Call function for rectangle:", PerimeterRectangle(r))
	fmt.Println("Call method for rectangle:", r.Perimeter())

}
