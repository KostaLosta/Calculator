когда запускаю файл мейн из папки цмд, выводится ошибка что не найдена директория. 
чтобы работало запускаю файл мейн из папки цмд и добавляю в рабочую область папку пкдж, но тогда если делаю коммиты, то добавляются только изменения в открытой папке.

далее по поводу постгри/

во первых выдает ошибку при установке драйвера go get github.com/lib/pq и другого github.com/jackc/pgx
	

В файл main.go добавил:

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
