package main

import (
	"encoding/json"
	"net/http"
	"quiz/basic/ddd"
	"quiz/basic/httpx"
	"quiz/features/submits"
	"strconv"
)

const apiVersion = 1

func (api *api) route() *http.ServeMux {
	mux := http.NewServeMux()

	prefix := "/api/v" + strconv.Itoa(apiVersion)

	mux.HandleFunc(prefix+"/healthcheck", api.healthCheck)

	submitsRepo := ddd.NewMemoRepository[*submits.QuestionSubmit](nil)

	httpx.HandleRequests(mux, prefix+"/questionSubmits", submits.NewQuestionSubmitRouteHandler(&submitsRepo))

	return mux
}

func (api *api) healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": api.config.env,
		"version":     version,
	}

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
