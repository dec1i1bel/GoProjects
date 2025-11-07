// функции с переменным количеством аргументов

package main

import (
	"fmt"
	"os"
)

// переменное кол-во аргументов типа float64
func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}

func main() {
	sum := addFloats("Addingvnumbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)
	// работает, поскольку содержимое s не распаковано
	everything(s)

	// не получится напрямую передать []string как []interface{}
	// Сначала необходимо преобразовать!
	empty := make([]interface{}, len(os.Args[1:]))

	// преобразовать []string в[]interface{}, чтобы использовать операцию распаковки
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	// распаковываем содержимое empty
	everything(empty...)

	// другой способ преобразования []string в []interface{}
	arguments := os.Args[1:]
	empty = make([]interface{}, len(arguments))
	for i := range arguments {
		empty[i] = arguments[i]
	}
	everything(empty...)

	// это тоже работает, поскольку вы трижды передаете всю переменную str, а не ее содержимое. Итак, срез содержит три элемента‚ и каждый из них равен содержимому переменной str
	str := []string{"one", "two", "three"}
	everything(str, str, str)
}
