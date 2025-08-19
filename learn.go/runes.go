package main

import (
	"fmt"
)

func main() {
	// Smile is a rune
	aString := "Hello World! 😀"
	fmt.Println("first character", string(aString[0]))

	// руны
	r := '€' // руна объявлена, если строка заключена в одинарные кавычки. руна - символ в этой строке

	// вывод рун в виде байтовых кодов
	fmt.Println("As an int32 value:", r)
	// преобразование руны в текст
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	// вывести существующую строку в виде рун
	for _, v := range aString {
		fmt.Printf("%x", v)
	}
	fmt.Println()

	// вывести существущую строку в виде символов
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
