package app

import (
	"quiz/basic"
	"quiz/domain"
)

type SubmitQuestionCommand struct {
	AuthorId    string
	CreatorId   string
	CreatorName string
	Question    string
	Answers     [4]string
}

func HandleSubmitQuestionCommand(
	cmd SubmitQuestionCommand,
	r *basic.Repository[*domain.QuestionSubmit]) {

	s := domain.NewQuestionSubmit("xyz", cmd.CreatorId, cmd.Question, cmd.Answers)
	(*r).Add(&s)
}
