package configuration

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

// Config global config settings
type Config struct {
	App              AppSettings
	SprintSettings   Sprint
	ConnectionString string
	Home             HomeSettings
	Workstream       WorkstreamSettings
}

// AppSettings settings for each page
type AppSettings struct {
	Title      string
	CSSPath    string
	JqueryPath string
}

// Sprint sprint configuration settings
type Sprint struct {
	VelocityIncreaseGoalConstant int
	DefaultAvailability          int
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

	retVal.ConnectionString = os.Getenv("CONNECTION_STRING_ZENVELUSER")

	if err != nil {
		fmt.Println(err.Error())
	}
	return retVal
}
