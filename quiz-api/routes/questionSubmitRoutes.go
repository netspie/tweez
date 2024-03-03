package routes

import (
	"fmt"
	"net/http"
	"quiz/app"
	"quiz/basic"
	"quiz/domain"
)

type QuestionSubmitRouteHandler struct {
	repo *basic.Repository[*domain.QuestionSubmit]
}

func (handler QuestionSubmitRouteHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlePost Question Submit")
	app.HandleSubmitQuestionCommand(app.SubmitQuestionCommand{}, handler.repo)
	app.HandleConfirmQuestionCommand(app.ConfirmQuestionCommand{}, handler.repo)
}

func (handler QuestionSubmitRouteHandler) HandleGetSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandleGetSingle Question Submit")
	app.HandleGetSubmitQuestionQuery(app.GetSubmitQuestionQuery{}, handler.repo)
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
