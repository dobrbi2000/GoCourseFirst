package main

import "fmt"

type Rectangle struct {
	lenght, width int
}

//вызывается только у копии
func Perimeter(r Rectangle) int {
	return 2 * (r.lenght + r.width)
}

//вызывается у копии+ссылка
func (r Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

func (r Rectangle) SetLength(newLenght int) {
	r.lenght = newLenght
}

func main() {

	rectangleValue := Rectangle{12, 12}

	fmt.Println("Call function for rectangleValue:", Perimeter(rectangleValue))
	fmt.Println("Call method for rectangleValue:", rectangleValue.Perimeter())

	rectangleValuePointer := &rectangleValue

	fmt.Println("Call method for rectangleValue:", rectangleValuePointer.Perimeter())
	//Perimeter(rectangleValuePointer) // к 10 строке

	fmt.Println("Before call method SetLength:", rectangleValue)
	rectangleValue.SetLength(9999)
	fmt.Println("After call method SetLength:", rectangleValue)

	rectangleValuePointer.SetLength(9999999999)
	fmt.Println("After call method SetLenght for &rectangleValue", rectangleValue)

}
