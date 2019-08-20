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

// SprintDetail the detail view of the sprint
type SprintDetail struct {
	ID                                   int
	Name                                 string
	StartDate                            string
	EndDate                              string
	EngineerHeader                       string
	AvailabilityThisSprintHeader         string
	AvailabilityLastSprintHeader         string
	CapacityByAvailabilityHeader         string
	TargetPointsHeader                   string
	CommittedPointsThisSprintHeader      string
	CommittedPointsLastSprintHeader      string
	CompletedPointsLastSprintHeader      string
	RunningVelocityHeader                string
	LastSprintVelocityHeader             string
	TotalDaysAvailableHeader             string
	TotalDaysAvailableLastSprintHeader   string
	TotalCapacityByAvailabilityHeader    string
	TotalTargetPointsHeader              string
	TotalCommittedPointsHeader           string
	TotalCompletedPointsHeader           string
	TotalPointsCompletedLastSprintHeader string
	TotalsLineTitle                      string
	SprintLineItems                      []SprintLineItem
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

// GetSprintDetail get the sprint details
func GetSprintDetail(workstreamID int, sprintID int) SprintDetail {
	retVal := getSprintDetail(workstreamID, sprintID)
	retVal.EngineerHeader = "Engineer"
	retVal.AvailabilityThisSprintHeader = "Availability This Sprint"
	retVal.AvailabilityLastSprintHeader = "Availability Last Sprint"
	retVal.CapacityByAvailabilityHeader = "Capacity By Availability"
	retVal.TargetPointsHeader = "Target Points"
	retVal.CommittedPointsThisSprintHeader = "Committed Points This Sprint"
	retVal.CommittedPointsLastSprintHeader = "Committed Points Last Sprint"
	retVal.CompletedPointsLastSprintHeader = "Completed Points Last Sprint"
	retVal.RunningVelocityHeader = "Running Velocity"
	retVal.LastSprintVelocityHeader = "Last Sprint Velocity"
	retVal.TotalDaysAvailableHeader = "Total Days Available"
	retVal.TotalDaysAvailableLastSprintHeader =
		"Total Days Available Last Sprint"
	retVal.TotalCapacityByAvailabilityHeader =
		"Total Capacity By Availability"
	retVal.TotalTargetPointsHeader = "Total Target Points"
	retVal.TotalCommittedPointsHeader = "Committed Points"
	retVal.TotalCompletedPointsHeader = "Completed Points"
	retVal.TotalPointsCompletedLastSprintHeader = "Points Completed Last Sprint"
	retVal.TotalsLineTitle = "Totals"
	return retVal
}

// AddEngineerDetails add a new engineer. Velocity defaults to 0
func AddEngineerDetails(firstName string, lastName string, emailAddress string) {
	addEngineerDetails(firstName, lastName, emailAddress)
}

// AddSprintName add a new sprint
func AddSprintName(name string, startDate string, endDate string) {
	addSprintName(name, startDate, endDate)
}

// AddSprint add a new sprint to a workstream. A sprint requires a mapping between a workstream,
// a sprintNameID, an engineer, and a sprint line item
func AddSprint(workstreamID int, currentSprintNameID int, engineerID int) {
	// get previous sprint, if any
	previousSprintName := getPreviousSprintName(currentSprintNameID)
	defaultAvailability := config.SprintSettings.DefaultAvailability

	if previousSprintName.ID == -1 {
		sprintLineItem := SprintLineItem{
			CurrentAvailability:       10,
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
			CurrentAvailability:       10,
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

// ToDo: verify that this is no longer used, and delete dead code.
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
