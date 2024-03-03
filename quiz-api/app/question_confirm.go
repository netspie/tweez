package features

import (
	"quiz/basic"
	"quiz/domain"
)

type ConfirmQuestionCommand struct {
	UserId      string
	CreatorId   string
	CreatorName string
	Question    string
	Answers     [4]string
}

func HandleConfirmQuestionCommand(
	cmd ConfirmQuestionCommand,
	r *basic.Repository[*domain.QuestionSubmit]) {

}
