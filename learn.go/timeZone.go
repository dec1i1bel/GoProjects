package main

import (
	"fmt"
	"time"
)

func main() {
	// toDo: тз - параграф "Работа с часовыми поясами"
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", time.Now().In(loc))
}
