package main

import "fmt"

// Задача 2

func main2()  {
	
	// Попробуй объявить руну так: var bad rune = "A" (двойные кавычки).
	// var bad rune = "A"

	// Запусти, перепиши в комментариях текст ошибки и объясни, почему так нельзя.
	// fmt.Println(bad)

	// cannot use "A" (untyped string constant) as rune value in variable declaration
	// Литерал руны пишется в одинарных кавычках

	var good rune = 'A'

	fmt.Println(good)
}
