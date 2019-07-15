package controllers

import (
	"github.com/zenofbeer/go-zen-velocity/configuration"
	"github.com/zenofbeer/go-zen-velocity/data"
	"github.com/zenofbeer/go-zen-velocity/views"
)

var config = configuration.GetConfig()

// GetHomeViewModel get the view model for the home page
func GetHomeViewModel() views.HomeViewModel {
	home := views.HomeViewModel{
		Head:   getHeadViewModel(),
		Foot:   getFootViewModel(config.Home.PageScript),
		PageID: config.Home.PageID,
	}
	return home
}

// GetWorkstreamViewModel get the viewmodel for the workstream page
func GetWorkstreamViewModel(ID int) views.WorkstreamViewModel {
	retVal := views.WorkstreamViewModel{
		Head:        getHeadViewModel(),
		Foot:        getFootViewModel(config.Workstream.PageScript),
		PageID:      config.Workstream.PageID,
		DisplayName: data.GetWorkstreamName(ID),
		Overview:    data.GetWorkstreamOverview(ID),
	}
	return retVal
}

// GetworkstreamNames return a json object containing workstream
// names and IDs
func GetworkstreamNames() ([]byte, error) {
	return data.GetWorkstreamNames()
}

func getHeadViewModel() views.HeadViewModel {
	retVal := views.HeadViewModel{
		PageTitle:  config.App.Title,
		CSSPath:    config.App.CSSPath,
		JqueryPath: config.App.JqueryPath,
	}
	return retVal
}

func getFootViewModel(pageScript string) views.FootViewModel {
	retVal := views.FootViewModel{
		PageScript: pageScript,
	}
	return retVal
}
