package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {

	aBoolean, bBoolean := true, false

	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	var a int = 32

	fmt.Printf("Type %T\n", a)
	fmt.Printf("Type %T %d bytes\n", a, unsafe.Sizeof(a))

	name := "Евгений"

	fmt.Println("Length", name, len(name))
	fmt.Println("Amount:", name, utf8.RuneCountInString(name))

	totalString, subString := "ABCDEFGHL", "ABC"
	fmt.Println(strings.Contains(totalString, subString))

}
