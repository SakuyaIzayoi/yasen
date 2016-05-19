package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	fmt.Println("Importing ship data")
	importDB()
	fmt.Println("Done")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Ship struct {
	Id   int `json:"id"`
	No   int `json:"no"`
	Name struct {
		NameKanji  string `json:"ja_jp"`
		NameKana   string `json:"ja_kana"`
		NameRomaji string `json:"ja_romaji"`
		Suffix     int    `json:"suffix"`
	} `json:"name"`
	Stat struct {
		Fire       int `json:"fire"`
		FireMax    int `json:"fire_max"`
		Torpedo    int `json:"torpedo"`
		TorpedoMax int `json:"torpedo_max"`
		AA         int `json:"aa"`
		AAMax      int `json:"aa_max"`
		Asw        int `json:"asw"`
		AswMax     int `json:"asw_max"`
		HP         int `json:"hp"`
		HPMax      int `json:"hp_max"`
		Armor      int `json:"armor"`
		ArmorMax   int `json:"armor_max"`
		Evasion    int `json:"evasion"`
		EvasionMax int `json:"evasion_max"`
		Carry      int `json:"carry"`
		Speed      int `json:"speed"`
		Range      int `json:"range"`
		Los        int `json:"los"`
		LosMax     int `json:"los_max"`
		Luck       int `json:"luck"`
		LuckMax    int `json:"luck_max"`
	} `json:"stat"`
}

func importDB() {
	file, err := os.Open("./database/db/ships.json")
	checkErr(err)

	defer file.Close()

	// Scan in the JSON, parse it, and send it to the DB
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var ship Ship

		err := json.Unmarshal([]byte(str), &ship)
		checkErr(err)

		db, err := sql.Open("sqlite3", "./test.db")
		checkErr(err)

		result, err := db.Exec(
			"INSERT INTO ship VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)",
			ship.Id,
			ship.No,
			ship.Name.NameKanji,
			ship.Name.NameRomaji,
			ship.Name.Suffix,
			ship.Stat.Fire,
			ship.Stat.FireMax,
			ship.Stat.Torpedo,
			ship.Stat.TorpedoMax,
			ship.Stat.AA,
			ship.Stat.AAMax,
			ship.Stat.Asw,
			ship.Stat.AswMax,
			ship.Stat.HP,
			ship.Stat.HPMax,
			ship.Stat.Armor,
			ship.Stat.ArmorMax,
			ship.Stat.Evasion,
			ship.Stat.EvasionMax,
			ship.Stat.Carry,
			ship.Stat.Speed,
			ship.Stat.Range,
			ship.Stat.Los,
			ship.Stat.LosMax,
			ship.Stat.Luck,
			ship.Stat.LuckMax,
		)
		checkErr(err)

		if result != nil {
			_, err := result.RowsAffected()
			checkErr(err)
		}
	}

	checkErr(scanner.Err())
}
