package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Printf("printme: param <%d>", x)
}

// порядок выполнения горутин может меняться
func main() {
	go func(x int) {
		fmt.Printf("goroutine created. param: %d", x)
	}(10) // 10 - способ, с помощью которого передается параметр анонимной функции

	go printme(15) // выполенние функции с помощью горутины

	// плохая практика
	time.Sleep(time.Second)
	// ---

	fmt.Println("exiting...")
}
