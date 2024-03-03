package domain

type QuestionSubmit struct {
	id string
}

func NewQuestionSubmit(id string) QuestionSubmit {
	return QuestionSubmit{
		id: id,
	}
}

func (qs *QuestionSubmit) Id() string {
	return qs.id
}
