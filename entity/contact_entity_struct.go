package entity

import (
	"github.com/Nemo08/ctw2/tools"
)

//Структура контакта в entity
type Contact struct {
	ID tools.ID

	Name  string
	Email string
	Phone string
}

func NewContact(name, email, phone string) *Contact {
	return &Contact{
		ID:    tools.NewID(),
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
