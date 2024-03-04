package app

import (
	"errors"
	"quiz/basic"
	"quiz/basic/colx"
	"quiz/basic/uuidx"
	"quiz/domain"
)

type SubmitQuestionCommand struct {
	AuthorId    string
	CreatorId   string
	CreatorName string
	Question    string
	Answers     [4]string
	AnswerIdx   int
}

func HandleSubmitQuestionCommand(
	cmd SubmitQuestionCommand,
	r *basic.Repository[*domain.QuestionSubmit]) error {

	err := validate(cmd)
	if err != nil {
		return err
	}

	s := domain.NewQuestionSubmit("xyz", cmd.CreatorId, cmd.Question, cmd.Answers)
	(*r).Add(&s)

	return nil
}

func validate(
	cmd SubmitQuestionCommand) error {

	if !uuidx.IsValidUUID(cmd.AuthorId) {
		return errors.New("authorId must be a valid uuid")
	}

	if len(cmd.CreatorName) == 0 && len(cmd.CreatorId) == 0 {
		return errors.New("creatorId or creatorName must be provided")
	}

	if len(cmd.CreatorName) == 0 && !uuidx.IsValidUUID(cmd.CreatorId) {
		return errors.New("creatorId must be a valid uuid")
	}

	if len(cmd.CreatorName) < 3 {
		return errors.New("creatorName must be at least 3 characters long")
	}

	if len(cmd.Question) == 0 {
		return errors.New("question must be provided")
	}

	if colx.HasEmptyStrings(cmd.Answers[:]) {
		return errors.New("All 4 answers must be provided")
	}

	if colx.IsInRange(cmd.Answers[:], cmd.AnswerIdx) {
		return errors.New("Valid answer index must be provided")
	}

	return nil
}
