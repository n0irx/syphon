package handler

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// ExecCommand : exec shell command via os/exec
func ExecCommand(alias string) {
	db, _ := connectDb()
	rows, err := db.Query("select command from syphon where alias=?", alias)

	checkError(err)

	for rows.Next() {
		var command string
		rows.Scan(&command)
		if err := rows.Scan(&command); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}

		commandArgs := strings.Split(command, " ")
		out, err := exec.Command(commandArgs[0], commandArgs[1:]...).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	}
}
