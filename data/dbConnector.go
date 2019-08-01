package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zenofbeer/go-zen-velocity/configuration"
	// need to force import
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
)

var config = configuration.GetConfig()

// SprintSummary returns the sprint activity and performance
// status for a sprint
type SprintSummary struct {
	Name                     string
	WorkstreamID             int
	SprintID                 int
	WorkingDays              int
	PointsCommitted          int
	PointsAchieved           int
	TargetPercentageAchieved float64
	Productivity             float64
	ProductivityChange       float64
}

// SprintName contains a sprint name entity
type SprintName struct {
	ID   int
	Name string
}

// EngineerDetails contains the data from an engineer detail query result
type EngineerDetails struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

// SprintLineItem represents an engineer level line item in a sprint
type SprintLineItem struct {
	id                        int
	CurrentAvailability       int
	PreviousAvailability      int
	Capacity                  int
	TargetPoints              float64
	CommittedPointsThisSprint int
	CompletedPointsThisSprint int
	CompletedPointsLastSprint int
}

const workstreamNameTable string = "workstream_name"
const sprintNameTable string = "sprint_name"
const sprintSummaryTable string = "sprint_summary"
const workstreamSprintNameSprintSummaryMapTable string = "workstream_sprintname_sprintsummary_map"
const engineerDetailsTable string = "engineer_details"
const sprintLineItemTable string = "sprint_line_item"
const workstreamSprintEngineerSprintLineItemMapTable = "workstream_sprint_engineer_sprint_line_item_map"

func addWorkstream(name string) {
	db := getDatabase()
	queryString := fmt.Sprintf(
		`INSERT INTO %v
		(name) VALUES (?)`,
		workstreamNameTable,
	)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec(name)
	db.Close()
}

