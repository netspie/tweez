package submits

import (
	"errors"
	"net/http"
	"quiz/basic/ddd"
	"quiz/basic/httpx"
	"quiz/basic/slicex"
)

type UpdateQuestionSubmitApiRequest struct {
	Question  string    `json:"question"`
	Answers   [4]string `json:"answers"`
	AnswerIdx int       `json:"answerIdx"`
}

func (handler QuestionSubmitRouteHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	req := &UpdateQuestionSubmitApiRequest{}
	if httpx.ReadRequest(req, w, r) != nil {
		return
	}

	err := HandleUpdateQuestionSubmitCommand(
		UpdateQuestionSubmitCommand{
			Question:  req.Question,
			Answers:   req.Answers,
			AnswerIdx: req.AnswerIdx,
		}, handler.Repo)

	if err != nil {
		httpx.WriteResponse(httpx.ErrorResponse{Error: err.Error()}, w)
	}
}

type UpdateQuestionSubmitCommand struct {
	QuestionSubmitId string
	Question         string
	Answers          [4]string
	AnswerIdx        int
}

func HandleUpdateQuestionSubmitCommand(
	cmd UpdateQuestionSubmitCommand,
	r *ddd.Repository[*QuestionSubmit]) error {

	err := validateUpdateQuestionSubmitCommand(cmd)
	if err != nil {
		return err
	}

	s, res := (*r).Get(cmd.QuestionSubmitId)
	if !res {
		return errors.New("the question submit does not exist")
	}

	s.Question = cmd.Question
	s.Answers = cmd.Answers
	s.AnswerIdx = cmd.AnswerIdx

	(*r).Update(s)

	return nil
}

func validateUpdateQuestionSubmitCommand(
	cmd UpdateQuestionSubmitCommand) error {

	if len(cmd.Question) == 0 {
		return errors.New("question must be provided")
	}

	if slicex.HasEmptyStrings(cmd.Answers[:]) {
		return errors.New("all 4 answers must be provided")
	}

	if slicex.IsInRange(cmd.Answers[:], cmd.AnswerIdx) {
		return errors.New("valid answer index must be provided")
	}

	return nil
}
