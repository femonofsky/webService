package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// method constructor
func newUserController() *userController {
	return &userController{
		regexp.MustCompile(`^/users/(\d+)/?`),
	}
}