package shared

type QuestionTypeEnum string

const (
	QuestionTypeEnumName     QuestionTypeEnum = "name"
	QuestionTypeEnumText     QuestionTypeEnum = "text"
	QuestionTypeEnumEmail    QuestionTypeEnum = "email"
	QuestionTypeEnumPhone    QuestionTypeEnum = "phone"
	QuestionTypeEnumTextarea QuestionTypeEnum = "textarea"
	QuestionTypeEnumSelect   QuestionTypeEnum = "select"
	QuestionTypeEnumRadios   QuestionTypeEnum = "radios"
)

// Question
// Routing form questions.
type Question struct {
	AnswerChoices []string         `json:"answer_choices"`
	Name          string           `json:"name"`
	Required      bool             `json:"required"`
	Type          QuestionTypeEnum `json:"type"`
	UUID          string           `json:"uuid"`
}
