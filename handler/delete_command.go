package handler

import (
	"fmt"
	"strconv"
)

// DeleteCommandByID : delete shell command by id
func DeleteCommandByID(id int) {
	db, _ := connectDb()
	sid := strconv.Itoa(id)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from syphon where id=?")
	row, err := stmt.Exec(sid)

	checkError(err)

	affectedRow, err := row.RowsAffected()

	if affectedRow == 0 {
		fmt.Printf("Error: No command with id=%d", id)
	} else {
		tx.Commit()
		fmt.Printf("Success: Command with id=%d has been deleted", id)
	}

}

// DeleteCommandByAlias : delete shell command by alias
func DeleteCommandByAlias(alias string) {
	db, _ := connectDb()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from syphon where alias=?")
	row, err := stmt.Exec(alias)

	checkError(err)

	affectedRow, err := row.RowsAffected()

	if affectedRow == 0 {
		fmt.Printf("Error: No command with alias='%s'", alias)
	} else {
		tx.Commit()
		fmt.Printf("Success: Command with alias='%s' has been deleted", alias)
	}
}
