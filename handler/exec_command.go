package handler

import (
	"log"
	"os"
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
		cmd := exec.Command(commandArgs[0], commandArgs[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}

	}
}
