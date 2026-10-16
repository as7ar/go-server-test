package api

import (
	"encoding/json"
	"net/http"

	"github.com/asxar/backend-test/internal/model"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.NewEncoder(w).Encode(model.Response{Code: http.StatusOK, Content: "Hello, World!"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
