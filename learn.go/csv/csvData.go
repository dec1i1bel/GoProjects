package main

import (
	"encoding/csv"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Number  string
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

	// CSV-файл читается весь сразу - ReadAll() (построчно - Read())
	// тип данных lines — [][]string
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
