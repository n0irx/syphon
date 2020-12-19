package handler

import (
	"database/sql"
	"strconv"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// UpdateCommand : update command from database with new command
func UpdateCommand(db *sql.DB, id2 int, alias string, command string, category string) {
	db, _ = connectDb()
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update syphon set alias=?,command=?,category=? where id=?")
	_, err := stmt.Exec(alias, command, category, sid)
	checkError(err)
	tx.Commit()
}
