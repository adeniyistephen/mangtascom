package api

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/adeniyistephen/mangtascom/business"
)

func Handle(log *log.Logger) *mux.Router {
	// Create the router
	r := mux.NewRouter()

	// Register word counter routes
	wc := wordCountGroup {
		business.New(log),
	}

	r.HandleFunc("/wordcounter", wc.WordCounter).Methods("POST")

	return r
}