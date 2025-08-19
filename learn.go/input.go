package main

import "fmt"

// получаем пользовательский ввод
func main() {
	fmt.Printf("your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("your name is", name)
}
