package main

import "fmt"

type Robot struct {
	serialNumber string
	factorName   string
	age          int
}

func printStudent(std Robot) {
	fmt.Println("========================")
	fmt.Println("Serial Number:", std.serialNumber)
	fmt.Println("Factor Name:", std.factorName)
	fmt.Println("Age:", std.age)

}

func main() {
	robot1 := Robot{
		serialNumber: "A1B2COD",
		factorName:   "ROBO2000",
		age:          100,
	}

	// fmt.Println("Robot1:", robot1)

	printStudent(robot1)

	// анонимная структура

	anonRobot := struct {
		serialNumber int
		age          int
		name         string
	}{
		serialNumber: 98541,
		age:          95,
		name:         "ROBBOB2000",
	}
	fmt.Println("AnonRobot:", anonRobot)

	// доступ к полям
	robot2 := Robot{"A1B2COD2", "ROBO2001", 101}
	fmt.Println("SerialNumber:", robot2.serialNumber)
	robot2.age += 2
	fmt.Println("newAge:", robot2.age)

	emptyRobot1 := Robot{}
	var emptyRobot2 Robot

	printStudent(emptyRobot1)
	printStudent(emptyRobot2)

	robotPointer := &Robot{
		serialNumber: "AB1C",
		factorName:   "Bob",
		age:          999,
	}
	fmt.Println("Value robotPointer", robotPointer)

	robotPointer2 := robotPointer
	(*robotPointer2).age += 3

	fmt.Println("Value after modify:", robotPointer2)

	fmt.Println("Age1:", (*robotPointer).age)
	fmt.Println("Age2:", robotPointer.age) //неявное разыменование
}
