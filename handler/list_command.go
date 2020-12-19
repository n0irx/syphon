package handler

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// GetCommands : get all shell commands from database
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
