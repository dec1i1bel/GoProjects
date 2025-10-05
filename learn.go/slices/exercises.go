package main

import (
	"fmt"
	"os"
)

// создаём срез из двух массивов
func arraysToSlice(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	res = append(res, arr1...)
	res = append(res, arr2...)

	return res
}

// создаём массив из двух массивов
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

// создаём массив из двух срезов
func slicesToArray(sl1 []int, sl2 []int) []int {
	var res []int

	for _, i := range sl1 {
		res = append(res, i)
	}

	for _, j := range sl2 {
		res = append(res, j)
	}

	return res
}

func osArgsToSlice() {
	args := os.Args
	res := make(map[int]string)
	for k, v := range args {
		res[k] = v
	}
	fmt.Println("os_args: ", res)
}

func main() {
	osArgsToSlice()
	arr1 := []int{46, 34, 6, 4, 32}
	arr2 := []int{635534, 5423}

	sl1 := arraysToSlice(arr1, arr2)
	fmt.Println("sl1:", sl1)
	fmt.Println("sl1_len:", len(sl1), "cap:", cap(sl1))

	arr3 := arraysToArray(arr1, arr2)
	fmt.Println("arr3:", arr3)

	sl2 := []int{34, 65, 45, 8, 5}
	fmt.Println("sl2:", sl2)
	arr4 := slicesToArray(sl1, sl2)
	fmt.Println("arr4:", arr4)
}
