package main

import (
	"fmt"
	"slices"
)

// Возврат функцией именованных возвращаемых значений - отсортированных целых чисел

func namedIntsSortedAsc(intsAsc []int, intsDesc []int) (intsSortedAsc []int, intsSortedDesc []int) {
	slices.Sort(intsAsc)
	intsSortedAsc = intsAsc

	slices.SortFunc(intsDesc, func(a, b int) int {
		return b - a
	})

	intsSortedDesc = intsDesc

	return
}

func main() {
	a, b := namedIntsSortedAsc([]int{4, 5, 65, 4, 3, 565, 56, 5, 4}, []int{6464, 4543, 543, 456, 231})

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
