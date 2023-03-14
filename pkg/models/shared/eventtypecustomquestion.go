package shared

type EventTypeCustomQuestionTypeEnum string

const (
	EventTypeCustomQuestionTypeEnumString       EventTypeCustomQuestionTypeEnum = "string"
	EventTypeCustomQuestionTypeEnumText         EventTypeCustomQuestionTypeEnum = "text"
	EventTypeCustomQuestionTypeEnumPhoneNumber  EventTypeCustomQuestionTypeEnum = "phone_number"
	EventTypeCustomQuestionTypeEnumSingleSelect EventTypeCustomQuestionTypeEnum = "single_select"
	EventTypeCustomQuestionTypeEnumMultiSelect  EventTypeCustomQuestionTypeEnum = "multi_select"
)

type EventTypeCustomQuestion struct {
	AnswerChoices []string                        `json:"answer_choices"`
	Enabled       bool                            `json:"enabled"`
	IncludeOther  bool                            `json:"include_other"`
	Name          string                          `json:"name"`
	Position      float64                         `json:"position"`
	Required      bool                            `json:"required"`
	Type          EventTypeCustomQuestionTypeEnum `json:"type"`
}
