package main

import "crypto/rand"

// генерация безопасных случайных чисел пакетом crypto\rand - для задач, связанных с безопасностью

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	// Функция crypto/rand.Read() случайным образом генерирует числа, которые заполняют весь байтовый срез b. придется декодировать этот байтовый срез, используя base64.URLEncoding.EncodeToString(b)‚ чтобы получить допустимую строку без каких-либо управляющих или непечатаемых символов
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
