package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func getFullShipName(id int) string {
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	rows, err := db.Query("SELECT name, name_en, suffix FROM ship WHERE shipid = $1", id)
	checkErr(err)

	var name string
	var name_en string
	var suffix string

	for rows.Next() {
		var suffixIndex int

		err := rows.Scan(&name, &name_en, &suffixIndex)
		checkErr(err)

		if strings.Compare("", name_en) == 0 {
			name_en = name
		}

		switch suffixIndex {
		case 1:
			suffix = " Kai"
		case 5:
			suffix = " Kou Kai"
		case 4:
			suffix = " Kou"
		case 6:
			suffix = " Kou Kai Ni"
		case 9:
			suffix = " Kai Ni A"
		case 2:
			suffix = " Kai Ni"
		case 8:
			suffix = " drei"
		case 3:
			suffix = " Kou"
		case 10:
			suffix = " Kai Ni Otsu"
		case 7:
			suffix = " zwei"
		default:
			suffix = ""
		}
	}
	return name_en + suffix
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
