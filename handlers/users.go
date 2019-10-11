package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LorneCurrie/gorilla-mux-test/models/user"
)

// GetAllUsers returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	u := user.NewUser("John Doe").SetAge(38).Approve(true)
	fmt.Println(u)

	js, err := json.Marshal(*u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v\n", r.Method)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
