package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type user struct {
	id             int
	name           string
	tutorial_score int
	last_step      string
	create_at      string
	update_at      string
}

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS USER (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"tutorial_score" INT,
		"last_step" TEXT,
		"create_at" datetime default CURRENT_TIMESTAMP,
		"update_at" dateitme default CURRENT_TIMESTAMP
	);`

	statement, err := db.Prepare(createTableSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer statement.Close()

	statement.Exec()
	log.Println("database table created!")

}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO blockchainedu(word, definition, category) VALUES(?, ?, ?)`
	statement, err := db.Prepare((insertNoteSQL))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted edu note successfully")
}

func DisplayAllUser() {
	row, err := db.Query("SELECT * FROM USER ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}

	var users = []user{}

	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var tutorial_score int
		var last_step string
		var create_at string
		// var update_at string
		row.Scan(&id, &name, &tutorial_score, &last_step, &create_at)
		log.Println("[", name, "] ", "Score: ", tutorial_score, "Step: ", last_step, "ca: ", create_at)
		users = append(users)
	}

	return user
}
