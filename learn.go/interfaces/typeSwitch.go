// переключение поведения в зависимости от типа входных данных
package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("entry type")
	default:
		fmt.Println("Not supported type: %T\n", T)
	}
}
