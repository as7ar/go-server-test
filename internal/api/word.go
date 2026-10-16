package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/asxar/backend-test/cmd/api/assets"
	"github.com/asxar/backend-test/internal/model"
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
	startWith := r.URL.Query().Get("start-with")
	endWith := r.URL.Query().Get("end-with")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if len(r.URL.Query()) <= 0 {
		content := model.Response{
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
		if search != "" && !strings.Contains(word, search) {
			continue
		}
		if startWith != "" && !strings.HasPrefix(word, startWith) {
			continue
		}
		if endWith != "" && !strings.HasSuffix(word, endWith) {
			continue
		}
		if !checkLength(word, r) {
			continue
		}

		result = append(result, word)
	}

	content := model.Response{
		Code:    http.StatusOK,
		Content: result,
	}

	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return
	}
}

func checkLength(word string, r *http.Request) bool {
	query := r.URL.Query()

	length, _ := strconv.Atoi(query.Get("len"))
	minLen, _ := strconv.Atoi(query.Get("min-len"))
	maxLen, _ := strconv.Atoi(query.Get("max-len"))

	wordLen := utf8.RuneCountInString(word)

	return (!query.Has("len") || wordLen == length) &&
		(!query.Has("min-len") || wordLen >= minLen) &&
		(!query.Has("max-len") || wordLen <= maxLen)
}
