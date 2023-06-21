package main

import (
	"Calculator/pkg/calc"
	"Calculator/pkg/printResult"
	record "Calculator/pkg/recordFile"
	"Calculator/pkg/scan"
)

var (
	history = "history.txt"
)

func main() {
	// последовательность работы
	// 1. работа с файлом
	// 1.1  Создать файл, куда записывать результаты, если он уже создаy то ничего не делать

	record.CreateFile(history)

	// 2. Получение данных, приведения их к типу с которым надо работать, сама работа
	name, oneInt, twoInt, symbol := scan.Scan()

	sum := calc.Calc(oneInt, twoInt, symbol)

	// 3. Открывавем файл куда записывать, записываем
	record.Filewrite(history, name, oneInt, twoInt, sum)

	// когда приложение отработало корректно, выводим результат
	printResult.PrintResult(sum)

}
