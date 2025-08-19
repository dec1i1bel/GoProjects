package main

/*
примеры команд cmd:
go run dates.go 14:10
	   "14 December 2020"
*/

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}
	// Это лишь дата?
	dateString := os.Args[1]
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time", d.Day(), d.Month(), d.Year())
	}

	//  Это дата + время?
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Это дата + время с месяцем, представленным в виде числа?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Это лишь время?
	d, err = time.Parse("15:04", dateString) // символ : должен присутствовать в проверяемой строке
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// как работать со временем эпохи UNIX: Текущую дату и время в формате unix можно получить, используя time.Now().Unix(), после чего преобразовать ее в значение time.Time с помощью вызова time.Unix().
	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// преобразуем время в time.Time типа времени Unix
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(start)
	fmt.Println("Execution time:", duration)
}
