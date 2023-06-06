package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	word := "Тестовое слово"
	fmt.Printf("String %s\n", word)
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i])
	}
	fmt.Println()

	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}

	fmt.Println()

	fmt.Printf("Runes: ")
	sliceRune := []rune(word)
	for i := 0; i < len(sliceRune); i++ {
		fmt.Printf("%c ", sliceRune[i])
	}

	fmt.Println()

	for id, runeVal := range word {
		fmt.Printf("%c starts at position %d\n", runeVal, id)
	}

	myBiteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myBiteSlice)
	fmt.Println(myStr)

	fmt.Println("Lenght of Вася:", len("Вася"))
	fmt.Println("Lenght of Вася:", utf8.RuneCountInString("Вася"))

	word1, word2 := "Вася", "Петя"

	if word1 != word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	word3 := word1 + word2
	fmt.Println(word3)

	firstName := "Alex"
	secondName := "Petrov"
	fullName := fmt.Sprintf("%s %s", firstName, secondName)
	fmt.Println("FullName:", fullName)

	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'T'
	fullName = string(fullNameSlice)
	fmt.Println("New string:", fullName)

	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d is %T", numInt, numInt)

}
