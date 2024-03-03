package httpx

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequest[T any](apiRequest *T, w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		BadRequestError(w)
		return err
	}

	err = json.Unmarshal(body, apiRequest)
	if err != nil {
		BadRequestError(w)
		return err
	}

	return nil
}

func WriteResponse(response interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(response)
	if err != nil {
		BadRequestError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
