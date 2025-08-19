package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 567
	input := strconv.Itoa(n)
	fmt.Printf("n = %v Itoa() converted: %v\n", n, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("n = %v FormatInt() converted: %v\n", n, input)
	input = string(n) // получаем ASCII-представление n. если например n == 100, то преобразование даст d
	fmt.Printf("n = %v string() converted: %v\n", n, input)
}
