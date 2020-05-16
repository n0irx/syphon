package handler

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

// Command struct object
type Command struct {
	id       int
	alias    string
	command  string
	category string
}

// Init database
func Init() {
	db, _ := sql.Open("sqlite3", "syphon.db")
	db.Exec("create table if not exists syphon (id integer primary key autoincrement, alias text unique, command text, category text)")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// SanityCheck for sanity
func SanityCheck(message string) {
	fmt.Println(message)
}

// AddCommand add command to db
func AddCommand(alias string, command string, category string) {
	Init()
	db, _ := sql.Open("sqlite3", "syphon.db")
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into syphon (alias,command,category) values (?,?,?)")
	_, err := stmt.Exec(alias, command, category)
	checkError(err)
	tx.Commit()
	fmt.Println("Done adding command")
}

// GetCommands get commands
func GetCommands() {
	Init()
	db, _ := sql.Open("sqlite3", "syphon.db")
	rows, err := db.Query("select * from syphon")

	checkError(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "alias", "command", "category"})

	for rows.Next() {
		var tempCommand Command
		err =
			rows.Scan(&tempCommand.id, &tempCommand.alias, &tempCommand.command, &tempCommand.category)
		checkError(err)
		values := []string{strconv.Itoa(tempCommand.id), tempCommand.alias, tempCommand.command, tempCommand.category}
		table.Append(values)
	}
	table.Render()
}

// UpdateCommand update command
func UpdateCommand(db *sql.DB, id2 int, alias string, command string, category string) {
	Init()
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update syphon set alias=?,command=?,category=? where id=?")
	_, err := stmt.Exec(alias, command, category, sid)
	checkError(err)
	tx.Commit()
}

// DeleteCommand delete command
func DeleteCommand(db *sql.DB, id2 int) {
	Init()
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from syphon where id=?")
	_, err := stmt.Exec(sid)
	checkError(err)
	tx.Commit()
}
