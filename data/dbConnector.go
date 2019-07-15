package data

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/zenofbeer/go-zen-velocity/configuration"
	// need to force import
	_ "github.com/mattn/go-sqlite3"
)

var config = configuration.GetConfig()

func getAllWorkstreamNames() []byte {
	buildDatabase()
	fmt.Println("database built")
	initDatabaseAndSeed()
	db := getDatabase()

	rowCount, _ := db.Query("SELECT COUNT(*) as count FROM  workstreamnames")
	nameCount := checkCount(rowCount)
	rows, _ := db.Query("SELECT id, name FROM workstreamnames")
	var myID int
	var name string

	names := make([]WorkstreamName, nameCount)
	i := 0
	for rows.Next() {
		rows.Scan(&myID, &name)
		names[i] = WorkstreamName{
			ID:   myID,
			Name: name,
		}
		i++
	}

	data := WorkstreamNameList{
		ListTitle:       "The list title",
		WorkstreamNames: names,
	}

	rows.Close()
	rowCount.Close()
	db.Close()

	dataJSON, _ := json.Marshal(data)
	return dataJSON
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	return count
}

func getDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "vt.db")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func buildDatabase() {
	db := getDatabase()
	query, err := db.Prepare("CREATE TABLE IF NOT EXISTS workstreamnames (id INTEGER PRIMARY KEY, name TEXT, UNIQUE(name))")
	if err != nil {
		fmt.Println(err)
	}
	query.Exec()
	db.Close()
}

func initDatabaseAndSeed() {
	db := getDatabase()

	query, _ := db.Prepare("INSERT INTO workstreamnames (name) VALUES (?)")
	query.Exec("Workstream AAA000")
	query.Exec("Workstream AAA001")
	query.Exec("Workstream AAA002")

	db.Close()
}
