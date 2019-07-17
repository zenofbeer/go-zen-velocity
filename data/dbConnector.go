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

// EngineerDetails contains the data from an engineer detail query result
type EngineerDetails struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

const workstreamNameTable string = "workstream_name"
const sprintNameTable string = "sprint_name"
const sprintSummaryTable string = "sprint_summary"
const workstreamSprintNameSprintSummaryMapTable string = "workstream_sprintName_sprintSummary_Map"
const engineerDetailsTable string = "engineer_details"

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

func getWorkstreamOverview(ID int) []SprintSummary {
	db := getDatabase()

	countQuery := fmt.Sprintf("SELECT COUNT(*) as count FROM %v WHERE workstreamId=%v", workstreamSprintNameSprintSummaryMapTable, ID)
	count, _ := db.Query(countQuery)
	queryString := fmt.Sprintf(
		"SELECT %v.name, %v.id, %v.workingDays, %v.pointsCommitted, %v.pointsAchieved "+
			"FROM %v "+
			"INNER JOIN %v "+
			"ON %v.workstreamId = %v "+
			"INNER JOIN %v "+
			"ON %v.id = %v.sprintSummaryId "+
			"AND %v.sprintNameId = %v.id "+
			"AND %v.sprintSummaryId = %v.id",
		sprintNameTable, sprintNameTable, sprintSummaryTable, sprintSummaryTable, sprintSummaryTable,
		sprintNameTable, workstreamSprintNameSprintSummaryMapTable,
		workstreamSprintNameSprintSummaryMapTable, ID,
		sprintSummaryTable,
		sprintSummaryTable, workstreamSprintNameSprintSummaryMapTable,

		workstreamSprintNameSprintSummaryMapTable, sprintNameTable,
		workstreamSprintNameSprintSummaryMapTable, sprintSummaryTable,
	)

	results, err := db.Query(queryString)
	checkError(err)
	resultCount := checkCount(count)
	summaries := make([]SprintSummary, resultCount)
	var sprintName string
	var sprintID int
	var workingDays int
	var pointsCommitted int
	var pointsAchieved int
	counter := 0

	for results.Next() {
		results.Scan(&sprintName, &sprintID, &workingDays, &pointsCommitted, &pointsAchieved)
		summaries[counter] = SprintSummary{
			Name:            sprintName,
			WorkstreamID:    ID,
			SprintID:        sprintID,
			WorkingDays:     workingDays,
			PointsCommitted: pointsCommitted,
			PointsAchieved:  pointsAchieved,
		}

		counter++
	}
	return summaries
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
	fmt.Println(connStr)
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
