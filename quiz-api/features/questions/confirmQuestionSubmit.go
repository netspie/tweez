package questions

import (
	"quiz/basic/ddd"
	"quiz/features/submits"
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
	r *ddd.Repository[*submits.QuestionSubmit]) {

}
