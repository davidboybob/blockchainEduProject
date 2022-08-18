package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type User struct {
	Id             int
	Name           string
	Tutorial_score int
	Last_step      string
	Create_at      string
	Update_at      string
}

func OpenDatabase() error {
	var err error

	Db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	//#region User Table 존재 여부 확인
	var isExist string
	err = Db.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE NAME="USER"`).Scan(&isExist)
	if err != nil {
		log.Fatal(err.Error())
	}

	if isExist == "0" {
		fmt.Println(`Can't find UserTable. Run CreateTable function!`)
		CreateTable()
	}
	//#endregion

	return Db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS USER (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"tutorial_score" INT default 0,
		"last_step" TEXT default tutorial,
		"create_at" datetime default CURRENT_TIMESTAMP,
		"update_at" datetime default CURRENT_TIMESTAMP
	);`

	statement, err := Db.Prepare(createTableSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer statement.Close()

	statement.Exec()
	log.Println("database table created!")

}
func InsertUser(name string) {
	insertUserSQL := `INSERT INTO User(name) VALUES (?)`
	statement, _ := Db.Prepare(insertUserSQL)

	_, err := statement.Exec(name)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted Your Info successfully")
}

// func GetUser(db *sql.DB, name string) (User, err) {
// 	var user User
// 	rows := db.QueryRow("select * from USER where id = 0")
// 	err := rows.Scan(&)
// }

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO blockchainedu(word, definition, category) VALUES(?, ?, ?)`
	statement, err := Db.Prepare((insertNoteSQL))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted edu note successfully")
}

func DisplayUser(db *sql.DB) (*User, error) {
	var user User
	row, err := db.Query("SELECT * FROM USER WHERE id = 1")
	if err != nil {
		log.Fatalln(err) // Fatal = 에러를 나타내는 출력문
		// log.Println(err) //
		return &user, err
	}

	defer row.Close()

	for row.Next() {

		err := row.Scan(&user.Id, &user.Name, &user.Tutorial_score, &user.Last_step, &user.Create_at, &user.Update_at)
		if err != nil {
			log.Println(err)
			return &User{}, nil
		}
	}

	return &user, nil
}

func UpdateUserScore(score int) {
	insertScoreSQL := `UPDATE User SET tutorial_score = (?) WHERE id = 1`
	statement, _ := Db.Prepare(insertScoreSQL)

	_, err := statement.Exec(score)
	if err != nil {
		log.Fatalln(err)
	}
}
