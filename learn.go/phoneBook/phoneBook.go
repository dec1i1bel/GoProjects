package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// шаблон файла csv: templates/CSVFILE.csv
type Entry1 struct {
	Name, Surname, Tel string
	LastAccess         int64
}

// шаблон файла csv: templates/CSVFILE2.csv
type Entry2 struct {
	Name, Surname, AreaCode, Tel string
	LastAccess                   int64
}

/* Сортировка записей через sort.Interface */
// нужно иметь отдельный тип данных‚ для которого реализован sort.Interface
type PhoneBook1 []Entry1
type PhoneBook2 []Entry2

var data1 = PhoneBook1{}
var data2 = PhoneBook2{}

var indexByPhone = make(map[string]int)     // индекс бд csv по номеру телефона
var indexByLastAccess = make(map[int64]int) // индекс бд csv по полю LastAccess
var CSVFILE = "CSVFILE.csv"
var fieldsDelimiter string = "#"

// var curTime = strconv.FormatInt(time.Now().Unix(), 10)
var curTime = time.Now().Unix()

// реализация sort.Interface
func (a PhoneBook1) Len() int {
	return len(a)
}

func (a PhoneBook1) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a PhoneBook1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PhoneBook2) Len() int {
	return len(a)
}

func (a PhoneBook2) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a PhoneBook2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

/* end */

func setSCVFILE() error {
	filepath := os.Getenv("PHONEBOOK") // получаем значение переменной среды PHONEBOOK
	if filepath != "" {
		CSVFILE = filepath
	}

	// если файла не существует - создаём пустой
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("creating...", CSVFILE) // создание файла в существующей папке
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return err
		}
		f.Close()
		// populate(5)
	}

	fileInfo, _ := os.Stat(CSVFILE)
	// это обычный файл UNIX?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular UNIX file")
		return err
	}

	return nil
}

func searchByPhone(key string) *Entry1 {
	i, ok := indexByPhone[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = curTime
	return &data[i]
}

func list(d interface{}) {
	switch T := d.(type) {
	case PhoneBook1:
		data := d.(PhoneBook1)
		for _, v := range data {
			fmt.Println(v)
		}
	case PhoneBook2:
		data := d.(PhoneBook2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		fmt.Println("Not supported type: %T\n", T)
	}
}

// заполнение рандомными данными
func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		insert(Entry1{name, surname, n, curTime})
	}
}

// объяс-е кода - в randomValuesGenerating/newPass.go
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

// объяс-е кода - в randomValuesGenerating/randomNumbers.go
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createIndexByPhone() error {
	indexByPhone = make(map[string]int)
	for i, k := range data1 {
		key := k.Tel
		indexByPhone[key] = i
	}

	return nil
}

func createIndexByLastAccess() error {
	indexByLastAccess := make(map[int64]int)

	for k, v := range data1 {
		indexByLastAccess[v.LastAccess] = k
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

	lines, err := csv.NewReader(f).ReadAll() // построчно - Read()
	if err != nil {
		return [][]string{}, err
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
				return [][]string{}, errors.New("Unknown File Format. Only 4 or 5 fields acceptable")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				la, _ := strconv.ParseInt(line[3], 10, 64)
				temp := Entry1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: la,
				}
				data1 = append(data1, temp)
			}
		} else {
			if len(line) == 5 {
				la, _ := strconv.ParseInt(line[3], 10, 64)
				temp := Entry2{
					Name:       line[0],
					Surname:    line[1],
					AreaCode:   line[2],
					Tel:        line[3],
					LastAccess: la,
				}
				data2 = append(data2, temp)
			}
		}
	}

	return lines, nil
}

// параметр - пустой интерфейс
func sortData(data interface{}) {
	// определяет тип данных среза, который передается в качестве пустого интерфейса этой функции путем использования переключателя типа
	switch T := data.(type) {
	case PhoneBook1:
		d := data.(PhoneBook1)
		sort.Sort(PhoneBook1(d))
		list(d)
	case PhoneBook2:
		d := data.(PhoneBook2)
		sort.Sort(PhoneBook2(d))
		list(d)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

// запись в CSV-файл
func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	curTimeStr := ""
	csvwriter.Comma = '#'

	for _, row := range data1 {
		curTimeStr = strconv.FormatInt(row.LastAccess, 10)
		temp := []string{row.Name, row.Surname, row.Tel, curTimeStr}
		_ = csvwriter.Write(temp)
		csvwriter.Flush()
	}

	return nil
}

func deleteEntryByPhone(key string) error {
	// поиск по индексу телефонного номера, чтобы найти место записи в срезе с данными. Если его нет - сообщение об ошибке
	i, ok := indexByPhone[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}

	// Если номер телефона найден, то вы удаляете соответствующую запись из среза data
	// путём склеивания в новый срез части до элемента и после`
	data1 = append(data1[:i], data1[i+1:]...)

	// обновить индекс (удалить из него запись)
	delete(indexByPhone, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func insert(entry Entry1) error {
	// если запись уже есть - не добавляем
	_, ok := indexByPhone[(entry).Tel]
	if ok {
		return fmt.Errorf("%s already exists", entry.Tel)
	}

	data1 = append(data1, entry)

	// обновить индекс
	_ = createIndexByPhone()
	_ = createIndexByLastAccess()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func matchPhone(phone string) bool {
	t := []byte(phone)
	ctel := regexp.MustCompile(`^+?\d+$`)
	return ctel.Match(t)
}

func initS(Name, Surname, Phone string) Entry1 {
	if !matchPhone(Phone) {
		fmt.Println("incorrect phone")
		return Entry1{}
	}

	return Entry1{Name: Name, Surname: Surname, Tel: Phone, LastAccess: curTime}
}

// телефонная книга заполняется из csv-файла
func main() {
	setSCVFILE()

	arguments := os.Args
	if len((arguments)) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	lines, err := readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	lastAccessRaw := ""

	for _, line := range lines {
		lineRaw := strings.Fields(line[0]) // срез из строк ,разделённых пробелом
		linePrepared := strings.Split(lineRaw[0], fieldsDelimiter)
		lastAccessRaw = linePrepared[3]
		lastAccess, err := strconv.ParseInt(lastAccessRaw, 10, 64)

		if err != nil {
			fmt.Printf("Incorrect value lastAccess: %d", lastAccessRaw)
			return
		}

		temp := Entry1{
			Name:       linePrepared[0],
			Surname:    linePrepared[1],
			Tel:        linePrepared[2],
			LastAccess: lastAccess,
		}

		data1 = append(data1, temp)
	}

	errIndPhone := createIndexByPhone()
	if errIndPhone != nil {
		fmt.Println("error creating index by phone: ", errIndPhone)
		return
	}

	errIndLastAccess := createIndexByLastAccess()
	if errIndLastAccess != nil {
		fmt.Println("error creating index by last access: ", errIndLastAccess)
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
		if !matchPhone(t) {
			fmt.Println("Not a valid phone:", t)
			return
		}

		err := deleteEntryByPhone(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search <Number>")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchPhone(t) {
			fmt.Println("Not a valid phone:", t)
			return
		}

		temp := searchByPhone(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}

		fmt.Println(*temp)
	case "list":
		if len(data1) == 0 {
			sortData(data2)
			list(data2)
		} else {
			sortData(data1)
			list(data1)
		}

	default:
		fmt.Println("Not a valid option")
	}
}
