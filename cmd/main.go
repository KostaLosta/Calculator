package main

import (
	"Calculator/pkg/calc"
	"Calculator/pkg/printResult"
	record "Calculator/pkg/recordFile"
	"Calculator/pkg/scan"
	"database/sql"
	"fmt"
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

	connStr := "user=postgres password=mypass dbname=calculat sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into calculat (username, onenumber,twonumber, resultat) values (name, oneInt, twoInt, sum)")

	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

}
