package main

import "fmt"

func main() {
	var (
		name     string
		lastName string
		age      int
	)

	fmt.Scan(&name, &lastName, &age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", name, lastName, age)

}
