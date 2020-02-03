package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	// Initializing User Controller
	uc := newUserController()

	// Adding User routes and Handlers
	http.Handle("/users",  *uc)
	http.Handle("/users/",  *uc)

}

// Convert Data to JSON Object
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}