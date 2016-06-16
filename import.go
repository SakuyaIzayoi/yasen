package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

/*
Ship struct
	Contains all ship details including stats, equippables, and remodel
	information
*/
type Ship struct {
	ID   int `json:"id"`
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

/*
Item struct
	Contains item details such as stat bonuses and equip groups
*/
type Item struct {
	ID     int `json:"id"`
	Rarity int `json:"rarity"`
	Type   int `json:"type"`
	Name   struct {
		NameKanji string `json:"ja_jp"`
	} `json:"name"`
	Stat struct {
		Fire     int  `json:"fire"`
		Torpedo  int  `json:"torpedo"`
		Bomb     int  `json:"bomb"`
		Asw      int  `json:"asw"`
		AA       int  `json:"aa"`
		Armor    int8 `json:"armor"`
		Evasion  int8 `json:"evasion"`
		Hit      int8 `json:"hit"`
		Los      int  `json:"los"`
		Range    int  `json:"range,omitempty"`
		Distance int  `json:"distance"`
	} `json:"stat"`
}

func importItems() {
	file, err := os.Open("./database/db/items.json")
	checkErr(err)
	defer file.Close()

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)

	_, err = db.Exec("BEGIN TRANSACTION")
	checkErr(err)
	defer db.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var item Item

		fmt.Println(str)
		err := json.Unmarshal([]byte(str), &item)
		checkErr(err)

		_, err = db.Exec(
			"INSERT INTO item VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)",
			item.ID,
			item.Rarity,
			item.Type,
			item.Name.NameKanji,
			item.Stat.Fire,
			item.Stat.Torpedo,
			item.Stat.Bomb,
			item.Stat.Asw,
			item.Stat.AA,
			item.Stat.Armor,
			item.Stat.Evasion,
			item.Stat.Hit,
			item.Stat.Los,
			item.Stat.Range,
			item.Stat.Distance,
		)
		checkErr(err)
	}
	checkErr(scanner.Err())

	_, err = db.Exec("END TRANSACTION")
	checkErr(err)
}

func importShips() {
	file, err := os.Open("./database/db/ships.json")
	checkErr(err)
	defer file.Close()

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)

	_, err = db.Exec("BEGIN TRANSACTION")
	checkErr(err)
	defer db.Close()

	// Scan in the JSON, parse it, and send it to the DB
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var ship Ship

		err := json.Unmarshal([]byte(str), &ship)
		checkErr(err)

		if ship.Name.NameRomaji == "" {
			ship.Name.NameRomaji = ship.Name.NameKanji
		}

		_, err = db.Exec(
			"INSERT INTO ship VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)",
			ship.ID,
			ship.No,
			ship.Name.NameKanji,
			strings.Title(ship.Name.NameRomaji)+getSuffixString(ship.Name.Suffix),
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
	}
	checkErr(scanner.Err())

	_, err = db.Exec("END TRANSACTION")
	checkErr(err)
}
