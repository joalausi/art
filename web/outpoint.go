package web

import (
	"art/processing"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func renderPage(w http.ResponseWriter, page *Page) {
	artTemplate, err := template.ParseFiles("web/templates/mainpage.html")
	if err != nil {
		log.Fatalf("Error parsing template: %s\n", err.Error())
		panic(err)
	}

	err = artTemplate.Execute(w, page)
	if err != nil {
		log.Printf("Error executing template: %s", err.Error())
		return
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request with path %s", r.URL.Path)

	artPage := &Page{}

	renderPage(w, artPage)

	w.WriteHeader(http.StatusOK)
}

func postDecoder(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got /decoder request with path %s", r.URL.Path)

	if r.Method != "POST" {
		getRoot(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		log.Printf("Error parsing form data: %s", err.Error())
		return
	}

	log.Printf("Received setting %s", r.FormValue("setting"))
	log.Printf("Received data:\n%s", r.FormValue("input_data"))

	inputData := strings.Split(r.FormValue("input_data"), "\n")
	setting := r.FormValue("setting")

	encode := false
	if setting == "encode" {
		encode = true
	}

	outputData := processing.ProcessData(&inputData, encode)

	log.Printf("Created output:\n%s", strings.Join(outputData, "\n"))

	artPage := &Page{
		Setting:    encode,
		InputData:  []byte(r.FormValue("input_data")),
		OutputData: []byte(strings.Join(outputData, "\n")),
	}

	renderPage(w, artPage)

	w.WriteHeader(http.StatusAccepted)
}
