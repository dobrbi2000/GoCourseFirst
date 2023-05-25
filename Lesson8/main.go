package main

import "fmt"

func main() {

	// Массивы

	var arr [5]int
	fmt.Println(arr)

	arr[0] = 10
	arr[1] = 20
	arr[3] = 100
	arr[4] = -50
	fmt.Println(arr)

	newArr := [5]int{10, 20, 30, 40}
	fmt.Println(newArr)

	valueArr := [...]int{15, 25, 35, 45, 55}
	fmt.Println(valueArr, len(valueArr))

	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First array:", first)
	fmt.Println("Second array:", second)

	// var aArr [5]int
	// var bArr [6]int
	// aArr[0] = 100
	// bArr = aArr

	var sum float64
	floatArr := [...]float64{12.5, 13.5, 15.5, 16.5, 17.5}

	// for i := 0; i < len(floatArr); i++ {
	// 	fmt.Printf("%d element of array is %.2f\n", i, floatArr[i])
	// }

	for id, val := range floatArr {
		fmt.Printf("%d element of array is %.2f\n", id, val)
		sum += val
	}
	fmt.Println("Sum:", sum)
}
