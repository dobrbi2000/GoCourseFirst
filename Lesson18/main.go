package main

import "fmt"

type Factory struct {
	yearProduce int
	infoShort   string
	infoLong    string
}

type Robot struct {
	name    string
	serial  string
	factory Factory
}

type Robotron struct {
	name   string
	serial string
	age    int
	Factory
}

func main() {

	robot := Robot{
		name:   "Max",
		serial: "2000",
		factory: Factory{
			yearProduce: 2084,
			infoShort:   "Cat&Man",
			infoLong:    "Cat&Man Inc.",
		},
	}

	fmt.Println("Name:", robot.name)
	fmt.Println("Serial№:", robot.serial)
	fmt.Println("Year of creation:", robot.factory.yearProduce)
	fmt.Println("Information:", robot.factory.infoShort)

	robotron1 := Robotron{
		name:   "Petrovich",
		serial: "8500",
		age:    601,
		Factory: Factory{
			yearProduce: 2800,
			infoShort:   "Short info",
		},
	}

	fmt.Println("Name:", robotron1.name)
	fmt.Println("Year:", robotron1.yearProduce)
	fmt.Println("Short Info:", robotron1.infoShort)

}
