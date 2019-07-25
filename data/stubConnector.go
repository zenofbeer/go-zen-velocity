package data

import (
	"math"
)

// WorkstreamName the workstream name and ID
type WorkstreamName struct {
	ID   int
	Name string
}

// WorkstreamOverview overview of the workstream
type WorkstreamOverview struct {
	NameTitle                     string
	WorkingDaysTitle              string
	PointsCommittedTitle          string
	PointsAchievedTitle           string
	TargetPercentageAchievedTitle string
	ProductivityTitle             string
	ProductivityChangeTitle       string
	SprintSummaries               []SprintSummary
}

// WorkstreamNameList a list of WorkstreamName
type WorkstreamNameList struct {
	ListTitle       string
	WorkstreamNames []WorkstreamName
}

// AddWorkstreamName add a workstream
func AddWorkstreamName(name string) {
	addWorkstream(name)
}

// GetWorkstreamNames return a json object containing
// workstream names and IDs
func GetWorkstreamNames() ([]byte, error) {
	return getAllWorkstreamNames(), nil
}

// GetWorkstreamName get name by id
func GetWorkstreamName(ID int) string {
	return getWorkstreamNameByID(ID)
}

// GetWorkstreamOverview ...
func GetWorkstreamOverview(ID int) WorkstreamOverview {
	return getWorkstreamOverviewOld(ID)
}

// AddEngineerDetails add a new engineer. Velocity defaults to 0
func AddEngineerDetails(firstName string, lastName string, emailAddress string) {
	addEngineerDetails(firstName, lastName, emailAddress)
}

// AddSprintName add a new sprint
func AddSprintName(name string) {
	addSprintName(name)
}

// AddSprint add a new sprint to a workstream. A sprint requires a mapping between a workstream,
// a sprintNameID, an engineer, and a sprint line item
func AddSprint(workstreamID int, currentSprintNameID int, engineerID int) {
	// get previous sprint, if any
	previousSprintName := getPreviousSprintName(currentSprintNameID)

	if previousSprintName.ID == -1 {
		sprintLineItem := SprintLineItem{
			CurrentAvailability:       0,
			PreviousAvailability:      0,
			Capacity:                  0,
			TargetPoints:              0,
			CommittedPointsThisSprint: 0,
			CompletedPointsThisSprint: 0,
			CompletedPointsLastSprint: 0,
		}
		sprintLineItemID := addSprintLineItem(sprintLineItem)
		addWorkstreamSprintEngineerSprintLineItemMap(workstreamID, currentSprintNameID, engineerID, sprintLineItemID)
	}

	// get engineer
	//engineer := getEngineerDetails(engineerID)

	// get previous sprint line item by engineerID & previousSprintName.ID, if any

	// if previousSprintName != nil

	// calculate new line item fields
	// build SprintLineItem struct from calculated fields and engineer data
	// add sprint line item
	// add record in workstream_sprint_engineer_line_item_map
	// when this is working execute above concurrently
}

func getWorkstreamOverviewOld(ID int) WorkstreamOverview {
	summary := getWorkstreamOverview(ID)

	for i := 0; i < len(summary); i++ {
		pa := summary[i].PointsAchieved
		pc := summary[i].PointsCommitted
		wd := summary[i].WorkingDays
		tpa := getFloatToTwo((float64(pa) / float64(pc)) * 100)
		p := getFloatToTwo((float64(pa) / float64(wd)) * 100)

		summary[i].TargetPercentageAchieved = tpa
		summary[i].Productivity = p

		if i > 0 {
			tp := summary[i].Productivity
			lp := summary[i-1].Productivity
			pc := tp - lp
			summary[i].ProductivityChange = getFloatToTwo(pc)
		}
	}

	retVal := WorkstreamOverview{
		NameTitle:                     "Sprint",
		WorkingDaysTitle:              "Working Days",
		PointsCommittedTitle:          "Points Committed",
		PointsAchievedTitle:           "Points Achieved",
		TargetPercentageAchievedTitle: "Percentage of Target Achieved",
		ProductivityTitle:             "Productivity",
		ProductivityChangeTitle:       "Productivity Change",
		SprintSummaries:               summary,
	}
	return retVal
}

func getFloatToTwo(rawNumber float64) float64 {
	return math.Round(rawNumber*100) / 100
}
