package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	A := T{1, "F2", 3.0}
	fmt.Println("A:", A)
	// С помощью Elem() и указателя на переменную A она может быть изменена при необходимости.
	r := reflect.ValueOf(&A).Elem()
	fmt.Println("string value:", r.String())
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfa := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfa, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("changed!")
		}
		fmt.Println("A:", A)
	}
}
