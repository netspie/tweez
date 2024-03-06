package submits

import "time"

type QuestionSubmit struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatorId string    `json:"creatorId"`
	Question  string    `json:"question"`
	Answers   [4]string `json:"answers"`
	AnswerIdx int       `json:"answer"`
	Confirmed bool      `json:"confirmed"`
}

func NewQuestionSubmit(
	id string,
	creatorId string,
	question string,
	answers [4]string,
	answerIdx int) QuestionSubmit {

	return QuestionSubmit{
		Id:        id,
		CreatedAt: time.Now(),
		CreatorId: creatorId,
		Question:  question,
		Answers:   answers,
		AnswerIdx: answerIdx,
		Confirmed: false,
	}
}

func (qs QuestionSubmit) GetId() string {
	return qs.Id
}
