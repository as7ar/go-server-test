package test

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/asxar/backend-test/cmd/api/assets"
	"github.com/asxar/backend-test/internal/api"
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if search == "" {
		content := api.Response{
			Code:    http.StatusBadRequest,
			Content: "Missing search parameter",
		}

		err := json.NewEncoder(w).Encode(content)
		if err != nil {
			return
		}
		return
	}

	result := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, search) {
			result = append(result, word)
		}
	}

	content := api.Response{
		Code:    http.StatusOK,
		Content: result,
	}

	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return
	}
}
