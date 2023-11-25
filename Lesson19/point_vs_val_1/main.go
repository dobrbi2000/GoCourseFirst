package main

import "fmt"

type Rectangle struct {
	lenght, width int
}

func (r *Rectangle) Area() int {
	return r.lenght * r.width
}

func Area(r *Rectangle) int {
	return r.lenght * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	rectangleAsValue := Rectangle{10, 10}

	rectangleAsPointer := &rectangleAsValue

	fmt.Println("Call Area function for &rectangle", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &rectangle", rectangleAsPointer.Area())

	//вызываем исходное значение
	fmt.Println("Call Area method for rectangle", rectangleAsValue.Area())
	//Area(rectangleAsValue)

	fmt.Println("Before changing width:", rectangleAsValue)
	rectangleAsPointer.SetWidth(777)
	fmt.Println("After change with method SetWidth for &rectangle", rectangleAsValue)

	rectangleAsValue.SetWidth(555) //вызов метода
	fmt.Println("After change with method SetWidth for &rectangle", rectangleAsValue)
}
