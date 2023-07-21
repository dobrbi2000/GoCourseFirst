package main

import "fmt"

func mutationSlice(sls []int) {
	sls[1] = 100
	sls[2] = 900
}

func main() {

	newArr := [3]int{1, 2, 3}
	fmt.Println("newArr before mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutationSlice:", newArr)

}
