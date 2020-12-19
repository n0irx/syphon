package handler

import "fmt"

// AddCommand add command to db
func AddCommand(alias string, command string, category string) {
	db, _ := connectDb()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into syphon (alias,command,category) values (?,?,?)")
	_, err := stmt.Exec(alias, command, category)
	checkError(err)
	tx.Commit()
	fmt.Printf("Success: Command '%s' with alias %s on category: %s has been added.", command, alias, category)
}
