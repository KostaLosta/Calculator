package main

import (
	"Calculator/pcg/calc"
	"Calculator/pcg/printResult"
	record "Calculator/pcg/recordFile" // посмотри сюда  я в качестве примера, назвал по разному папку и пакет
	"Calculator/pcg/scan"
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
