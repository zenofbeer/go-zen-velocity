package config

// Config global config settings
type Config struct {
	Page PageSettings
}

// PageSettings settings for each page
type PageSettings struct {
	Title      string
	CSSPath    string
	JqueryPath string
}
