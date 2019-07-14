package main

import (
	"fmt"
	"html/template"
	"net/http"

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

var config = configuration.GetConfig()

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/velocity", handler).Methods("GET")
	r.HandleFunc("/velocity/workstreamNames", getWorkstreamNameList).Methods("POST")
	r.HandleFunc("/velocity/workstreamHome", getWorkstreamHome).Methods("GET")

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
		PageTitle:  config.Page.Title,
		CSSPath:    config.Page.CSSPath,
		JqueryPath: config.Page.JqueryPath,
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

	data := SiteTemplate{
		PageTitle:  config.Page.Title,
		CSSPath:    config.Page.CSSPath,
		JqueryPath: config.Page.JqueryPath,
		PageID:     "workstreamHome",
		PageScript: "resources/scripts/index.js",
	}

	fmt.Println(r.FormValue("displayName"))

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
