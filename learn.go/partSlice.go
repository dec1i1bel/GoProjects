package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	I := len(aSlice)
	// 1-ые 5 эл-в - 2 вар-та получения
	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])
	// последние 2 эл-та - 2 вар-та получения
	fmt.Println(aSlice[I-2 : I])
	fmt.Println(aSlice[I-2:])
	// 1-ые 5 эл-в, помещаемые в срез ёмкостью 10-0=10
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))
	// эл-ты с индексами 2,3,4
	// ёмкость 10-2=8, длина 3
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))
	// эл-ты с индексами 0,1,2,3,4
	// ёмкость = 6-0=6, длина = 5, так как выбрал 1-ые 5 эл-в
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
