/*
os.Args - аргументы конандной строки. первый - всегда имя исполняемого файла. то есть os.Args никогдане пустой. остальные - аргументы, разделённые пробелами. os.Args - срез
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments")
		return
	}

	fmt.Println("run:", arguments[0])

	var min, max float64

	// начинаем с i := 1, так как i := 0 - путь к файлу билда в папке /tmp
	for i := 1; i < len((arguments)); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}

		if i == 1 {
			min = n
			max = n
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
