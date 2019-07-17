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
	buildEngineerDetailsTable(db, seed)
	buildSprintLineItemTable(db, seed)
	buildWorkstreamSprintEngineerSprintLineItemMap(db, seed)

	db.Close()
}

// DB v1.0.0

func buildWorkstreamNameTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(128) NOT NULL UNIQUE, PRIMARY KEY (id))", workstreamNameTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()
	if seed {
		seedWorkstreamNameTable(db)
	}
}

func buildSprintNameTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(128) NOT NULL UNIQUE, PRIMARY KEY (id))", sprintNameTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedSprintNameTable(db)
	}
}

func buildSprintSummaryTable(db *sql.DB, seed bool) {
	query, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + sprintSummaryTable + " (id INT(10) NOT NULL AUTO_INCREMENT, workingDays INT, pointsCommitted INT, pointsAchieved INT, PRIMARY KEY(id))")
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

func buildEngineerDetailsTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id INT(10) NOT NULL AUTO_INCREMENT, firstName TEXT, lastName TEXT, email VARCHAR(128) NOT NULL UNIQUE, velocity INTEGER, PRIMARY KEY(id))", engineerDetailsTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedEngineerDetailsTable(db)
	}
}

func buildSprintLineItemTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v "+
			"(id INT(10) NOT NULL AUTO_INCREMENT, "+
			"current_availability INT, "+
			"previous_availability INT, "+
			"capacity INT, "+
			"target_points INT, "+
			"committed_points_this_sprint INT, "+
			"completed_points_this_sprint INT, "+
			"completed_points_last_sprint INT, "+
			"PRIMARY KEY (id))",
		sprintLineItemTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedSprintLineItemTable(db)
	}
}

func buildWorkstreamSprintEngineerSprintLineItemMap(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v 
		(workstream_id INT NOT NULL,
		sprint_id INT NOT NULL,
		engineer_id INT NOT NULL,
		sprint_line_item_id INT NOT NULL,
		PRIMARY KEY (workstream_id, sprint_id, engineer_id, sprint_line_item_id))`,
		workstreamSprintEngineerSprintLineItemMapTable,
	)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()
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
	query.Exec(30, 24, 0)
	query.Exec(35, 27, 21)

	query.Exec(37, 28, 16)
	query.Exec(32, 21, 5)
	query.Exec(40, 22, 15)

	query.Exec(41, 23, 20)
	query.Exec(36, 25, 17)
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

func seedEngineerDetailsTable(db *sql.DB) {
	queryString := fmt.Sprintf("INSERT INTO %v (firstName, lastName, email, velocity) VALUES (?, ?, ?, ?)", engineerDetailsTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec("Bruce", "Dickinson", "a@mail.com", 0)
	query.Exec("Steve", "Harris", "b@mail.com", 0)
	query.Exec("Nicko", "McBrain", "c@mail.com", 0)
	query.Exec("Adrian", "Smith", "d@mail.com", 0)
	query.Exec("Dave", "Murray", "e@mail.com", 0)
	query.Exec("Janick", "Gers", "f@mail.com", 0)
	query.Exec("Paul", "Di`Anno", "g@mail.com", 0)
	query.Exec("Blaze", "Bayley", "h@mail.com", 0)
	query.Exec("Clive", "Burr", "i@mail.com", 0)
	query.Exec("Dennis", "Stratton", "j@mail.com", 0)
	query.Exec("Thunderstick", "Joe", "k@mail.com", 0)
	query.Exec("Doug", "Sampson", "l@mail.com", 0)
}

func seedSprintLineItemTable(db *sql.DB) {
	queryString := fmt.Sprintf(
		`INSERT INTO %v (
			current_availability, 
			previous_availability, 
			capacity, 
			target_points, 
			committed_points_this_sprint, 
			completed_points_this_sprint, 
			completed_points_last_sprint) VALUES(?,?,?,?,?,?,?)`,
		sprintLineItemTable)

	query, err := db.Prepare(queryString)
	checkError(err)

	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)
	query.Exec(10, 10, 11, 10, 10, 10, 10)

}
