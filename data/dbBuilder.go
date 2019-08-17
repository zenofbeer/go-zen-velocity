package data

import (
	"database/sql"
	"fmt"
)

func dbBuilder(seed bool) {
	db := getDatabase()

	buildWorkstreamNameTable(db, seed)
	buildEngineerDetailsTable(db, seed)
	buildSprintNameTable(db, seed)
	buildSprintLineItemTable(db)
	buildWorkstreamSprintEngineerSprintLineItemMap(db)

	if seed {

		// add an empty sprint
		AddSprint(1, 1, 1)

		// copy a previous sprint
		AddSprint(1, 2, 1)
		AddSprint(1, 2, 2)
		AddSprint(1, 2, 3)
		AddSprint(1, 2, 4)

		AddSprint(2, 2, 5)
		AddSprint(2, 2, 6)
		AddSprint(2, 2, 7)
		AddSprint(2, 2, 8)

		AddSprint(2, 2, 9)
		AddSprint(2, 2, 10)
		AddSprint(2, 2, 11)
		AddSprint(2, 2, 12)
	}

	/*
		buildWorkstreamSprintNameSprintSummaryMapTable(db, seed)

	*/
	db.Close()
}

// DB v1.0.0

func buildWorkstreamNameTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v 
		(id INT(10) NOT NULL AUTO_INCREMENT, 
		name VARCHAR(128) NOT NULL UNIQUE, 
		PRIMARY KEY (id)) ENGINE=InnoDB;`, workstreamNameTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()
	if seed {
		seedWorkstreamNameTable()
	}
}

func buildSprintNameTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v(
		id INT(11) NOT NULL AUTO_INCREMENT,
		name VARCHAR(128) NOT NULL UNIQUE,
		start_date DATE NOT NULL,
		end_date DATE NOT NULL,
		PRIMARY KEY(id)
		) ENGINE=InnoDB;`,
		sprintNameTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedSprintNameTable()
	}
}

func buildWorkstreamSprintNameSprintSummaryMapTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (workstreamId INTEGER NOT NULL, sprintNameId INTEGER NOT NULL, sprintSummaryId INTEGER NOT NULL, PRIMARY KEY (workstreamId, sprintNameId, sprintSummaryId)) ENGINE=InnoDB;",
		workstreamSprintNameSprintSummaryMapTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedWorkstreamSprintNameSprintSummaryMapTable(db)
	}
}

func buildEngineerDetailsTable(db *sql.DB, seed bool) {
	queryString := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v 
		(id INT(10) NOT NULL AUTO_INCREMENT, 
		first_name TEXT, 
		last_name TEXT, 
		email VARCHAR(128) NOT NULL UNIQUE, 
		velocity INTEGER, 
		PRIMARY KEY(id)) ENGINE=InnoDB;`,
		engineerDetailsTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()

	if seed {
		seedEngineerDetailsTable()
	}
}

func buildSprintLineItemTable(db *sql.DB) {
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
			"PRIMARY KEY (id)) ENGINE=InnoDB;",
		sprintLineItemTable)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()
}

func buildWorkstreamSprintEngineerSprintLineItemMap(db *sql.DB) {
	queryString := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v 
		(workstream_id INT NOT NULL,
		sprint_id INT NOT NULL,
		engineer_id INT NOT NULL,
		sprint_line_item_id INT NOT NULL,
		PRIMARY KEY (workstream_id, sprint_id, engineer_id)) ENGINE=InnoDB;`,
		workstreamSprintEngineerSprintLineItemMapTable,
	)

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec()
}

func seedWorkstreamNameTable() {
	AddWorkstreamName("Workstream AAA000")
	AddWorkstreamName("Workstream AAA001")
	AddWorkstreamName("Workstream AAA002")
}

func seedSprintNameTable() {
	AddSprintName("2019.06.20", "2019-06-20", "2019-07-03")
	AddSprintName("2019.07.04", "2019-07-04", "2019-07-17")
	AddSprintName("2019.07.18", "2019-07-18", "2019-07-31")
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

func seedEngineerDetailsTable() {
	AddEngineerDetails("Bruce", "Dickinson", "a@mail.com")
	AddEngineerDetails("Steve", "Harris", "b@mail.com")
	AddEngineerDetails("Nicko", "McBrain", "c@mail.com")
	AddEngineerDetails("Adrian", "Smith", "d@mail.com")
	AddEngineerDetails("Dave", "Murray", "e@email.com")
	AddEngineerDetails("Janick", "Gers", "f@mail.com")
	AddEngineerDetails("Paul", "Di`Anno", "g@mail.com")
	AddEngineerDetails("Blaze", "Bayley", "H@mail.com")
	AddEngineerDetails("Clive", "Burr", "i@mail.com")
	AddEngineerDetails("Dennis", "Stratton", "j@mail.com")
	AddEngineerDetails("Thunderstick", "Joe", "k@mail.com")
	AddEngineerDetails("Doug", "Sampson", "l@mail.com")
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
