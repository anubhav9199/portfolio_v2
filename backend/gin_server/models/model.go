package models


type Services struct {
	Image		string	`json:"image,omitempty" validate:"required"`
	Title		string	`json:"title,omitempty" validate:"required"`
	Description	string	`json:"description,omitempty" validate:"required"`
}