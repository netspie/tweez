package app

import (
	"quiz/basic"
	"quiz/domain"
)

type SubmitQuestionCommand struct {
	UserId      string
	CreatorId   string
	CreatorName string
	Question    string
	Answers     [4]string
}

func HandleSubmitQuestionCommand(
	cmd SubmitQuestionCommand,
	r *basic.Repository[*domain.QuestionSubmit]) {

}
