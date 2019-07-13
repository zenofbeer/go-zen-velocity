package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/zenofbeer/go-zen-velocity/config"

	"github.com/tkanos/gonfig"

	"github.com/gorilla/mux"
	"github.com/zenofbeer/go-zen-velocity/controllers"
)

// SiteTemplate contains the base site fields
type SiteTemplate struct {
	PageTitle  string
	CSSPath    string
	JqueryPath string
	PageScript string
}

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/velocity", handler).Methods("GET")
	r.HandleFunc("/velocity/workstreamNames", getWorkstreamNameList).Methods("POST")
	r.HandleFunc("/velocity/workstreamHome", getWorkstreamHome).Methods("POST")

	staticFileDirectory := http.Dir("./resources/")
	staticFileHandler := http.StripPrefix("/resources/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/resources/").Handler(staticFileHandler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	config := config.Config{}
	err := gonfig.GetConf("./config/config.json", &config)
	if err != nil {
		fmt.Println(err.Error)
	}

	tmpl := template.Must(template.ParseFiles("./resources/index.html"))
	data := SiteTemplate{
		PageTitle:  config.Page.Title,
		CSSPath:    config.Page.CSSPath,
		JqueryPath: config.Page.JqueryPath,
		PageScript: "resources/scripts/index.js",
	}
	/*
		data := TodoPageData{
			PageTitle: "Velocity Tracker",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
	*/
	tmpl.Execute(w, data)
}

func getWorkstreamHome(w http.ResponseWriter, r *http.Request) {
	postData := r.FormValue("displayName")
	city := r.FormValue("id")
	fmt.Println(postData)
	fmt.Println(city)
}

// should be a get, but using a post as a sample for getting postData
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