func getAllWorkstreamNames() []byte {
	dbBuilder(true)
	db := getDatabase()
	rowCountQuery := fmt.Sprintf("SELECT COUNT(*) as count FROM %v", workstreamNameTable)
	rowCount, _ := db.Query(rowCountQuery)
	nameCount := checkCount(rowCount)
	query := fmt.Sprintf("SELECT * FROM %v", workstreamNameTable)
	rows, _ := db.Query(query)
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

func getWorkstreamNameByID(ID int) string {
	db := getDatabase()
	row, _ := db.Query("SELECT name FROM " + workstreamNameTable + " WHERE id=" + strconv.Itoa(ID) + " ORDER BY id DESC LIMIT 1")
	var name string
	for row.Next() {
		row.Scan(&name)
	}
	row.Close()
	db.Close()
	return name
}

func addEngineerDetails(firstName string, lastName string, emailAddress string) {
	db := getDatabase()
	queryString := fmt.Sprintf(
		`INSERT INTO %v 
		(first_name,
			last_name,
			email,
			velocity) VALUES(?, ?, ?, ?)`,
		engineerDetailsTable)
	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec(firstName, lastName, emailAddress, 0)
	db.Close()
}

func getEngineerDetails(engineerID int) EngineerDetails {
	db := getDatabase()

	queryString := fmt.Sprintf(
		"SELECT * FROM %v WHERE id = %v LIMIT 1",
		engineerDetailsTable, engineerID)

	query, err := db.Query(queryString)
	checkError(err)

	var id int
	var firstName string
	var lastName string
	var emailAddress string
	var retVal EngineerDetails

	for query.Next() {
		query.Scan(&id, &firstName, &lastName, &emailAddress)
		retVal = EngineerDetails{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     emailAddress,
		}
	}

	return retVal
}

func addSprintName(name string) {
	queryString := fmt.Sprintf(`
		INSERT INTO %v (name) 
		VALUES (?)`,
		sprintNameTable)

	db := getDatabase()

	query, err := db.Prepare(queryString)
	checkError(err)
	query.Exec(name)

	db.Close()
}

func addSprintLineItem(lineItem SprintLineItem, workstreamID int, sprintID int, engineerID int) {
	db := getDatabase()
	tx, err := db.Begin()
	checkError(err)

	defer db.Close()

	lineItemID := insertSprintLineItem(lineItem, tx)
	addWorkstreamSprintEngineerSprintLineItemMap(workstreamID, sprintID, engineerID, lineItemID, tx)

	err = tx.Commit()
	db.Close()
}

func insertSprintLineItem(lineItem SprintLineItem, tx *sql.Tx) int {
	queryString := fmt.Sprintf(
		`INSERT INTO %v (
			current_availability,
			previous_availability,
			capacity,
			target_points,
			committed_points_this_sprint,
			completed_points_this_sprint,
			completed_points_last_sprint)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
		sprintLineItemTable)

	//query, err := tx.Prepare(queryString)

	result, err := tx.Exec(queryString,
		lineItem.CurrentAvailability,
		lineItem.PreviousAvailability,
		lineItem.Capacity,
		lineItem.TargetPoints,
		lineItem.CommittedPointsThisSprint,
		lineItem.CompletedPointsThisSprint,
		lineItem.CompletedPointsLastSprint,
	)

	if err != nil {
		tx.Rollback()
		checkError(err)
	}

	ID, err := result.LastInsertId()

	return int(ID)
}

func getSprintLineItem(sprintNameID int, engineerID int) SprintLineItem {
	queryString := fmt.Sprintf(`
		SELECT
			current_availability,
			previous_availability,
			capacity,
			target_points,
			committed_points_this_sprint,
			completed_points_this_sprint,
			completed_points_last_sprint
		FROM %v
		INNDR JOIN %v
		ON %v.engineer_id=%v
		AND %v.sprint_id=%v
		AND %v.sprint_line_item_id=id`,
		sprintLineItemTable,
		workstreamSprintEngineerSprintLineItemMapTable,
		workstreamSprintEngineerSprintLineItemMapTable, engineerID,
		workstreamSprintEngineerSprintLineItemMapTable, sprintNameID,
		workstreamSprintEngineerSprintLineItemMapTable)
	db := getDatabase()
	result, err := db.Query(queryString)
	checkError(err)
	var sprintLineItem SprintLineItem
	var currentAvailability int
	var previousAvailability int
	var capacity int
	var targetPoints float64
	var committedPointsThisSprint int
	var completedPointsThisSprint int
	var completedPointsLastSprint int

	for result.Next() {
		result.Scan(
			&currentAvailability, &previousAvailability, &capacity, &targetPoints,
			&committedPointsThisSprint, &completedPointsThisSprint, &completedPointsLastSprint)
		sprintLineItem.CurrentAvailability = currentAvailability
		sprintLineItem.PreviousAvailability = previousAvailability
		sprintLineItem.Capacity = capacity
		sprintLineItem.TargetPoints = targetPoints
		sprintLineItem.CommittedPointsThisSprint = committedPointsThisSprint
		sprintLineItem.CompletedPointsThisSprint = completedPointsThisSprint
		sprintLineItem.CompletedPointsLastSprint = completedPointsLastSprint
	}
	db.Close()
	return sprintLineItem
}

func getPreviousSprintName(ID int) SprintName {
	queryString := fmt.Sprintf(
		`SELECT * FROM %v 
		WHERE id = (
			SELECT MAX(id)
			FROM %v 
			WHERE id < %v)`,
		sprintNameTable,
		sprintNameTable,
		ID)
	db := getDatabase()
	result, err := db.Query(queryString)
	checkError(err)
	var sprintName SprintName
	previousID := -1
	previousName := ""

	sprintName = SprintName{
		ID:   previousID,
		Name: "",
	}

	for result.Next() {
		result.Scan(&previousID, &previousName)
		sprintName.ID = previousID
		sprintName.Name = previousName
	}

	db.Close()

	return sprintName

}

func getWorkstreamOverview(ID int) []SprintSummary {
	sprintNames := getSprintNameByWorkstreamID(ID)
	sprintSummaries := make([]SprintSummary, len(sprintNames))

	for i := range sprintNames {
		sprintSummaries[i] = SprintSummary{
			Name:                     sprintNames[i],
			WorkstreamID:             ID,
			SprintID:                 i,
			WorkingDays:              20,
			PointsCommitted:          7 + i,
			PointsAchieved:           9 + i,
			TargetPercentageAchieved: 5,
			Productivity:             70,
			ProductivityChange:       2,
		}
	}
	return sprintSummaries
}

func getSprintNameByWorkstreamID(ID int) []string {
	db := getDatabase()
	defer db.Close()
	queryString := fmt.Sprintf(`
	SELECT name
	FROM %v
	INNER JOIN %v
	ON workstream_id=%v
	AND %v.sprint_id=%v.id
	GROUP BY %v.name
	ORDER BY %v.id`,
		sprintNameTable,
		workstreamSprintEngineerSprintLineItemMapTable,
		ID,
		workstreamSprintEngineerSprintLineItemMapTable,
		sprintNameTable,
		sprintNameTable,
		sprintNameTable)
	result, _ := db.Query(queryString)
	var names []string
	for result.Next() {
		var name string
		err := result.Scan(&name)
		checkError(err)
		names = append(names, name)
	}
	return names
}

func addWorkstreamSprintEngineerSprintLineItemMap(workstreamID int, sprintID int, engineerID int, sprintLineItemID int, tx *sql.Tx) {
	queryString := fmt.Sprintf(
		`INSERT INTO %v (
			workstream_id,
			sprint_id,
			engineer_id,
			sprint_line_item_id)
			VALUES(?, ?, ?, ?)`,
		workstreamSprintEngineerSprintLineItemMapTable)

	result, err := tx.Exec(queryString, workstreamID, sprintID, engineerID, sprintLineItemID)
	if err != nil {
		tx.Rollback()
		println(result)
		println("here I am.")
	}
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
	//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
	connStr := config.ConnectionString
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
	}

	return db

	/* db, err := sql.Open("sqlite3", "vt.db")
	if err != nil {
		fmt.Println(err)
	}
	return db */
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
