package main

import "fmt"

func arraysToSlice(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	res = append(res, arr1...)
	res = append(res, arr2...)

	return res
}

func arraysToArray(arr1 []int, arr2 []int) []int {
	res := []int{}

	for _, v := range arr1 {
		res = append(res, v)
	}

	for _, v := range arr2 {
		res = append(res, v)
	}

	return res
}

func main() {
	arr1 := []int{46, 34, 6, 4, 32}
	arr2 := []int{635534, 5423}
	sl1 := arraysToSlice(arr1, arr2)
	fmt.Println("1:", sl1)
	fmt.Println("1_len:", len(sl1), "cap:", cap(sl1))
	sl2 := arraysToArray(arr1, arr2)
	fmt.Println("2:", sl2)
}
