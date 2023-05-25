package main

import (
	"fmt"
	"strconv"
)

func Scan() {

	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Вашему вниманию предоставляется универсальный калькулятор, разработанный начинающим программистом")

	var one, two, sum string

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

	switch symbol {
	case "+":
		sum = one + two
	case "-":
		sum = one - two
	case "*":
		sum = one * two
	case "/":
		sum = one / two
	default:
		fmt.Println("ОШИБКА")
	}
	fmt.Println(sum)
}

func main() {
	Scan()

}
