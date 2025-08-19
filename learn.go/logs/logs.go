package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("fatal hello world")
	}

	log.Panic("panic hello world")
}
