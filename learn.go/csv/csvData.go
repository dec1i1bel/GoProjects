package main

// запуск:
// go run csvData.go csv.data output.data
// output.data - генерируется при выполнении

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

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
func saveCSVFile(filepath string, delimiter []rune) error {
	csvfile, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = delimiter[0]

	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
		csvwriter.Flush()
	}

	return nil
}

// с помощью readCSVFile() помещает прочитанное в срез myData. lines — это срез с двумя измерениями, и каждая строка в lines уже разделена на поля.
func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("Usage: <input file name> <output file name> \"<file line fields delimiter>\"")
		return
	}

	input := args[1]
	output := args[2]
	delimiter := args[3]
	lines, err := readCSVFile(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}

		myData = append(myData, temp)
		fmt.Println("temp: ", temp)
	}

	delimiterRuned := []rune(delimiter)

	err = saveCSVFile(output, delimiterRuned)

	if err != nil {
		fmt.Println(err)
		return
	}
}
