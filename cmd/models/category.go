package models

type Category struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required,min=2,max=25"`
	Description string `json:"description,omitempty"`
}
