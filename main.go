package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" //ОБЯЗАТЕЛЬНАЯ СТРОЧКА! Подключает драйвер sqlite3
	//Нижнее подчеркивание говорит о том, что функции из этого пакета использоваться не будут
	//И импорт в таких случаях нужен для запуска метода init() внутри пакета
	//В нашем случае он регистрирует sqlite3, который используется в sql.Open()
)

type Gopher struct {
	name           string
	age            int
	dateOfCatching time.Time
}

func (u *Gopher) incAge() {
	u.age++
}

func (u Gopher) speak() {
	fmt.Println("Go B")
}

func main() {
	dbPath := "gopher.db" //Это название файла, в котором будет храниться наша бд

	db, err := sql.Open("sqlite", dbPath) //Подключаемся к бд Появляется обьект бд через который мы будем отправлять SQL запросы
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close() //эта строчка закроет соединение с бд после завершения работы приложения

	// дальше ваш код из прошлых практик //

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS gophers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
		return
	}
	log.Printf("Table 'gophers' created.")

	for true {

		var с string

		fmt.Print("Выберите действие: \nCreate a \nRead b \nUpdate с \nDelete  d\n\n")

		fmt.Scan(&с)

		switch с {

		case "a":
			create(db)

		case "b":
			read(db)

		case "c":
			update(db)

		case "d":
			delete(db)
		}
	}
}

func create(db *sql.DB) {

	var name string
	var age int
	var email string

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	fmt.Print("Введите почту: Наверное она нужна сусликам")
	fmt.Scan(&email)

	insertSQL := `INSERT INTO gophers (name, email, age) VALUES (?, ?, ?);`
	_, err := db.Exec(insertSQL, name, email, age)
	if err != nil {
		log.Fatalf("failed to insert: %v", err)
	}
}

func read(db *sql.DB) {

	readSQL := `SELECT name, email, age 
				FROM gophers`
	rows, err := db.Query(readSQL)
	if err != nil {
		log.Fatalf("failed to insert: %v", err)
	}
	defer rows.Close()

	for rows.Next() { // Для каждой строки резултата
		var name, email string
		var age int

		err := rows.Scan(&name, &email, &age)
		if err != nil {
			log.Fatal(err)
		}

		// 4. Выводим
		fmt.Printf("Имя: %s | Email: %s | Возраст: %d\n", name, email, age)
	}
}

func update(bd *sql.DB) { //NEN

	var name string
	var age int
	var email string
	var id int

	fmt.Print("Введите id для редактирования: ")
	fmt.Scan(&id)
	fmt.Print("Введите новое имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите новый возраст: ")
	fmt.Scan(&age)
	fmt.Print("Введите новую почту: ")
	fmt.Scan(&email)

	updateSQL := `UPDATE gophers SET name = ?, email = ?, age = ? WHERE id = ?;`
	_, err := bd.Exec(updateSQL, name, email, age, id)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}
}

func delete(bd *sql.DB) {

	var id int

	fmt.Print("Введите id для удаления: ")
	fmt.Scan(&id)

	deleteSQL := `DELETE FROM gophers WHERE id = ?;`
	_, err := bd.Exec(deleteSQL, id)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}
}
