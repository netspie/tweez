package main

import (
	"encoding/json"
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

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
