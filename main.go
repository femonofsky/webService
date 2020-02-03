package main

import (
	"github.com/pluarlsight/webservice/controllers"
	"net/http"
)

func main() {
	const PORT = ":3000"
	// Attach All Controller
	controllers.RegisterControllers()

	// listens on the TCP network address addr
	http.ListenAndServe(PORT, nil)
}
