package main

import (
	"fmt"
	"unicode/utf8"
)

// Задача 1

func main1() {

	// Объяви переменную var r rune = 'A' и выведи её как число и как символ.
	var r rune = 'A'
	fmt.Println(r)

	// Создай строку s := "Go!".
	s := "Go!"

	// Выведи: len(s) и результат utf8.RuneCountInString(s).
	fmt.Println(len(s), utf8.RuneCountInString(s))

	// Добавь строку с кириллицей: ru := "Привет"
	// и снова сравни len(ru) и utf8.RuneCountInString(ru).
	ru := "Привет"
	fmt.Println(len(ru), utf8.RuneCountInString(ru))
}
