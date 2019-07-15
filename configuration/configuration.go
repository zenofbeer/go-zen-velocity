package configuration

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

// Config global config settings
type Config struct {
	App AppSettings
}

// AppSettings settings for each page
type AppSettings struct {
	Title      string
	CSSPath    string
	JqueryPath string
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
