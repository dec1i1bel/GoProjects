package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i < finish; i++ {
		fmt.Println(i, "")
	}
	fmt.Println("===============")
	// time.Sleep - неправильный способ синхронизации горутин
	time.Sleep(100 * time.Microsecond)
}

func main() {
	for i := 0; i < 5; i++ {
		go myPrint(i, 5) // создание горутины
	}

	time.Sleep(time.Second)
}
