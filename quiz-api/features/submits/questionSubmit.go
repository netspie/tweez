package submits

import "time"

type QuestionSubmit struct {
	Idv       string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatorId string    `json:"creatorId"`
	Question  string    `json:"question"`
	Answers   [4]string `json:"answers"`
	AnswerIdx int       `json:"answer"`
}

func NewQuestionSubmit(
	id string,
	creatorId string,
	question string,
	answers [4]string,
	answerIdx int) QuestionSubmit {

	return QuestionSubmit{
		Idv:       id,
		CreatedAt: time.Now(),
		CreatorId: creatorId,
		Question:  question,
		Answers:   answers,
		AnswerIdx: answerIdx,
	}
}

func (qs QuestionSubmit) Id() string {
	return qs.Idv
}
