package data

import (
	"database/sql"
	"fmt"
)

func dbBuilder(seed bool) {
	db := getDatabase()

	buildWorkstreamNameTable(db, seed)
	buildSprintNameTable(db, seed)
	buildSprintSummaryTable(db, seed)
	buildWorkstreamSprintNameSprintSummaryMapTable(db, seed)

	db.Close()
}

// DB v1.0.0

func buildWorkstreamNameTable(db *sql.DB, seed bool) {
	query, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + workstreamNameTable + "(id INTEGER PRIMARY KEY, name TEXT, UNIQUE(name))")
	checkError(err)
	query.Exec()
	if seed {
		seedWorkstreamNameTable(db)
	}
}

func buildSprintNameTable(db *sql.DB, seed bool) {
	query, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + sprintNameTable + " (id INTEGER PRIMARY KEY, name TEXT, UNIQUE(name))")
	checkError(err)
	query.Exec()

	if seed {
		seedSprintNameTable(db)
	}
}

func buildSprintSummaryTable(db *sql.DB, seed bool) {
	query, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + sprintSummaryTable + " (id INTEGER PRIMARY KEY, workingDays INTEGER, pointsCommitted INTEGER, pointsAchieved INTEGER)")
	checkError(err)
	query.Exec()

	if seed {
		seedSprintSummaryTable(db)
	}
}

func buildWorkstreamSprintNameSprintSummaryMapTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (workstreamId INTEGER NOT NULL, sprintNameId INTEGER NOT NULL, sprintSummaryId INTEGER NOT NULL, PRIMARY KEY (workstreamId, sprintNameId, sprintSummaryId))",
		workstreamSprintNameSprintSummaryMapTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedWorkstreamSprintNameSprintSummaryMapTable(db)
	}
}

func seedWorkstreamNameTable(db *sql.DB) {
	query, err := db.Prepare("INSERT INTO " + workstreamNameTable + " (name) VALUES (?)")
	checkError(err)
	query.Exec("Workstream AAA000")
	query.Exec("Workstream AAA001")
	query.Exec("Workstream AAA002")
}

func seedSprintNameTable(db *sql.DB) {
	query, err := db.Prepare("INSERT INTO " + sprintNameTable + " (name) VALUES (?)")
	checkError(err)
	query.Exec("2019.06.20")
	query.Exec("2019.07.04")
	query.Exec("2019.07.17")
}

func seedSprintSummaryTable(db *sql.DB) {
	query, err := db.Prepare("INSERT INTO " + sprintSummaryTable + " (workingDays, pointsCommitted, pointsAchieved) VALUES(?, ?, ?)")
	checkError(err)
	query.Exec(34, 26, 13)
	query.Exec(30, 22, 0)
	query.Exec(35, 27, 20)

	query.Exec(37, 28, 15)
	query.Exec(32, 21, 5)
	query.Exec(40, 22, 15)

	query.Exec(40, 23, 20)
	query.Exec(35, 25, 15)
	query.Exec(45, 30, 10)
}

func seedWorkstreamSprintNameSprintSummaryMapTable(db *sql.DB) {
	query, err := db.Prepare("INSERT INTO " + workstreamSprintNameSprintSummaryMapTable + " (workstreamId, sprintNameId, sprintSummaryId) VALUES(?, ?, ?)")
	checkError(err)
	query.Exec(1, 1, 1)
	query.Exec(1, 2, 2)
	query.Exec(1, 3, 3)

	query.Exec(2, 1, 4)
	query.Exec(2, 2, 5)
	query.Exec(2, 3, 6)

	query.Exec(3, 1, 7)
	query.Exec(3, 2, 8)
	query.Exec(3, 3, 9)
}
