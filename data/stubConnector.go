package data

// WorkstreamName the workstream name and ID
type WorkstreamName struct {
	ID   int
	Name string
}

// SprintSummary returns the sprint activity and performance
// status for a sprint
type SprintSummary struct {
	Name                     string
	WorkingDays              int
	PointsCommitted          int
	PointsAchieved           int
	TargetPercentageAchieved float32
	Productivity             float32
	ProductivityChange       float32
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
	return getWorkstreamOverview(ID)
}

func getWorkstreamOverview(ID int) WorkstreamOverview {
	summary := []SprintSummary{
		SprintSummary{
			Name:                     "2019.06.20",
			WorkingDays:              34,
			PointsCommitted:          26,
			PointsAchieved:           13,
			TargetPercentageAchieved: 50.00,
			Productivity:             0.00,
			ProductivityChange:       0.00,
		},
		SprintSummary{
			Name:                     "2019.07.04",
			WorkingDays:              30,
			PointsCommitted:          22,
			PointsAchieved:           0,
			TargetPercentageAchieved: 0.00,
			Productivity:             0.00,
			ProductivityChange:       -38.24,
		},
		SprintSummary{
			Name:                     "2019.07.17",
			WorkingDays:              35,
			PointsCommitted:          27,
			PointsAchieved:           20,
			TargetPercentageAchieved: 52.00,
			Productivity:             1.00,
			ProductivityChange:       -37.00,
		},
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
