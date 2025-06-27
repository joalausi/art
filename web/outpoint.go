package web

import (
	"art/processing"
	"html/template"
	"io"
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
	if r.Method != http.MethodGet || r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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

	contentType := r.Header.Get("Content-Type")
	var input string
	var setting string
	if strings.HasPrefix(contentType, "text/plain") {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		input = string(body)
		setting = r.URL.Query().Get("setting")
	} else {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			log.Printf("Error parsing form data: %s", err.Error())
			return
		}

		input = r.FormValue("input_data")
		setting = r.FormValue("setting")
	}

	log.Printf("Received setting %s", setting)
	log.Printf("Received data:\n%s", input)

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
