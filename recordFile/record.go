package record

import (
	"fmt"
	"os"
)

func CreateFile(nameFile string) {

	if _, err := os.Stat(nameFile); os.IsNotExist(err) {
		file, err := os.Create(nameFile)
		if err != nil {
			fmt.Println("Не удалось создать файл:", err)
			os.Exit(1)
		}
		file.Close()

	}
}

func Filewrite(nameFile string, name string, oneInt, twoInt, sum int) {
	recordLine := fmt.Sprintf("Пользватель %s, ввел первое число %d, второе %d, результат: %d", name, oneInt, twoInt, sum)

	f, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.WriteString("\n")

	f.WriteString(recordLine)
}
