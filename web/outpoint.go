package web

import (
	"art/processing"
	"html/template"
	"log"
	"net/http"
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

	w.WriteHeader(http.StatusOK)
	renderPage(w, artPage)
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

	input := r.FormValue("input_data")
	setting := r.FormValue("setting")

	encode := false
	if setting == "encode" {
		encode = true
	}

	var (
		output    string
		decodeErr error
	)

	if encode {
		output = processing.EncodeMultiLine(input)
	} else {
		output, decodeErr = processing.DecodeMultiLine(input)
	}

	if decodeErr != nil {
		http.Error(w, "Invalid encoded string", http.StatusBadRequest)
		return
	}

	log.Printf("Created output:\n%s", output)

	artPage := &Page{
		Setting:    encode,
		InputData:  []byte(input),
		OutputData: []byte(output),
	}

	w.WriteHeader(http.StatusAccepted)
	renderPage(w, artPage)
}
