package main

import (
	"fmt"
	"sync"
)

func writeToChannel(c chan int, x int) {
	c <- x   // запись значения в канал
	close(c) // закрытие канала
}

// Записывает в канал значение, не закрывая его
func printer(ch chan bool) {
	ch <- true
}

func main() {
	c := make(chan int, 1) // создан канал с буфером размером 1. как только мы заполним буфер, мы сможем закрыть канал, а горутина продолжит свое выполнение и вернется.
	// Канал, который не буферизован, ведет себя по-другому: когда вы отправляете значение в этот канал, он блокируется до момента, пока кто-то получит это значение
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	// считываем из канала и выводим значение, не сохраняя его в отдельной переменной
	fmt.Println("Read:", <-c)

	// Определяем, закрыт ли канал. игнорируем читаемое значение — если бы канал был открыт, то значение было бы отброшено.
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	// создаём небуферизованный канал и 5 горутин без синхронизации (т.к. не исп-ем Add())
	var ch chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	n := 0
	// цикл range работает с каналами.
	for i := range ch {
		fmt.Println("range. iter: ", n)
		fmt.Println("i:", i)

		if i == true {
			n++
		}

		if n > 2 {
			fmt.Println("n", n)
			close(ch)
			break
		}
	}

	// Чтение из закрытого канала возвращает нулевое значение его типа данных (здесь - false)
	fmt.Println("Reading closed channel:", <-ch)
}
