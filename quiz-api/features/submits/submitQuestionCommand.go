package submits

import (
	"errors"
	"net/http"
	"quiz/basic/ddd"
	"quiz/basic/httpx"
	"quiz/basic/slicex"
	"quiz/basic/uuidx"

	"github.com/google/uuid"
)

type SubmitQuestionApiRequest struct {
	CreatorId   string    `json:"creatorId"`
	CreatorName string    `json:"creatorName"`
	Question    string    `json:"question"`
	Answers     [4]string `json:"answers"`
	AnswerIdx   int       `json:"answerIdx"`
}

func (handler QuestionSubmitRouteHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	req := &SubmitQuestionApiRequest{}
	if httpx.ReadRequest(req, w, r) != nil {
		return
	}

	err := HandleSubmitQuestionCommand(
		SubmitQuestionCommand{
			AuthorId:    "",
			CreatorId:   req.CreatorId,
			CreatorName: req.CreatorName,
			Question:    req.Question,
			Answers:     req.Answers,
			AnswerIdx:   req.AnswerIdx,
		}, handler.Repo)

	if err != nil {
		httpx.WriteResponse(httpx.ErrorResponse{Error: err.Error()}, w)
	}
}

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
	r *ddd.Repository[*QuestionSubmit]) error {

	err := validateCommand(cmd)
	if err != nil {
		return err
	}

	s := NewQuestionSubmit(uuid.NewString(), cmd.CreatorId, cmd.Question, cmd.Answers, cmd.AnswerIdx)
	(*r).Add(&s)

	return nil
}

func validateCommand(
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

	if slicex.HasEmptyStrings(cmd.Answers[:]) {
		return errors.New("all 4 answers must be provided")
	}

	if slicex.IsInRange(cmd.Answers[:], cmd.AnswerIdx) {
		return errors.New("valid answer index must be provided")
	}

	return nil
}
