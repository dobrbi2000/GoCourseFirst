package main

import "fmt"

func main() {

	mapper := make(map[string]int)
	fmt.Println("Map:", mapper)

	mapper["Kate"] = 27
	mapper["Petr"] = 29
	fmt.Println("Map after adding:", mapper)

	mapper1 := map[string]int{
		"Katarina": 33,
		"Jack":     42,
	}
	mapper1["John"] = 29
	fmt.Println("Mapper1:", mapper1)

	Person := "Katarina"
	fmt.Println("Age of Person:", mapper1[Person])

	Person = "Kate"
	fmt.Println("Age of Person:", mapper1[Person])

	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	if value, ok := employee["Jack"]; ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Value doesn't exists")
	}

	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}

	fmt.Println("Before delete:", employee)
	delete(employee, "Den")
	fmt.Println("After delete", employee)

	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("Words map:", words)
	fmt.Println("newWords map:", newWords)

	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

}
