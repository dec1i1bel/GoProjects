package main

import (
	"proj1/db"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	db.Connect("root", "vkshmuk0707", "tcp", "127.0.0.1:3306", "bitrloc")
}
