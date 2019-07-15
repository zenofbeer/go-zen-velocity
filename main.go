package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/zenofbeer/go-zen-velocity/data"

	"github.com/gorilla/mux"
	"github.com/zenofbeer/go-zen-velocity/configuration"
	"github.com/zenofbeer/go-zen-velocity/controllers"
)

// SiteTemplate contains the base site fields
type SiteTemplate struct {
	PageTitle  string
	CSSPath    string
	JqueryPath string
	PageID     string
	PageScript string
}

// WorkstreamViewModel loads the data for the workstream home page
type WorkstreamViewModel struct {
	SiteTemplate
	WorkstreamName string
	Overview       data.WorkstreamOverview
}

var config = configuration.GetConfig()

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/velocity", handler).Methods("GET")
	r.HandleFunc("/velocity/workstreamNames", getWorkstreamNameList).Methods("POST")
	r.HandleFunc("/velocity/workstreamHome/{id:[0-9]+}", getWorkstreamHome).Methods("GET")

	staticFileDirectory := http.Dir("./resources/")
	staticFileHandler := http.StripPrefix("/resources/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/resources/").Handler(staticFileHandler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(
		template.ParseFiles(
			"./resources/templates/layout.html",
			"./resources/templates/head.html",
			"./resources/templates/index.html"))

	data := SiteTemplate{
		PageTitle:  config.App.Title,
		CSSPath:    config.App.CSSPath,
		JqueryPath: config.App.JqueryPath,
		PageID:     "index",
		PageScript: "resources/scripts/index.js",
	}
	templates.ExecuteTemplate(w, "layout", data)
}

func getWorkstreamHome(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(
		template.ParseFiles(
			"./resources/templates/layout.html",
			"./resources/templates/head.html",
			"./resources/templates/workstream.html"))

	params := mux.Vars(r)
	stringID := params["id"]
	workstreamID, _ := strconv.Atoi(stringID)
	displayName := controllers.GetWorkstreamName(workstreamID)

	siteTemplate := SiteTemplate{
		PageTitle:  config.App.Title,
		CSSPath:    config.App.CSSPath,
		JqueryPath: config.App.JqueryPath,
		PageID:     "workstreamHome",
		PageScript: "resources/scripts/index.js",
	}

	data := WorkstreamViewModel{
		SiteTemplate:   siteTemplate,
		WorkstreamName: displayName,
		Overview:       controllers.GetWorkstreamOverview(workstreamID),
	}

	templates.ExecuteTemplate(w, "layout", data)
}

func getWorkstreamNameList(w http.ResponseWriter, r *http.Request) {
	ajaxpostdata := r.FormValue("ajaxpostdata")
	fmt.Println("Receive ajax post data string ", ajaxpostdata)

	response, err := controllers.GetWorkstreamNames()
	if err != nil {
		fmt.Println(err.Error)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(response)
}
