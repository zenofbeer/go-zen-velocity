package controllers

import "encoding/json"

// WorkstreamName the workstream Name and ID
type WorkstreamName struct {
	ID   int
	Name string
}

// WorkstreamNameList a list of WorkstreamName
type WorkstreamNameList struct {
	ListTitle       string
	WorkstreamNames []WorkstreamName
}

// GetWorkstreamNames return a json object containing
// workstream names and IDs
func GetWorkstreamNames() ([]byte, error) {
	data := WorkstreamNameList{
		ListTitle: "The list title",
		WorkstreamNames: []WorkstreamName{
			{
				ID:   -1,
				Name: "Select a workstream",
			},
			{
				ID:   0,
				Name: "Air Cancel",
			},
			{
				ID:   1,
				Name: "Air Schedule Change",
			},
			{
				ID:   2,
				Name: "Shopping",
			},
		},
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataJSON, nil
}
