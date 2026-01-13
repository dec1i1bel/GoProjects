package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var RAND_MIN = 1
var RAND_MAX = 999

func main() {
	args := os.Args
	cnt := len(args)
	if cnt != 2 {
		fmt.Println("Please provide number of goroutines to create")
		return
	}

	routinesNumber := args[1]
	grNum, err := strconv.Atoi(routinesNumber)
	if err != nil {
		fmt.Println("Invalid argument: goroutines number")
		return
	}

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < grNum; i++ {
		waitGroup.Add(1) // перед каждой запускаемой горутиной увеличиваем их счётчик на 1 для избежания состояния гонки
		go func() {
			defer waitGroup.Done() // уменьшение счётчика после завершения горутины (Done() == Add(-1))
			x := rand.Intn(RAND_MAX-RAND_MIN) + RAND_MIN
			fmt.Printf("%d ", x)
		}()
	}
	fmt.Printf("\n%#v\n", waitGroup)
	waitGroup.Wait() // ожидает, пока счётчик не станет == 0, потом возвращается.
	fmt.Println("exiting...")
}
