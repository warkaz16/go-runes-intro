package main

import "fmt"

// Задача 3

func main() {

	s := "Жук"

	// Выведи s[0], s[1], s[2] как числа (байты).
	// fmt.Println(s[0], s[1], s[2],)

	fmt.Println(len(s))

	// Добавь цикл for i := 0; i < len(s); i++ { fmt.Println(i, s[i]) }
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// Добавь цикл for i, r := range s { fmt.Println(i, r, string(r)) }
	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

}
