package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926
const (
	C1 = "c1c1c1"
	C2 = "c2c2c2"
	C3 = "c3c3c3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)
	const (
		// iota - генератор контант
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	/*
		объявление констант через iota выше эквивалентно:
		const (
			Zero = 0
			One = 1
			Two = 2
			Three = 3
			Four = 4
		)
	*/

	fmt.Println(One)
	fmt.Println(Two)
	// вариант использования iota. нижние слеши - для пропуска нежелательных значений
	const (
		p2_0 Power2 = 1 << iota // синатксис << - битовый сдвиг - используется в том чсиле для возведения в степень двойки. например, 1 << n == 2 в степени n
		_                       // пропускаем генерацию константы p2_1 , которая была бы равна 2 в степени 1
		p2_2
		_ // пропускаем p2_3 == 2 в степени 3
		p2_4
		_ // пропускаем p2_5
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)

	// toDo: узнать, что за тип Power2 и что делает синтаксис <<
}
