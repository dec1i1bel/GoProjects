package main

import "fmt"

func main() {
	// определена только длина среза, тогда capacity == length
	a := make([]int, 4)
	fmt.Println("a: length:", len(a), "capacity:", cap(a))
	// инициализация среза. capacity == length
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("b: length:", len(b), "capacity:", cap(b))
	// capacity == length, не потому что Go так придумал, а мы так задали
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	// добавляем элемет сверх существуюещй ёмкости. тогда она удваивается до 8
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	fmt.Println("aSlice: length", len(aSlice), "capacity:", cap(aSlice))
	// добавляем 4 элемента. операция <...> разворачивает []int{-1, -2, -3, -4} в несколько аргументов, и append() по отдельности добавляет каждый аргумент в aSlice.
	// ёмкость снова удваивается
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)
	fmt.Println("aSlice: length:", len(aSlice), "capacity:", cap(aSlice))
}
