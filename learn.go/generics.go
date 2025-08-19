package main

import "fmt"

// [T any] - дженерик
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, "")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"one", "two", "three"}
	Print(Ints)
	Print(Strings)
}
