package main

import (
	"fmt"
	s "strings" // псевдоним
	"unicode"
)

// создаём глобальный псевдоним для часто используемой функции
var f = fmt.Printf

func main() {
	// strings.EqualFold() сравнивает две строки без учета их регистра
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALi"))
	// проверки с учётом регистра
	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))

	// strings.Fields - разбивает строку на срез, используя пробел в качестве разделителя
	t := s.Fields("This is a string")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	// "" в качестве разделителя -> посимвольная обработка строки как среза
	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1)) // отрицательный 3ий параметр - кол-во замен не ограничено
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	f("SplitAfter: %s\n", s.SplitAfter("123+432++", "++"))
	f("SplitAfter: %s\n", s.SplitAfter("123+432++87980", "++"))

	// функция обрезки оставляет в строке только буквы, удаляя цифры и всё прочее, на основе unicode.IsLetter()
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("trimFunc: %s\n", s.TrimFunc("123 abc ABC \t.", trimFunction))
}
