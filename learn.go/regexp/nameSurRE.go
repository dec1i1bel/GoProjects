package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Need arguments")
		return
	}

	for i := 1; i < len(args); i++ {
		matched := ""
		if matchNameSur(args[i]) {
			matched = "matched"
		} else {
			matched = "not matched"
		}

		fmt.Println(matched)
	}
}
