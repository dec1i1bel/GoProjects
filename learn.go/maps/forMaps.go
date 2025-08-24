package main

import "fmt"

func main() {
	// инициализируем карту с ключами и значениями типа string
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "a value"
	// перебор пар ключ-значение
	for key, v := range aMap {
		fmt.Println("key:", key, "value:", v)
	}
	// перебор только значений
	for _, v := range aMap {
		fmt.Println("#", v)
	}
}
