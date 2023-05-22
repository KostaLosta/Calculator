package main

import "fmt"

func Calculator() {

	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Вашему вниманию предоставляется универсальный калькулятор? разработанный начинающим программистом")

	var x, y, z int

	fmt.Println("Введите число 1:")
	fmt.Scanf("%d\n", &x)

	var symbol string
	fmt.Println("Введите арифметическое действие:")
	fmt.Scanf("%s\n", &symbol)

	fmt.Println("Введите число 2:")
	fmt.Scanf("%d\n", &y)

	switch symbol {
	case "+":
		z = x + y
		fmt.Println(name, ", Ваш результат", z)
	case "-":
		z = x - y
		fmt.Println(name, ", Ваш результат", z)
	case "*":
		z = x * y
		fmt.Println(name, ", Ваш результат", z)
	case "/":
		z = x / y
		fmt.Println(name, ", Ваш результат", z)
	}
}
func main() {
	Calculator()
}
