package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var dbConnect *sql.DB
var dbConf *mysql.Config

type Conf struct {
	user   string
	passwd string
	net    string
	host   string
	name   string
}

func (con Conf) createConnection() *mysql.Config {
	dbConf := mysql.NewConfig()
	dbConf.User = con.user
	dbConf.Passwd = con.passwd
	dbConf.Net = con.net
	dbConf.Addr = con.host
	dbConf.DBName = con.name

	return dbConf
}

func Connect(user string, passwd string, net string, host string, name string) *sql.DB {
	conf := Conf{user, passwd, net, host, name}
	createCon := conf.createConnection()

	var err error
	dbConnect, err = sql.Open("mysql", createCon.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := dbConnect.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("DB Connected")

	return dbConnect
}
