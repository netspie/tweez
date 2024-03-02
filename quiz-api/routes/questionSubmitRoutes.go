package routes

import (
	"fmt"
	"net/http"
)

type QuestionSubmitRouteHandler struct{}

func (handler QuestionSubmitRouteHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlePost Question Submit")
}

func (handler QuestionSubmitRouteHandler) HandleGetSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleGetSingle Question Submit")
}

func (handler QuestionSubmitRouteHandler) HandleGetMany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleGetMany Question Submit")
}

func (handler QuestionSubmitRouteHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlePut Question Submit")
}

func (handler QuestionSubmitRouteHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleDelete Question Submit")
}
