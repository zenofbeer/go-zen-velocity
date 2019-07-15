package views

// HeadViewModel model for the head fields in all pages
type HeadViewModel struct {
	PageTitle  string
	CSSPath    string
	JqueryPath string
	PageID     string
	PageScript string
}

// HomeViewModel view model for rendering the home page (index.html)
type HomeViewModel struct {
	Head HeadViewModel
}
