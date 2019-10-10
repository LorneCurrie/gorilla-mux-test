package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handle the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r.Method)
	w.Header().Set("Made-up-header", "42")
}
