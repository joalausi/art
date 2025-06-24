package web

import (
	"errors"
	"log"
	"net/http"
	"os"
)

type ArtServer struct {
}

func (server *ArtServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/decoder", postDecoder)
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	err := http.ListenAndServe(":22459", mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err.Error())
		os.Exit(1)
	}
}
