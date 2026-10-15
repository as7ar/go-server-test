package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age, err := strconv.ParseUint(r.URL.Query().Get("age"), 10, 0)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	user := User{
		Name: name,
		Age:  uint(age),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		return
	}
}
