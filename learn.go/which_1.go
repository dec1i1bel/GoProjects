package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// аналог утилиты which для unix
	// примеры команд:
	// go run which_1.go which - ищет в путях $PATH упоминания which

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("No argment provided")
		return
	}

	// file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	// срез хранит аргументы командной строки - имена файлов, которые нужно найти в путях $PATH
	files := make([]string, 0)
	i := 1

	// цикл по аргументам cmd, начиная с 1 (выбор непрерывной части среза), т.к. 0 - сама команда`
	for _, arg := range arguments[i:] {
		files = append(files, arg)
	}

	for _, directory := range pathSplit {
		// путь до исполняемого файла
		for _, file := range files {
			fullpath := filepath.Join(directory, file)
			// проверяем существование пути
			fileinfo, err := os.Stat(fullpath)

			if err == nil {
				mode := fileinfo.Mode()
				// это обычный файл?
				if mode.IsRegular() {
					// является ли он исполняемым?
					if mode&0111 != 0 {
						fmt.Println(fullpath)
					}
				}
			}
		}
	}
}
