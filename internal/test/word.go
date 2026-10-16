package test

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/asxar/backend-test/cmd/api/assets"
)

var words []string

func InitWords() {
	if err := json.Unmarshal(assets.WordsJSON, &words); err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d words", len(words))
}

func HandleWord(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	if search == "" {
		http.Error(w, "Missing search parameter", http.StatusBadRequest)
		return
	}

	result := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, search) {
			result = append(result, word)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
