package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/shibukawa/configdir"
)

// Command struct object
type Command struct {
	id       int
	alias    string
	command  string
	category string
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
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
