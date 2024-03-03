package app

import (
	"quiz/basic"
	"quiz/domain"
)

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
	r *basic.Repository[*domain.QuestionSubmit]) {

}
