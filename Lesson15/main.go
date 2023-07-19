package main

import "fmt"

func main() {

	// указатели
	var variable int = 100
	var pointer *int = &variable

	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	var zeroPointer *int
	fmt.Printf("Type of %T and value %v\n", zeroPointer, zeroPointer)

	var numericValue int = 32
	var poinerToNumeric *int = &numericValue

	fmt.Printf("Value in numericValue is %v\nAdress is %v\n", *poinerToNumeric, poinerToNumeric)

	// var zeroVar int
	// var zeroPoint *int = &zeroVar

	zeroPoint := new(int)

	fmt.Printf("Value in *zeroPoint %v\nAdress is %v\n", *zeroPoint, zeroPoint)

	zeroPointerToInt := new(int)
	fmt.Printf("Adress is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Adress is %v and New Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)

	sample := 1
	samplePointer := &sample
	fmt.Println("Origin value of sample", sample)
	changeParam(samplePointer)
	fmt.Println("After changing value of sample", sample)

	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and adress %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr1: %T and adress %v and value %v\n", ptr2, ptr2, *ptr2)
}

func returnPointer() *int {
	var numeric int = 123
	return &numeric
}

func changeParam(val *int) {
	*val += 100
}
