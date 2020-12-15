package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"github.com/shibukawa/configdir"
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
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDb() (*sql.DB, error) {

	configDirs := configdir.New("n0irx", "syphon")
	cacheDir := configDirs.QueryCacheFolder()

	err := os.MkdirAll(cacheDir.Path, 0755)

	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(cacheDir.Path, "syphon.db")
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	db.Exec("create table if not exists syphon (id integer primary key autoincrement, alias text unique, command text, category text)")

	return db, err
}

// SanityCheck for sanity
func SanityCheck(message string) {
	fmt.Println(message)
}

// AddCommand add command to db
func AddCommand(alias string, command string, category string) {
	db, _ := connectDb()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into syphon (alias,command,category) values (?,?,?)")
	_, err := stmt.Exec(alias, command, category)
	checkError(err)
	tx.Commit()
	fmt.Printf("\nCommand added \n\nCommand: \t%s \nAlias: \t%s \nCategory: \t%s", command, alias, category)
}

// GetCommands get commands
func GetCommands() {
	db, _ := connectDb()
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
	db, _ = connectDb()
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update syphon set alias=?,command=?,category=? where id=?")
	_, err := stmt.Exec(alias, command, category, sid)
	checkError(err)
	tx.Commit()
}

// DeleteCommandByID delete command by Id
func DeleteCommandByID(id2 int) {
	db, _ := connectDb()
	sid := strconv.Itoa(id2) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from syphon where id=?")
	_, err := stmt.Exec(sid)
	checkError(err)
	tx.Commit()
	fmt.Printf("Command deleted: \nid: %d", id2)
}

// DeleteCommandByAlias delete command by alias
func DeleteCommandByAlias(alias string) {
	db, _ := connectDb()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from syphon where alias=?")
	_, err := stmt.Exec(alias)
	checkError(err)
	tx.Commit()
	fmt.Printf("Command deleted: \alias: %s", alias)
}
