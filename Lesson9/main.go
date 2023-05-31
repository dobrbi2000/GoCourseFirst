package main

import (
	"fmt"
)

func main() {
	arr := [4]int{10, 20, 30, 40}
	var arrSlice []int = arr[0:2]
	fmt.Println("Slice[0:2]:", arrSlice)

	arr2Slice := []int{10, 20, 30, 40}
	fmt.Println("Slice_2:", arr2Slice)

	originArr := [...]int{30, 40, 50, 60, 70, 80}
	// firstSlace := originArr[1:4]

	// for i, _ := range firstSlace {
	// 	firstSlace[i]++
	// }
	// fmt.Println("Original Array:", originArr)
	// fmt.Println("First Slice:", firstSlace)

	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before mod Array:", originArr, "\nfSlice:", fSlice, "\nsSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After mod Array:", originArr, "\nfSlice:", fSlice, "\nsSlice:", sSlice)

	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "capacity:", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "capacity:", cap(wordsSlice))

	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	slace_1 := make([]int, 10, 15)

	// 1. инициализируем Arr = [15]int
	// 2. затем по нему делаем срез arr[0:10]

	fmt.Println(slace_1)

	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	myWords = append(myWords, "four")
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	fmt.Println("myWords:", myWords)

	mSlice := [][]int{
		{1, 2},
		{2, 3, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)

}
