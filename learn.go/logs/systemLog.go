package main

import (
	"log"
	"log/syslog"
)

// добавление записей в системный журнал unix

func main() {
	// создаём запись. в линуксе посмотреть системный журнал - journalctl -xe
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		log.Println(err) // добавляем в него запись об ошибке
		return
	} else {
		// добавляем в него запись "Everything is fine"
		log.SetOutput(sysLog)
		log.Print("Everything is fine")
	}
}
