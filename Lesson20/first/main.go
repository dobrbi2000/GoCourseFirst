package main

import (
	"fmt"
	"unicode/utf8"
)

type BigWord interface {
	IsBig() bool
}

type MyString string

func (ms MyString) IsBig() bool {
	if utf8.RuneCountInString(string(ms)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MyString("StringStringStringStringStringString")
	var InterfaceSample BigWord
	InterfaceSample = sample

	fmt.Println("IsBig?", InterfaceSample.IsBig())

}
