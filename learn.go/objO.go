// возможности ООП
package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

// реализация IntА. Структура c удовлетворяет IntA, поскольку реализует foo()
func (VarC c) foo() {
	fmt.Printf("Foo Processing", VarC)
}

// // реализация IntB B. Структура c удовлетворяет IntB. Поскольку структура c удовлетворяет как IntA, так и IntB, она неявно удовлетворяет IntC, который представляет собой композицию интерфейсов IntA и IntB
func (VarC c) bar() {
	fmt.Println("Bar processing", VarC)
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

// структура compose получает поля анонимной структуры a, то есть использует её
type compose struct {
	field1 int
	a
}

// разные структуры могут иметь одноимённые методы
func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	// c состоит из структуры а и структуры b
	var iC c = c{a{120, 12}, b{"-12", -12}}
	iC.A.A()
	iC.B.A()

	// не будет работать:
	// iComp := compose{field1: 123, a{456, 789}}
	// iComp := compose{field1: 123, XX: 456, YY: 789}
	// а так будет:
	iComp := compose{123, a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	iC.bar()
	processA(iC)
}
