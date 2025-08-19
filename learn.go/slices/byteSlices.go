package main

import "fmt"

func main() {
	// байтовый срез
	b := make([]byte, 12)
	// пустой байтовый срез. содержит 12 нулей
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €") // фраза преобразуется в байтовый срез, пригодный для записи в файл
	fmt.Println("Byte slice:", b)
	// вывод содержимого байтового среза в виде текста
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
	// длина b
	fmt.Println("Length of b", len(b)) // байтовый срез b содержит 12 символов, но его длина - 14. символ € занимает несколько байт
}
