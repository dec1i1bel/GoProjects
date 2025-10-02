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
	// toDo: дальше по задаче
}

func main() {
	arrToMap()
}
