package submits

import (
	"net/http"
	"quiz/basic/ddd"
	"quiz/basic/httpx"
	"strings"
)

func (handler QuestionSubmitRouteHandler) HandleGetSingle(w http.ResponseWriter, r *http.Request) {
	routeParams := strings.Split(r.RequestURI, "/")

	q := HandleGetSubmitQuestionQuery(GetSubmitQuestionQuery{
		QuestionSubmitId: routeParams[len(routeParams)-1],
	}, handler.Repo)

	httpx.WriteResponse(q, w)
}

type GetSubmitQuestionQuery struct {
	QuestionSubmitId string
}

type GetSubmitQuestionQueryOut struct {
	UserId      string
	CreatorId   string
	CreatorName string
	Question    string
	Answers     [4]string
}

func HandleGetSubmitQuestionQuery(
	cmd GetSubmitQuestionQuery,
	r *ddd.Repository[*QuestionSubmit]) *QuestionSubmit {

	item, _ := (*r).Get(cmd.QuestionSubmitId)
	return item
}
