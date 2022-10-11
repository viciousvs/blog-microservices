package update

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DTO struct {
	UUID    string `json:"uuid,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func (a DTO) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Title, validation.Length(3, 0)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Content, validation.Length(2, 0)),
		validation.Field(&a.UUID, validation.Required, validation.Length(36, 0)),
	)
}

func (a DTO) EmptyBody() bool {
	if a.Title == "" && a.Content == "" {
		return true
	}
	return false
}
