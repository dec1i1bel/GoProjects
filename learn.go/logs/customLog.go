package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// создаём глобальную переменную LOGFILE, в которой хранится путь к каталогу кастомного журнала. os.TempDir() - каталог "/tmp", который обычно очищается при каждой перезагрузке пк. это в учебнмо коде, в боевом - что-то своё
	LOGFILE := path.Join(os.TempDir(), "mGo.log")

	// вызов os.OpenFile() создает файл журнала для записи,
	// если он еще не существует, или же открывает его для записи
	// путем добавления новых данных в конце (ос.O_APPEND)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	// функция выполнится непосредственно перед возвратом из main()
	defer f.Close()

	// добавляем в лог имя файла и строку, на которая инициировала запись
	// log.Llongfile вместо log.Lshortfile - полный путь к файлу
	LstdFlags := log.Ldate | log.Lshortfile
	// создаём файл журнала на основе открытого f и пишем в него 2 сообщения с Println. данные только пишутся в журнал, не выводятся
	iLog := log.New(f, "iLog ", LstdFlags)
	iLog.Println("mastering go 3")
	iLog.SetFlags((log.Lshortfile | log.LstdFlags))
	iLog.Println("Another log entry")
}
