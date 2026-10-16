package test

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

var words []string

func InitWords() {
	data, err := os.ReadFile("assets/words.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &words); err != nil {
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
