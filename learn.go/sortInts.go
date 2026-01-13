package main

import (
	"fmt"
	"slices"
)

func sortIntsAsc(ints []int) []int {
	slices.Sort(ints)
	return ints
}

func sortIntsDesc(ints []int) []int {
	slices.SortFunc(ints, func(a, b int) int {
		return b - a
	})

	return ints
}

func main() {
	sortedAsc := sortIntsAsc([]int{5, 54, 65, 6, 454, 5, 3})
	fmt.Println("Ints sorted asc:", sortedAsc)

	sortedDesc := sortIntsAsc([]int{54, 6, 7, 554, 6, 5, 54, 5})
	fmt.Println("Ints sorted desc:", sortedDesc)
}
