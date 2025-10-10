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
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	// возвращается значение reflect.Value переменной A:
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())
	// получаем тип данных переменной — в данном случае переменной A
	iType := r.Type()
	fmt.Printf("i Type %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	//  посетить все поля структуры и изучить их характеристики
	for i := 0; i < r.NumField(); i++ {
		// вывод имени, типа данных и значения полей:
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		// есть ли другие структуры, встроенные в запись?
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// нужно преобразование в строку, чтобы можно было сравнить
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// то же, что и раньше, но с использованием внутреннего значения
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
