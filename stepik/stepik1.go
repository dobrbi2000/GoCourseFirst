package main

import "fmt"

func main() {

	//Формат выходных данных
	//Выведите "YES", если все цифры числа различны, в противном случае - "NO".

	var num int //122
	fmt.Scan(&num)

	num1 := num % 10         //Переменная num1 содержит число 2
	num2 := (num % 100) / 10 //Переменная num2 содержит число 2
	num3 := num / 100        //Переменная num3 содержит число 1

	fmt.Println(num1, num2, num3)

	if num1 == num2 || num2 == num3 || num3 == num1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
