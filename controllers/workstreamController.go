package controllers

import "github.com/zenofbeer/go-zen-velocity/data"

// GetWorkstreamName ...
func GetWorkstreamName(ID int) string {
	return data.GetWorkstreamName(ID)
}

// GetWorkstreamOverview ...
func GetWorkstreamOverview(ID int) data.WorkstreamOverview {
	return data.GetWorkstreamOverview(ID)
}
