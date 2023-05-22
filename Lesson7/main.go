package main

import "fmt"

func main() {

	// switch

	// var price int
	// fmt.Scan(&price)

	// switch price {
	// case 100:
	// 	fmt.Println("First case")
	// case 110:
	// 	fmt.Println("Second case")
	// case 120:
	// 	fmt.Println("Third case")
	// case 130:
	// 	fmt.Println("Another case")
	// default:
	// 	fmt.Println("Default case")
	// }

	// var ageGroup string = "k"
	// switch ageGroup {
	// case "a", "b", "c":
	// 	fmt.Println("AgeGroup1")
	// case "d", "e":
	// 	fmt.Println("AgeGroup2")
	// default:
	// 	fmt.Println("AgeGroupNull")
	// }
	// var age int
	// fmt.Scan(&age)

	// switch {
	// case age <= 18:
	// 	fmt.Println("AgeGroup_18")
	// case age > 18 && age <= 30:
	// 	fmt.Println("AgeGroup_18-30")
	// case age > 30 && age <= 100:
	// 	fmt.Println("AgeGroup_30-100")
	// default:
	// 	fmt.Println("AgeGroupNull")

	// }

	var number int
	fmt.Scan(&number)
	switch {
	case number < 100:
		fmt.Printf("%d is less than 100\n", number)
		fallthrough
	case number < 200:
		fmt.Printf("%d is less than 200\n", number)
	}
}
