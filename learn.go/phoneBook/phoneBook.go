package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"regexp"
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
var index = make(map[string]int)
var db = "db.csv"
var curTime = strconv.FormatInt(time.Now().Unix(), 10)

// линейный поиск по срезу data. это медленно, но достаточно при небольшом количестве записей
func search(key string) *Entry {
	/* поиск без исопльзования индекса
	for i, v := range data { // i,v - индекс и значение элемента
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
	*/

	// поиск использует индекс:
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = curTime
	return &data[i]
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
		return fmt.Errorf("%s cannot be found", key)
	}

	// Если номер телефона найден, то вы удаляете соответствующую запись из среза data
	// путём склеивания в новый срез части до элемента и после`
	data = append(data[:i], data[i+1:]...)

	// обновить индекс (удалить из него запись). забота о нем — та цена, которую вы платите за дополнительную скорость, возникающую благодаря ему
	delete(index, key)

	// сохранить обновленные данные
	err := saveCSVFile(db)
	if err != nil {
		return err
	}

	return nil
}

func insert(pS Entry) error {
	// если запись уже есть - не добавляем
	_, ok := index[(pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	data = append(data, pS)

	/* заполнение хардкодом
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471", t})
	data = append(data, Entry{"Mary", "Doe", "2109416871", t})
	data = append(data, Entry{"John", "Black", "2109416123", t})
	*/

	// обновить индекс
	_ = createIndex()
	err := saveCSVFile(db)
	if err != nil {
		return err
	}
	return nil
}

func matchTel(tel string) bool {
	t := []byte(tel)
	ctel := regexp.MustCompile(`^+?\d+$`)
	return ctel.Match(t)
}

func initS(Name, Surname, Tel string) Entry {
	isTelMatch := matchTel(Tel)
	if !isTelMatch {
		fmt.Println("incorrect phone")
		return Entry{}
	}

	return Entry{Name: Name, Surname: Surname, Tel: Tel, LastAccess: curTime}
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

	lines, err := readCSVFile(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		linePrepared := strings.Fields(line[0]) // срез из строк ,разделённых пробелом
		temp := Entry{
			Name:       linePrepared[0],
			Surname:    linePrepared[1],
			Tel:        linePrepared[2],
			LastAccess: linePrepared[3],
		}

		data = append(data, temp)
	}

	err = createIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}

		t := strings.ReplaceAll(arguments[4], "-", "")

		entry := initS(arguments[2], arguments[3], t)
		fmt.Println()
		if entry.Name != "" {
			err := insert(entry)
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
