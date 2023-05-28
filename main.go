package main

import (
	"Calculator/calc"
	"Calculator/scan"
	"fmt"
	"os"
)

func print(sum int) {
	fmt.Println("Результат:", sum)
}
func main() {

	name, oneInt, twoInt, symbol := scan.Scan()

	sum := calc.Calc(oneInt, twoInt, symbol)

	print(sum)

	recordLine := fmt.Sprintf("Пользватель %s, ввел первое число %d, второе %d, результат: %d", name, oneInt, twoInt, sum)

	// функция создания файла

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(recordLine)

}

// калькулятор считает, записывает, файл создается с добавлением в него значений.
// Однако при каждом запуске происходит добавление последний данных, я думаю это из-за того что пересоздается и перезаписывается новый файл7 Подскажи как добавлять данные в имеющийся файл.
