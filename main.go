package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Calculator() {

	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Вашему вниманию предоставляется универсальный калькулятор, разработанный начинающим программистом")

	var one, two, sum string

	fmt.Println("Введите число 1:")
	fmt.Scanf("%d\n", &one)

	fmt.Println(reflect.TypeOf(one))

	oneInt, err := strconv.Atoi(one)
	if err != nil {
		fmt.Println("ошибка, введите число повторно")
		fmt.Scanf("%s\n", &one)
	}

	fmt.Println(reflect.TypeOf(oneInt))

	var symbol string
	fmt.Println("Введите арифметическое действие:")
	fmt.Scanf("%s\n", &symbol)

	fmt.Println("Введите число 2:")
	fmt.Scanf("%d\n", &two)

	fmt.Println(reflect.TypeOf(two))

	//twoInt, err := strconv.Atoi(two)
	//if err != nil {
	//	fmt.Println("ошибка, введите число повторно")
	//	fmt.Scanf("%s\n", &two)
	//}
	//fmt.Println(reflect.TypeOf(twoInt))

	switch symbol {
	case "+":
		sum = one + two
		fmt.Println(name, ", Ваш результат", sum)
	case "-":
		sum = one - two
		fmt.Println(name, ", Ваш результат", sum)
	case "*":
		sum = one * two
		fmt.Println(name, ", Ваш результат", sum)
	case "/":
		sum = one / two
		fmt.Println(name, ", Ваш результат", sum)
	}
}
func main() {
	Calculator()
}
