package main

import (
	"Calculator/pkg/calc"
	"Calculator/pkg/config"
	record "Calculator/pkg/recordFile"
	"Calculator/pkg/scan"
	"Calculator/pkg/sqlstore"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var (
	history = "history.txt"
)

func main() {
	// последовательность работы
	// 1. работа с файлом
	// 1.1  Создать файл, куда записывать результаты, если он уже создаy то ничего не делать
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	record.CreateFile(history)
	// строка подключения с данными

	actsDSN := fmt.Sprintf(`host=%s dbname=%s user=%s password=%s port=%s sslmode=disable`,
		config.Service().DBHost, config.Service().DBName, config.Service().DBUser, config.Service().DBPass, config.Service().DBPort)

	// подключение в бд
	actsConn, err := sqlstore.New("postgres", actsDSN, config.Service().DBActsPoolSize, config.Service().DBConnTimeout)
	if err != nil {
		_ = level.Error(logger).Log("msg", "failed to connect to acts storage", "err", err)
		return
	}

	// инициализация интрефейса который отвечает за методы которые используются для работы с базой данный ты их создаешь там и вызываешь где нужно
	store := sqlstore.NewStorage(actsConn)

	// 2. Получение данных, приведения их к типу с которым надо работать, сама работа
	name, oneInt, twoInt, symbol := scan.Scan()

	sum := calc.Calc(oneInt, twoInt, symbol)

	// вызываем метод который записывает в базу данных
	err = store.Create(name, oneInt, twoInt, sum)
	if err != nil {
		_ = level.Error(logger).Log("БРО ВМЕСТО ПРИНТОВ ЛОГГером пользуются", "err create DB", "err", err)
		return

	}
	// 3. Открывавем файл куда записывать, записываем
	record.Filewrite(history, name, oneInt, twoInt, sum)
}
