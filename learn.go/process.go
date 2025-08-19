package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	// проверка входных данных сначала на допустимым целым числом, потом - допустимым числом с плавающей точкой. так как каждое допустимое целое число также является допустимым значением с плавающей запятой.
	var total, nInts, nFloats int // допустимые типы
	invalid := make([]string, 0)  // срез для нечисловых значений
	i := 1

	// [i:] - выбор непрерывной части среза
	for _, k := range arguments[i:] {
		// проверяем на int
		_, err := strconv.Atoi(k)

		if err == nil {
			total++
			nInts++
			continue
		}

		// проверяем на float
		_, err = strconv.ParseFloat(k, 64)

		if err == nil {
			total++
			nFloats++
			continue
		}

		// если никуда не добавлено, значит недопустимое значение
		invalid = append(invalid, k)
		total++
	}

	fmt.Println("#total:", total, "#ints:", nInts, "#floats:", nFloats)
}
