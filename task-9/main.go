package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Задача 9. Очистка пароля

// Удалить все пробельные символы (используй unicode.IsSpace).
// Запретить непечатные символы. (используй unicode.IsControl().
// Перевести латинские буквы в верхний регистр.
// Сохранить остальные руны как есть (в т.ч. эмодзи, кириллицу, знаки).

func CleanPassword(s string) string {
	slice := []rune{}

	for _, r := range s {

		if unicode.IsSpace(r) || unicode.IsControl(r) {
		} else {
			slice = append(slice, r)
		}
	}
	return strings.ToUpper(string(slice))
}

func main() {
	fmt.Println(CleanPassword("tu \nt u k🙂"))
}
