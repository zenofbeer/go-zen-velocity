package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zenofbeer/go-zen-velocity/controllers"
)

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/velocity", handler).Methods("GET")
	r.HandleFunc("/velocity/workstreamNames", getWorkstreamNameList).Methods("POST")
	r.HandleFunc("/velocity/workstreamHome/{id:[0-9]+}", getWorkstreamHome).Methods("GET")
	r.HandleFunc("/velocity/workstreamHome/{id:[0-9]+}/sprint/{id:[0-9]+}", sprintDetailHandler).Methods("GET")

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
			"./resources/templates/foot.html",
			"./resources/templates/index.html"))

	data := controllers.GetHomeViewModel()
	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Println(err)
	}
}

func getWorkstreamHome(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(
		template.ParseFiles(
			"./resources/templates/layout.html",
			"./resources/templates/head.html",
			"./resources/templates/foot.html",
			"./resources/templates/workstream.html"))

	params := mux.Vars(r)
	stringID := params["id"]
	workstreamID, _ := strconv.Atoi(stringID)

	data := controllers.GetWorkstreamViewModel(workstreamID)
	templates.ExecuteTemplate(w, "layout", data)
}

func sprintDetailHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(
		template.ParseFiles(
			"./resources/templates/layout.html",
			"./resources/templates/head.html",
			"./resources/templates/foot.html",
			"./resources/templates/sprintDetail.html"))

	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getWorkstreamNameList(w http.ResponseWriter, r *http.Request) {
	ajaxpostdata := r.FormValue("ajaxpostdata")
	fmt.Println("Receive ajax post data string ", ajaxpostdata)

	response, err := controllers.GetworkstreamNames()
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(response)
}
