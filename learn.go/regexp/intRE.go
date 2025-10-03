package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Need arguments")
		return
	}

	numTrues := 0
	numFalses := 0

	for i := 1; i < len(args); i++ {
		matched := ""
		if matchInt(args[i]) {
			matched = "matched"
			numTrues++
		} else {
			matched = "not matched"
			numFalses++
		}

		fmt.Println(matched)
	}

	fmt.Printf("trues: %d, falses: %d", numTrues, numFalses)

}
