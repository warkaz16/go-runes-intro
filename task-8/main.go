package main

import (
	"fmt"
	"unicode/utf8"
)

// Задача 8

// Функция SubstrRunes(s string, start, length int) string — вернуть подстроку по индексам рун.
// Проверь на "Интукод": возьми 4 руны, начиная с руны с индексом 2. Должно получиться туко.

func SubstrRunes(s string, start, length int) string {
	runes := []rune(s)

	if utf8.RuneCountInString(s)-2 < length || start < 0 {
		return "0"
	}

	end := start + length

	return string(runes[start:end])
}

func main() {
	fmt.Println(SubstrRunes("Интукод", 2, 4))
}
