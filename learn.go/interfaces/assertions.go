package main

import "fmt"

// возвращает значение int, обёрнутое` в пустой интерфейс
func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber() // получаем пустой интерфейс со значением

	number := anInt.(int) // извлекаем значение типа int из переменной пустого интерфейса anInt
	number++
	fmt.Println(number)

	// завершится с ошибкой, но это контролируемый сбой
	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful:", value)
	} else {
		fmt.Println("Type assertion failed")
	}

	// нормальное выполнение приложения не гарантируется
	i := anInt.(int)
	fmt.Println("i:", i)

	// вызовет panic
	_ = anInt.(bool)
}
