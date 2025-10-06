package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}
type Record struct {
	Field1 string
	Field2 float64
	fields Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	// возвращается значение reflect.Value переменной A:
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())
	// Используя Type(), мы получаем тип данных переменной — в данном случае переменной A
	iType := r.Type()

	//  посетить все поля структуры и изучить их характеристики
	for i := 0; i < r.NumField(); i++ {
		// вывод имени, типа данных и значения полей:
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s", r.Field(i).Type())
		fmt.Printf("\tand value _%v_%n", r.Field(i).Interface())
	}
}
