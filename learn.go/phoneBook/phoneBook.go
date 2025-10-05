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
	LastAccess int64
}

var data = []Entry{}

var indexByPhone = make(map[string]int)     // индекс db.csv по номеру телефона
var indexByLastAccess = make(map[int64]int) // индекс db.csv по полю LastAccess
var db = "db.csv"

var decimalBase int = 10
var int64BitSize int = 64
var fieldsDelimiter string = "#"

// var curTime = strconv.FormatInt(time.Now().Unix(), decimalBase)
var curTime = time.Now().Unix()

func searchByPhone(key string) *Entry {
	i, ok := indexByPhone[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = curTime
	return &data[i]
}

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
		data = append(data, Entry{name, surname, n, curTime})
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
	for i, k := range data {
		key := k.Tel
		indexByPhone[key] = i
	}

	return nil
}

func createIndexByLastAccess() error {
	indexByLastAccess := make(map[int64]int)

	for k, v := range data {
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
	curTimeStr := ""
	csvwriter.Comma = '#'

	for _, row := range data {
		curTimeStr = strconv.FormatInt(row.LastAccess, decimalBase)
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
	data = append(data[:i], data[i+1:]...)

	// обновить индекс (удалить из него запись)
	delete(indexByPhone, key)

	err := saveCSVFile(db)
	if err != nil {
		return err
	}

	return nil
}

func insert(entry Entry) error {
	// если запись уже есть - не добавляем
	_, ok := indexByPhone[(entry).Tel]
	if ok {
		return fmt.Errorf("%s already exists", entry.Tel)
	}

	data = append(data, entry)

	// обновить индекс
	_ = createIndexByPhone()
	_ = createIndexByLastAccess()

	err := saveCSVFile(db)
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

func initS(Name, Surname, Phone string) Entry {
	if !matchPhone(Phone) {
		fmt.Println("incorrect phone")
		return Entry{}
	}

	return Entry{Name: Name, Surname: Surname, Tel: Phone, LastAccess: curTime}
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
		fmt.Println("creating...", db)
		f, err := os.Create(db)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, _ := os.Stat(db)
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

	lastAccessRaw := ""

	for _, line := range lines {
		lineRaw := strings.Fields(line[0]) // срез из строк ,разделённых пробелом
		linePrepared := strings.Split(lineRaw[0], fieldsDelimiter)
		lastAccessRaw = linePrepared[3]
		lastAccess, err := strconv.ParseInt(lastAccessRaw, decimalBase, int64BitSize)

		if err != nil {
			fmt.Printf("Incorrect value lastAccess: %d", lastAccessRaw)
			return
		}

		temp := Entry{
			Name:       linePrepared[0],
			Surname:    linePrepared[1],
			Tel:        linePrepared[2],
			LastAccess: lastAccess,
		}

		data = append(data, temp)
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
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
