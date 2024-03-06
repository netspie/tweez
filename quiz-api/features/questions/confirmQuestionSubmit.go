package questions

import (
	"errors"
	"quiz/basic/ddd"
	"quiz/basic/uuidx"
	"quiz/features/submits"

	"github.com/google/uuid"
)

type SubmitQuestionApiRequest struct {
	SubmitId string `json:"submitId"`
}

type ConfirmQuestionCommand struct {
	SubmitId string
}

func HandleConfirmQuestionCommand(
	cmd ConfirmQuestionCommand,
	submitRp *ddd.Repository[*submits.QuestionSubmit],
	questionRp *ddd.Repository[*Question]) error {

	err := validateCommand(cmd)
	if err != nil {
		return err
	}

	s, res := (*submitRp).Get(cmd.SubmitId)
	if !res {
		return errors.New("could not find given submit in the database")
	}

	q := NewQuestion(uuid.NewString(), s.CreatorId, s.Question, s.Answers, s.AnswerIdx)
	res = (*questionRp).Add(&q)
	if !res {
		return errors.New("could not add new question to the database")
	}

	(*s).Confirmed = true

	res = (*submitRp).Add(s)
	if !res {
		return errors.New("could not submit in the database")
	}

	return nil
}

func validateCommand(
	cmd ConfirmQuestionCommand) error {

	if !uuidx.IsValidUUID(cmd.SubmitId) {
		return errors.New("authorId must be a valid uuid")
	}

	return nil
}
