package main

import "fmt"

// Задача 7

// Напиши функцию ReverseRunes(s string) string, которая переворачивает строку по рунам (а не по байтам!).
// Проверь на "Привет", "你好", "🙂👍".

func ReverseRunes(s string) string {

	r := []rune(s)
	t := make([]rune, len(r))

	for i := 0; i <= len(r)-1; i++ {
		t[i] = r[len(r)-1-i]
	}

	return string(t)
}

func main() {
	s := "привет"
	fmt.Println(ReverseRunes(s))
}
