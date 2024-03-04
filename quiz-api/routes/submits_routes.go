package routes

import (
	"fmt"
	"net/http"
	"quiz/app"
	"quiz/basic"
	"quiz/basic/httpx"
	"quiz/domain"
	"strings"
)

type QuestionSubmitRouteHandler struct {
	Repo *basic.Repository[*domain.QuestionSubmit]
}

func NewQuestionSubmitRouteHandler(repo basic.Repository[*domain.QuestionSubmit]) QuestionSubmitRouteHandler {
	return QuestionSubmitRouteHandler{
		Repo: &repo,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SubmitQuestionApiRequest struct {
	CreatorId   string    `json:"creatorId"`
	CreatorName string    `json:"creatorName"`
	Question    string    `json:"question"`
	Answers     [4]string `json:"answers"`
}

func (handler QuestionSubmitRouteHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	req := &SubmitQuestionApiRequest{}
	if httpx.ReadRequest(req, w, r) != nil {
		return
	}

	// TODO: validate request
	err := app.HandleSubmitQuestionCommand(
		app.SubmitQuestionCommand{
			AuthorId:    "",
			CreatorId:   req.CreatorId,
			CreatorName: req.CreatorName,
			Question:    req.Question,
			Answers:     req.Answers,
		}, handler.Repo)

	if err != nil {
		httpx.WriteResponse(ErrorResponse{Error: err.Error()}, w)
	}
}

func (handler QuestionSubmitRouteHandler) HandleGetSingle(w http.ResponseWriter, r *http.Request) {
	routeParams := strings.Split(r.RequestURI, "/")

	q := app.HandleGetSubmitQuestionQuery(app.GetSubmitQuestionQuery{
		QuestionSubmitId: routeParams[len(routeParams)-1],
	}, handler.Repo)

	httpx.WriteResponse(q, w)
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
