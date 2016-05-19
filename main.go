package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM ship")
	checkErr(err)

	for rows.Next() {
		var id int
		var num int
		var name string
		err := rows.Scan(&id, &num, &name)
		checkErr(err)
		fmt.Printf("%s\n", name)
	}

	checkErr(rows.Err())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
