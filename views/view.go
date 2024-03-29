package views

import (
	"github.com/zenofbeer/go-zen-velocity/data"
)

// HeadViewModel model for the head fields in all pages
type HeadViewModel struct {
	PageTitle  string
	CSSPath    string
	JqueryPath string
}

// FootViewModel common footer fields
type FootViewModel struct {
	PageScript string
}

// HomeViewModel view model for rendering the home page (index.html)
type HomeViewModel struct {
	Head       HeadViewModel
	Foot       FootViewModel
	PageID     string
	PageScript string
}

// WorkstreamViewModel viewmodel for rendering the workstream page
type WorkstreamViewModel struct {
	Head        HeadViewModel
	Foot        FootViewModel
	PageID      string
	PageScript  string
	DisplayName string
	Overview    data.WorkstreamOverview
}

// SprintDetailViewModel represents a sprint detail view
type SprintDetailViewModel struct {
	Head       HeadViewModel
	Foot       FootViewModel
	PageID     string
	PageScript string
	Detail     data.SprintDetail
}

// WorkstreamName Model for displaying workstream names in dropdown lists
type WorkstreamName struct {
	ID   int
	Name string
}

// WorkstreamNameList Model for displaying workstream names list in a dropdown list
type WorkstreamNameList struct {
	ListTitle       string
	WorkstreamNames []WorkstreamName
}
