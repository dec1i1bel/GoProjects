package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

// линейный поиск по срезу data. это медленно, но достаточно при небольшом количестве записей
func search(key string) *Entry {
	for i, v := range data { // i,v - индекс и значение элемента
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

// список записей среза
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

// заполнение рандомными данными
func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
	}
}

// объяс-е кода ф-ии - в randomValuesGenerating/newPass.go
// toDo: ф-я ген-ет строку только из символов a-z,A-Z (т.к. в имени и фамилии дргуих не бывает)
func getString(len int) string {
	temp := ""
	startChar := "!"

	var min int = 0
	var max int = 94

	for i := 1; i <= len; i++ {
		myRand := random(min, max)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
	}

	return temp
}

// объяс-е кода ф-ии - в randomValuesGenerating/randomNumbers.go
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	args := os.Args

	if len(args) == 1 {
		// если аргументов нет, то завершаем приложение
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <args>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})
	populate(3)

	// какая команда была введена?
	switch args[1] {
	// поиск
	case "search":
		if len(args) != 3 {
			fmt.Println("Usage: search phone")
			return
		}

		result := search(args[2])

		if result == nil {
			fmt.Println("Entry not found: ", args[2])
			return
		}

		fmt.Println(*result)
	// вывод списка
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
