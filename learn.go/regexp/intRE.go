package main

import (
	"fmt"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	// сопоставить что-то, что начинается с — или +‚ является необязательным (?) и заканчивается любым количеством цифр (\d+). Кроме того‚ требуется, чтобы у нас была хотя бы одна цифра до конца проверяемой строки ($)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	fmt.Println(matchInt("1fs23"))
	fmt.Println(matchInt("123"))
	fmt.Println(matchInt("/123"))
	fmt.Println(matchInt("123.2"))
}
