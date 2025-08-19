package main

import "fmt"

func main() {
	// создаём пустой срез
	aSlice := []float64{}
	// его длина и ёмкость равны 0
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, 34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// срез длинной 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// добавляем элемент сверх заявленной длины, используя append
	t = append(t, -5)
	fmt.Println(t)

	// двумерный срез: создание и одновременно инициализация
	// измерений может быть больше
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	// получение всех элементов:
	// с помощью вложенного for
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, "")
		}
		fmt.Println()
	}

	// создание без инициализации с помощью make()
	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
