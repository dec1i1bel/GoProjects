/*
утилита, работающая с двумя форматами файла csv:
• формат 1 — имя, фамилия, номер телефона, время последнего доступа
• формат 2 — имя, фамилия, код города, номер телефона, время последнего доступа
*/
package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

type F1 struct {
	Name, Surname, Tel, LastAccess string
}

type F2 struct {
	Name, Surname, AreaCode, Tel, LastAccess string
}

type Book1 []F1
type Book2 []F2

var CSVFILE = ""
var d1 = Book1{}
var d2 = Book2{}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	// Первая строка CSV-файла определяет его формат. Следовательно, нам нужна переменная флага для указания того, имеем ли мы дело с первой строкой (firstLine)
	var firstLine bool = true
	var format1 bool = true
	for _, line := range lines {
		if firstLine {
			// в строке файла либо 4, либо 5 полей. всё остальное - ошибка
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false
			} else {
				return errors.New("Unknown File Format. Only 4 or 5 fields acceptable")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				d1 = append(d1, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					Surname:    line[1],
					AreaCode:   line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				d2 = append(d2, temp)
			}
		}
	}
	return nil
}

// параметр - пустой интерфейс
func sortData(data interface{}) {
	// определяет тип данных среза, который передается в качестве пустого интерфейса этой функции путем использования переключателя типа
	switch T := data.(type) {
	case Book1:
		d := data.(Book1)
		sort.Sort(Book1(d))
		list(d)
	case Book2:
		d := data.(Book2)
		sort.Sort(Book2(d))
		list(d)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func list(d interface{}) {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		for _, v := range data {
			fmt.Println(v)
		}
	case Book2:
		data := d.(Book2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		fmt.Println("Not supported type: %T\n", T)
	}
}

/*
реализация sort.Iterface

Пакет sort содержит интерфейс sort.Interface, который позволяет сортировать срезы в соответствии с вашими потребностями и данными при условии, что вы реализуете sort.Interface для пользовательских типов данных, хранящихся в срезах. Пакет sort определяет sort.Interface так:
type Interface interface {
    // Len — это количество элементов в коллекции
    Len() int
    // Less сообщает, должен ли элемент с индексом i сортироваться перед элементом с индексом j
    Less(i, j int) bool
    // Swap меняет местами элементы с индексами i и j
    Swap(i, j int)
}
*/

// реализация sort.Interface для Book1
func (a Book1) Len() int {
	return len(a)
}

func (a Book1) Less(i, j int) bool {
	// если в записях общая фамилия - сравниваем имена
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a Book1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// реализация sort.Interface для Book2
func (a Book2) Len() int {
	return len(a)
}

func (a Book2) Less(i, j int) bool {
	// если в записях общая фамилия - сравниваем имена
	if a[i].AreaCode == a[j].AreaCode {
		return a[i].Name < a[j].Name
	}
	return a[i].AreaCode < a[j].AreaCode
}

func (a Book2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

}
