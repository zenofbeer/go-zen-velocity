package data

import (
	"github.com/zenofbeer/go-zen-velocity/helpers"
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
	summary := getWorkstreamOverview(ID)

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
	defaultAvailability := config.SprintSettings.DefaultAvailability

	if previousSprintName.ID == -1 {
		sprintLineItem := SprintLineItem{
			CurrentAvailability:       defaultAvailability,
			PreviousAvailability:      0,
			Capacity:                  0,
			TargetPoints:              0,
			CommittedPointsThisSprint: 0,
			CompletedPointsThisSprint: 0,
			CompletedPointsLastSprint: 0,
		}
		addSprintLineItem(
			sprintLineItem, workstreamID, currentSprintNameID, engineerID)
	} else {
		engineer := getEngineerDetails(engineerID)
		// get previous sprint line item by engineerID & previousSprintName.ID
		previousSprintLineItem := getSprintLineItem(previousSprintName.ID, engineer.ID)
		// calculate new line item fields
		// build SprintLineItem struct from calculated fields and engineer data
		currentSprintLineItem := SprintLineItem{
			CurrentAvailability:       defaultAvailability,
			PreviousAvailability:      previousSprintLineItem.CurrentAvailability,
			Capacity:                  helpers.CalculateCapacityAsPercentage(10, previousSprintLineItem.CurrentAvailability),
			TargetPoints:              helpers.CalculateTargetPoints(previousSprintLineItem.CompletedPointsThisSprint, defaultAvailability),
			CommittedPointsThisSprint: 0,
			CompletedPointsThisSprint: 0,
			CompletedPointsLastSprint: previousSprintLineItem.CompletedPointsThisSprint,
		}
		addSprintLineItem(
			currentSprintLineItem, workstreamID, currentSprintNameID, engineerID)
	}
}

func getWorkstreamOverviewOld(ID int) WorkstreamOverview {
	summary := getWorkstreamOverview(ID)

	for i := 0; i < len(summary); i++ {
		pa := summary[i].PointsAchieved
		pc := summary[i].PointsCommitted
		wd := summary[i].WorkingDays
		tpa := helpers.RoundToTwoDecimals((float64(pa) / float64(pc)) * 100)
		p := helpers.RoundToTwoDecimals((float64(pa) / float64(wd)) * 100)

		summary[i].TargetPercentageAchieved = tpa
		summary[i].Productivity = p

		if i > 0 {
			tp := summary[i].Productivity
			lp := summary[i-1].Productivity
			pc := tp - lp
			summary[i].ProductivityChange = helpers.RoundToTwoDecimals(pc)
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
