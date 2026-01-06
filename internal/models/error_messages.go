package models

type ErrorMessage struct {
	Code         string
	Component    string
	ResponseType string
	One          string
	Other        string
}

func (f *ErrorMessage) TableName() string {
	return "error_messages"
}
