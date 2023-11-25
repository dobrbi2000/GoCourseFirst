package main

import "fmt"

type Rectangle struct {
	lenght, width int
}

func New(newLenght, newWidth int) *Rectangle {
	return &Rectangle{newLenght, newWidth}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

func (r *Rectangle) Area() int {
	return r.lenght * r.width
}

func main() {
	r1 := New(10, 20)
	fmt.Printf("Type as %T and value %v\n", r1, r1)
	fmt.Println("Perimeter:", r1.Perimeter())
	fmt.Println("Area", r1.Area())

}
