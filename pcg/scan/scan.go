package scan

import (
	"fmt"
	"strconv"
)

func Scan() (string, int, int, string) {

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
	return name, oneInt, twoInt, symbol

}
