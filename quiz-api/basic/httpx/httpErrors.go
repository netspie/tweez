package httpx

import "net/http"

type ErrorResponse struct {
	Error string `json:error`
}

func BadRequestError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
