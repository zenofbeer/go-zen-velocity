package data

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

func getWorkstreamOverviewOld(ID int) WorkstreamOverview {
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
