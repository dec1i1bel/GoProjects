package main

import "fmt"

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	// если у return нет аргументов, функция автоматически возвращает текущее значение каждого именованного возвращаемого значения в том порядке, в котором они были объявлены в сигнатуре. здесь будет return min,max
	return
}

func main() {
	min, max := minMax(5, 2)
	fmt.Println("Min:", min, "Max:", max)
}
