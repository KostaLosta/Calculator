package main

import (
	"fmt"
	"strconv"
)

func Scan() (int, int, string) {

	var name string
	fmt.Println("Вашему вниманию представляется универсальный калькуляторю. Ведите Ваше имя:")
	fmt.Scanf("%s\n", &name)

	var one, two string

	fmt.Println("Введите число 1:")
	fmt.Scanf("%s\n", &one)

	oneInt, err := strconv.Atoi(one)
	if err != nil {
		fmt.Println("ошибка, введите число повторно")
		fmt.Scanf("%s\n", &one)
	}

	var symbol string
	fmt.Println("Введите арифметическое действие:")
	fmt.Scanf("%s\n", &symbol)

	fmt.Println("Введите число 2:")
	fmt.Scanf("%s\n", &two)

	twoInt, err := strconv.Atoi(two)
	if err != nil {
		fmt.Println("ошибка, введите число повторно")
		fmt.Scanf("%s\n", &two)
	}
	return oneInt, twoInt, symbol

}

func calc(oneInt, twoInt int, symbol string) int {

	var sum int

	switch symbol {
	case "+":
		sum = oneInt + twoInt
	case "-":
		sum = oneInt - twoInt
	case "*":
		sum = oneInt * twoInt
	case "/":
		sum = oneInt / twoInt
	default:
		fmt.Println("ОШИБКА, неверное указано логарифмическое действие")
	}
	return sum

}

func main() {
	// ответ почему у тебя не получилось, в документации есть как создать проект в go
	// https://go.dev/doc/code
	// что я сделал точнее какие команды ввел
	//   1. go mod init Calculator
	//   2. go mod tidy
	//   3.  go install Calculator

	// что с этим делать, ответ это тоже просто запоминается на изусть действия при создании проекта на go

	//counter.CreatedPackageForBro()

	Scan() //  return 3 variables
	//counter.CreatedPackageForBro()
	calc()
	// calc() // return 1 variables

	// создать строку записи
	// recordLine  := fmt.Sprintf("Пользватель %s, ввел первое число %d второе %d, итог: %d" name, one, two, sum) тут 4 переменных откудо взялись, а почему Scan b calc вместе возвращают 4 переменных)

	// функция создания файла

	// функция записи в файл

}
