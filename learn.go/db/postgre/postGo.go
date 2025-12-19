package main

import (
	"fmt"
	"math/rand"
	"time"

	// тестируемый пакет
	"github.com/dec1i1bel/post05"
)

// величины - хакрдкод в целях тестирования
var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// генерирование случайной строки
func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar

		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5432
	post05.Username = "postgres"
	post05.Password = ""
	post05.Database = "postgres"

	data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.New(rand.NewSource(SEED)) //  rand.Seed(SEED) устарело
	random_username := getString(5)

	t := post05.Userdata{
		Username:    random_username,
		Name:        "Test",
		Surname:     "user 1",
		Description: "test user 1 description",
	}
	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("Error adding user", t.Username)
	}
	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}
	post05.AddUser(t)
	if id == -1 {
		fmt.Println("error adding 2nd user", t.Username)
	}
	t = post05.Userdata{
		Username:    random_username,
		Name:        "Test",
		Surname:     "User 1",
		Description: "this night not be me",
	}
	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
