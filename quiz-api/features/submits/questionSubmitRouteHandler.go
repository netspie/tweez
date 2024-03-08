package submits

import (
	"fmt"
	"net/http"
	"quiz/basic/ddd"
)

type QuestionSubmitRouteHandler struct {
	Repo *ddd.Repository[*QuestionSubmit]
}

func NewQuestionSubmitRouteHandler(repo ddd.Repository[*QuestionSubmit]) QuestionSubmitRouteHandler {
	return QuestionSubmitRouteHandler{
		Repo: &repo,
	}
}

func (handler QuestionSubmitRouteHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleDelete Question Submit")
}
