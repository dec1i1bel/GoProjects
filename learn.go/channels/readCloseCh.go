package main

import "fmt"

func main() {
	// Если читать закрытый канал без буферизации, то будет сбой
	willClose := make(chan complex64, 10)
	willClose <- -1
	willClose <- 1i
	// считывание данных из канала
	<-willClose
	<-willClose

	close(willClose)

	read := <-willClose // читаем из закрытого канала, получаем нулевое значение типа complex64
	fmt.Println(read)
}
