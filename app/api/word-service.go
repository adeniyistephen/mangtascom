package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adeniyistephen/mangtascom/business"
	"github.com/pkg/errors"
)

// wordCountGroup is a group of endpoints for word counter
type wordCountGroup struct {
	wordCount business.NewWordCount
}

// WordCounter is the handler for the word counter API that calls or wordcount method.
func (wcg wordCountGroup) WordCounter(w http.ResponseWriter, r *http.Request) {
	var text business.Word

	// Decode the incoming json data
	decoder := json.NewDecoder(r.Body)
	// Decode the incoming json data into the struct
	if err := decoder.Decode(&text); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call the business layer to get the word count
	wc, err := wcg.wordCount.WordCounter(text)
	if err != nil {
		log.Panic(errors.Wrap(err, "word count failed"))
	}

	// Return the word count
	respondWithJSON(w, http.StatusOK, wc)
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
