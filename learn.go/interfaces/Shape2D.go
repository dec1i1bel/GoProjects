package main

import (
	"fmt"
	"math"
)

// определение интерфейса Shape2D, которое требует реализации метода типа Perimeter()
type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

// тип circle реализует интерфейс Shape2D путем реализации метода типа Perimeter()
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f\n", a.R, a.Perimeter())

	//  нотация interface{}(a).(Shape2D) проверяет‚ удовлетворяет ли переменная а интерфейсу Shape2D, не используя ее базовое значение (circle{R: 1.5}
	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D")
	}
}
