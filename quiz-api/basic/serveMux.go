package basic

import (
	"net/http"
)

type RouteHandler interface {
	HandlePost(w http.ResponseWriter, r *http.Request)
	HandleGetSingle(w http.ResponseWriter, r *http.Request)
	HandleGetMany(w http.ResponseWriter, r *http.Request)
	HandlePut(w http.ResponseWriter, r *http.Request)
	HandleDelete(w http.ResponseWriter, r *http.Request)
}

func HandleRequests(mux *http.ServeMux, resourcePath string, handler RouteHandler) {
	mux.HandleFunc(
		resourcePath,
		func(w http.ResponseWriter, r *http.Request) { getOrCreateResourceHandler(w, r, handler) })

	mux.HandleFunc(
		resourcePath+"/",
		func(w http.ResponseWriter, r *http.Request) { getUpdateOrDeleteResourceHandler(w, r, handler) })
}

func getOrCreateResourceHandler(w http.ResponseWriter, r *http.Request, handler RouteHandler) {

	switch r.Method {
	case http.MethodGet:
		handler.HandleGetMany(w, r)
	case http.MethodPost:
		handler.HandlePost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func getUpdateOrDeleteResourceHandler(w http.ResponseWriter, r *http.Request, handler RouteHandler) {
	switch r.Method {
	case http.MethodGet:
		handler.HandleGetSingle(w, r)
	case http.MethodPut:
		handler.HandlePut(w, r)
	case http.MethodDelete:
		handler.HandleDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
