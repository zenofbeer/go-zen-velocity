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
