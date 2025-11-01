package main

import "fmt"

// Функция, возвращающая 2 значения
func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// Сортировка от меньшего к большему
func sortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func main() {
	n := 10
	d, s := doubleSquare(n)
	fmt.Println("Double of", n, "is", d)
	fmt.Println("Square of", n, "is", s)

	// анонимная функция

}
