// реализация интерфейса error на примере обработки ошибок при чтении файла

// Когда из файла больше нечего читать, Go возвращает ошибку io.EOF, которая, строго говоря, является не ошибочным состоянием, а логической частью чтения файла. Если файл полностью пуст, то вы все равно получите io.EOF при попытке его прочитать. Однако в ряде ситуаций это может вызвать проблемы, и вам понадобится способ отличить полностью пустой файл от файла, который был прочитан полностью. Один из способов решения этой задачи — использование интерфейса error

/*
тип error - это интерфейс, определённый как:

type error interface {
    Error() string
}
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	Ended bool
	Read  int
}

// реалиуем интерфейс error для emptyFile
func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t) but read (%d) bytes", e.Ended, e.Read)
}

// проверка файла на пустоту
func isFileEmpty(e error) bool {
	// утверждение типа для получения структуры emptyFile из переменной error
	v, ok := e.(emptyFile)
	if ok {
		// если прочитано 0 байт и достигнут конец файла
		if v.Read == 0 && v.Ended == true {
			return true
		}
	}
	return false
}

func readFile(file string) error {
	var err error
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	n := 0
	for {
		line, err := reader.ReadString('\n')
		n += len(line)
		// если достигнут конец файла
		if err == io.EOF {
			if n == 0 {
				return emptyFile{true, n}
			}
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse() // парсинг флагов командной строки вида -n
	if len(flag.Args()) == 0 {
		fmt.Println("usage: errorInt <file1>[<file2>...]")
		return
	}

	for _, file := range flag.Args() {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println(file, err)
		}
	}
}
