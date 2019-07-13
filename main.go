package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zenofbeer/go-zen-velocity/controllers"
)

// Todo ...
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData ...
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// WorkstreamSelection ...
type WorkstreamSelection struct {
	ID   int
	Name string
}

// WorkstreamSelectionList ...
type WorkstreamSelectionList struct {
	ListTitle   string
	Workstreams []WorkstreamSelection
}

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/velocity", handler).Methods("GET")
	r.HandleFunc("/receive/workstreamNames", handleAjax).Methods("POST")

	staticFileDirectory := http.Dir("./resources/")
	staticFileHandler := http.StripPrefix("/resources/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/resources/").Handler(staticFileHandler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./resources/layout.html"))
	data := TodoPageData{
		PageTitle: "My ToDo list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func handleAjax(w http.ResponseWriter, r *http.Request) {
	ajaxpostdata := r.FormValue("ajaxpostdata")
	fmt.Println("Receive ajax post data string ", ajaxpostdata)

	response, err := controllers.GetWorkstreamNames()
	if err != nil {
		fmt.Println(err.Error)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(response)
}
