package main

import "fmt"

func main() {
	// создаём карту с ключами типа string и значениями int
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("1_aMap:", aMap)

	aMap = nil
	fmt.Println("2_aMap:", aMap)

	// Перед использованием карты полезно проверять, не указывает ли она на nil. Оператор if aMap == nil позволяет нам определить, можем ли мы сохранить значение ключа/пары в aMap. В данном случае не можем, и попытка это сделать приведет к сбою в программе. Мы исправляем это с помощью оператора aMap = map[string]int{}.
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}

	aMap["test"] = 1
	aMap = nil
	// аварийное завершение программы
	aMap["test"] = 1
}
