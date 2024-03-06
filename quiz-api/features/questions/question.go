package questions

import "time"

type Question struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatorId string    `json:"creatorId"`
	Question  string    `json:"question"`
	Answers   [4]string `json:"answers"`
	AnswerIdx int       `json:"answer"`
}

func NewQuestion(
	id string,
	creatorId string,
	question string,
	answers [4]string,
	answerIdx int) Question {

	return Question{
		Id:        id,
		CreatedAt: time.Now(),
		CreatorId: creatorId,
		Question:  question,
		Answers:   answers,
		AnswerIdx: answerIdx,
	}
}

func (qs Question) GetId() string {
	return qs.Id
}
