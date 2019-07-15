package configuration

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

// Config global config settings
type Config struct {
	App        AppSettings
	Home       HomeSettings
	Workstream WorkstreamSettings
}

// AppSettings settings for each page
type AppSettings struct {
	Title      string
	CSSPath    string
	JqueryPath string
}

// HomeSettings settings for the home page
type HomeSettings struct {
	AppSettings AppSettings
	PageScript  string
	PageID      string
}

// WorkstreamSettings settings for the workstream page
type WorkstreamSettings struct {
	AppSettings AppSettings
	PageScript  string
	PageID      string
}

// GetConfig returns the configuration
func GetConfig() Config {
	retVal := Config{}

	err := gonfig.GetConf("./configuration/configuration.json", &retVal)
	if err != nil {
		fmt.Println(err.Error())
	}
	return retVal
}
