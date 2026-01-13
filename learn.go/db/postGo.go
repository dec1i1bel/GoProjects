package main

import (
	"fmt"
	"math/rand"

	"github.com/dec1i1bel/post05"
)

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
	post05.Password = "mysecretpassword"
	post05.Database = "postgres"
	isEmptyDb, err := post05.CheckIfEmptyDb()

	if err != nil {
		fmt.Println("func main: error check if empty db:", err)
		return
	}

	if isEmptyDb {
		fmt.Println("func main: db is empty, no ListUsers will be called")
	}

	if !isEmptyDb {
		data, err := post05.ListUsers()

		if err != nil {
			fmt.Println("error using ListUsers:", err)
		}

		for _, v := range data {
			fmt.Println("func main - row:", v)
		}
	}

	newUserName := getString(5)

	fmt.Println("newUserName:", newUserName)

	curUserId := 0
	if !isEmptyDb {
		curUserId = post05.FindUserId(newUserName)
	}

	udata := post05.UserData{
		Username:    newUserName,
		Name:        "Test Name post05",
		Surname:     "Test Surname post05",
		Description: "Test Description post05",
	}

	newUserId := 0
	if curUserId <= 0 {
		newUserId = post05.AddUser(udata)
	}

	if newUserId > 0 {
		udata = post05.UserData{
			ID:          newUserId,
			Username:    newUserName,
			Name:        "Test",
			Surname:     "User 1",
			Description: "this night not be me",
		}

		err = post05.UpdateUser(udata)
		if err != nil {
			fmt.Println("error using UpdateUser:", err)
		} else {
			fmt.Printf("user <%s> with id <%d> updated successfully\n", newUserName, newUserId)
		}

		err = post05.DeleteUser(newUserId)
		if err != nil {
			fmt.Println("error using DeleteUser", err)
		} else {
			fmt.Printf("user <%s> with id <%d> deleted successfully\n", newUserName, newUserId)
		}
	} else {
		fmt.Println("Error adding user", udata.Username)
	}
}
