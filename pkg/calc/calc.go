package calc

import (
	"fmt"
)

func Calc(oneInt, twoInt int, symbol string) int {
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
