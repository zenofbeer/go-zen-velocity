package controllers

import "github.com/zenofbeer/go-zen-velocity/data"

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
	return data.GetWorkstreamNames()
}
