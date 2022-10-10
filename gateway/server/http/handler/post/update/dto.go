package update

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DTO struct {
	UUID    string `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a DTO) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Title, validation.Length(3, 0)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Content, validation.Length(10, 0)),
		validation.Field(&a.UUID, validation.Required, validation.Length(36, 0)),
	)
}
