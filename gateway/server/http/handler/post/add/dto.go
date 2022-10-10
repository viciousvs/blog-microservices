package add

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a DTO) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Title, validation.Required, validation.Length(3, 0)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Content, validation.Required, validation.Length(10, 0)),
	)
}
