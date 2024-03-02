package main

import (
	"fmt"
	"net/http"
	"quiz/basic"
	"quiz/routes"
	"strconv"
)

const apiVersion = 1

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()

	prefix := "/api/v" + strconv.Itoa(apiVersion)

	mux.HandleFunc(prefix+"/healthcheck", app.healthCheck)

	basic.HandleRequests(mux, prefix+"/questionSubmits", routes.QuestionSubmitRouteHandler{})

	return mux
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}
