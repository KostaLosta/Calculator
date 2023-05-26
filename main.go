package main

import (
	"fmt"
	"strconv"
)

//
// func Scan() {   задаа функции получить значения от пользователя,  должна вернуть значения какие...
//
// var name string
// fmt.Println("Как тебя зовут?")
// fmt.Scanf("%s\n", &name)
//
// fmt.Println("Вашему вниманию предоставляется универсальный калькулятор, разработанный начинающим программистом")
//
// var one, two string
// var sum int
// fmt.Println("Введите число 1:")
// fmt.Scanf("%s\n", &one)
//
// oneInt, err := strconv.Atoi(one)
// if err != nil {
// fmt.Println("ошибка, введите число повторно")
// fmt.Scanf("%s\n", &one)
// }
// var symbol string
// fmt.Println("Введите арифметическое действие:")
// fmt.Scanf("%s\n", &symbol)
//
// fmt.Println("Введите число 2:")
// fmt.Scanf("%s\n", &two)
//
// twoInt, err := strconv.Atoi(two)
// if err != nil {
// fmt.Println("ошибка, введите число повторно")
// fmt.Scanf("%s\n", &two)
// }
//
// }

// func calc(){  //   задача фукнции  значения получить, производит операцию, значение  вернуть
// switch symbol {
// case "+":
// sum = oneInt + twoInt
// case "-":
// sum = oneInt - twoInt
// case "*":
// sum = oneInt * twoInt
// case "/":
// sum = oneInt / twoInt
// default:
// fmt.Println("ОШИБКА, неверное указано логарифмическое действие")
// }
// fmt.Println(sum)
//
//
// }

func Scan() { // почему опять одна функция которая выполняет две операции?? мы же так не договаривались

	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Вашему вниманию предоставляется универсальный калькулятор, разработанный начинающим программистом")

	var one, two string
	var sum int
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
	fmt.Println(sum)
}
func main() {

	Scan() //  return 3 variables

	// calc() // return 1 variables

	// создать строку записи
	// recordLine  := fmt.Sprintf("Пользватель %s, ввел первое число %d второе %d, итог: %d" name, one, two, sum) тут 4 переменных откудо взялись, а почему Scan b calc вместе возвращают 4 переменных)

	// функция создания файла

	// функция записи в файл

}
