package main

import (
	"fmt"
	"strings"
)

// читаем запись целиком и разделяем ее перед выполнением проверки. Кроме того, мы проводим дополнительную проверку, чтобы убедиться в том, что обрабатываемая запись содержит нужное количество полей. Каждая запись должна содержать три поля: имя, фамилию и номер телефона
func matchRecord(s string) bool {
	fields := strings.Split(s, ",")

	if len(fields) != 3 {
		return false
	}

	if !matchNameSur(fields[0]) {
		return false
	}

	if !matchNameSur(fields[1]) {
		return false
	}

	return matchTel(fields[2])
}

func main() {
	fmt.Println(matchRecord("will,smith,6786767"))
	fmt.Println(matchRecord("will,smith"))
	fmt.Println(matchRecord("smith,789677867"))
}
