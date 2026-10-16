package user

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/asxar/backend-test/internal/model"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age, err := strconv.ParseUint(r.URL.Query().Get("age"), 10, 0)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		content := model.Response{
			Code:    http.StatusBadRequest,
			Content: "Invalid Age",
		}

		err := json.NewEncoder(w).Encode(content)
		if err != nil {
			return
		}
		return
	}

	response := model.Response{
		Code: 200,
		Content: User{
			Name: name,
			Age:  uint(age),
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}

func HandleIPRequest(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Failed to get IP", http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintln(w, ip)
}
