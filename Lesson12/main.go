package main

import "fmt"

func add(a int, b int) int {
	result := a + b
	return result
}

func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

func rectangleParameters(lenght, width float64) (float64, float64) {
	perimetr := 2 * (lenght + width)
	area := lenght * width

	return perimetr, area
}

func namedReturn(a, b int) (perimetr, area int) {
	perimetr = 2 * (a + b)
	area = a * b

	return
}

func helloVariadic(a ...int) {
	// fmt.Printf("value %v and type %T\n", a, a)
	fmt.Println(a)
}

func someString(a, b int, words ...string) {
	fmt.Println("Parametr:", a)
	fmt.Println("Parametr:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func bigFunc(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)
}

func main() {
	fmt.Println("bigFunc:", bigFunc(10, 20))
	res := add(10, 20)
	fmt.Println("Result:", res)
	fmt.Println("Result of multi:", mult(1, 2, 3, 4))

	per, area := rectangleParameters(10.5, 10.5)
	newPer, _ := rectangleParameters(10, 10)
	secondPer, secondArea := namedReturn(6, 5)
	helloVariadic(10, 20, 30, 40, 50)
	helloVariadic(10, 20, 30, 40)
	helloVariadic()

	fmt.Println("Area:", area)
	fmt.Println("Perimetr:", per)
	fmt.Println("NewPerimetr:", newPer)
	fmt.Println("SecondPerimetr:", secondPer)
	fmt.Println("secondArea:", secondArea)

	someString(2, 10, "One", "Two")

	sum1 := sumVariadic(10, 20, 30, 40, 50)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)
}
