package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Задача 5

// Напиши функцию CountRunes(s string) int, которая возвращает количество рун в строке.
// Проверь на строках: "Go", "Привет", "你好", "🙂".

func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}

// Задача 6

// Напиши функцию OnlyLetters(s string) string, которая возвращает новую строку только из букв.
// Используй for _, r := range s и unicode.IsLetter(r).

func OnlyLetters(s string) string {
	slice := []rune{}

	for _, r := range s {

		if unicode.IsLetter(r) {
			slice = append(slice, r)
		}
	}
	return string(slice)
}

func main() {

	fmt.Println(OnlyLetters("asan3"))

	// Go := "Go"
	// Hello := "Привет"
	// Japan := "你好"
	// Smile := "🙂"

	// fmt.Println(CountRunes(Go))
	// fmt.Println(CountRunes(Hello))
	// fmt.Println(CountRunes(Japan))
	// fmt.Println(CountRunes(Smile))

}
