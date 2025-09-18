package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var data = []Entry{}
var index = map[string]int
var db = "db.csv"

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
		t := time.Now().Format(time.RFC850)
		data = append(data, Entry{name, surname, n, t})
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

// получаем доступ ко всему срезу data и помещаем пары индекса и значения среза на карту, используя значение в качестве ключа для карты и индекс среза в качестве значения карты
func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}

	return nil
}

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)

	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	// CSV-файл читается весь сразу - ReadAll() (построчно - Read())
	// тип данных lines — [][]string
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// запись в CSV-файл
func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	// изменение разделителя полей по умолчанию на табуляцию
	csvwriter.Comma = '\t'

	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
		csvwriter.Flush()
	}

	return nil
}

func deleteEntry(key string) error {
	// поиск по индексу телефонного номера, чтобы найти место записи в срезе с данными. Если его нет - сообщение об ошибке
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}

	// Если номер телефона найден, то вы удаляете соответствующую запись из среза data
	// склеивания в новый срез части до элемента и после`
	data = append(data[:i], data[i+1:]...)

	// обновить индекс (удалить из нег запись)‚ поскольку забота о нем — та цена, которую вы платите за дополнительную скорость, возникающую благодаря ему
	delete(index, key)

	// сохранить обновленные данные
	err := saveCSVFile(db)
	if err != nil {
		return err
	}

	return nil
}

// toDo: matchTel - аналогично matchInt из regexp/intRE.go. учесть, что возможен +7... или прочие знаки кроме цифр
func matchTel(tel string) bool {

}

// toDo: это копипаст из другого файла. доработать
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}

	return Entry{Name: N, Surname: S, Year: Y}
}

// toDo
func insert() {

}

// телефонная книга заполняется из csv-файла
func main() {
	arguments := os.Args
	if len((arguments)) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}
	// если файла не существует - создаём пустой
	_, err := os.Stat(db)
	if err != nil {
		fmt.Println("creating", db)
		// создаём файл
		f, err := os.Create(db)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(db)
	// это обычный файл UNIX?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(db, "not a regular UNIX file")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index")
		return
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert <Name>, <Surname> <Telephone>")
			return
		}

		t := strings.ReplaceAll(arguments[4], "-", "")

		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		temp := initS(arguments[2], arguments[3], t)

		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete <Number>")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search <Number>")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}

		fmt.Println(*temp)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}

/*
телефонная книгка запоняется хардкодом

func main() {
	args := os.Args

	if len(args) == 1 {
		// если аргументов нет, то завершаем приложение
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <args>\n", exe)
		return
	}

	t := time.Now().Format(time.RFC850)
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471", t})
	data = append(data, Entry{"Mary", "Doe", "2109416871", t})
	data = append(data, Entry{"John", "Black", "2109416123", t})
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
*/
