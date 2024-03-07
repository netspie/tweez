package submits

import (
	"net/http"
	"quiz/basic/ddd"
	"quiz/basic/httpx"
	"quiz/basic/slicex"
)

func (handler QuestionSubmitRouteHandler) HandleGetMany(w http.ResponseWriter, r *http.Request) {
	//routeParams := strings.Split(r.RequestURI, "/")
	q := HandleGetSubmitQuestionsQuery(GetSubmitQuestionsQuery{}, handler.Repo)
	httpx.WriteResponse(q, w)
}

type GetSubmitQuestionsQuery struct{}

type GetSubmitQuestionsQueryOut struct {
	QuestionIds []string `json:"questionIds"`
}

func HandleGetSubmitQuestionsQuery(
	cmd GetSubmitQuestionsQuery,
	r *ddd.Repository[*QuestionSubmit]) GetSubmitQuestionsQueryOut {

	items, _ := (*r).GetMany()
	return GetSubmitQuestionsQueryOut{QuestionIds: slicex.Map(items, func(qs *QuestionSubmit) string { return qs.Id })}
}
