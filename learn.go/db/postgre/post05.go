package post05

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Userdata struct {
	ID          int
	Name        string
	Surname     string
	Description string
}

var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

func openConnection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode := disable", Hostname, Port, Username, Password, Database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func exists(username string) int {
	username = strings.ToLower(username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := -1
	statement := fmt.Sprintf("SELECT 'id' FROM 'users' where username='%s'", username)
	rows, err := db.Query(statement)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Scan", err)
			return -1
		}
		userID = id
	}
	defer rows.Close()
	return userID
}

func addUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", Username)
		return -1
	}
	insertStatement := `insert into "users" ("username") values ($1)`
	// параметр $1 передаём в запрос аргументом db.Exec
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}
	insertStatement = `insert into "userdata" ("userid","name","surname","description") values ($1,$2,$3,$4)`
	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}
	return userID
}
