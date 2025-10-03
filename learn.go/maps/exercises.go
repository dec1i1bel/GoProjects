package main

import "fmt"

func arrToMap() {
	arr := [3]string{"one", "two", "three"}
	result := make(map[int]string)

	for k, v := range arr {
		result[k] = v
	}

	fmt.Println(result)
}

func mapTo2Slices() {
	m := make(map[string]string)
	m["k_1"] = "val_one"
	m["k_2"] = "val_two"
	m["k_3"] = "val_three"

	// toDo
	skeys := make([]string, 0)
	svals := make([]string, 0)

	for k, v := range m {
		skeys = append(skeys, k)
		svals = append(svals, v)
	}
}

func main() {
	arrToMap()
	mapTo2Slices()
}
