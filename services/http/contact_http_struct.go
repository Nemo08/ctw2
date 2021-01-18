package http

import (
	"github.com/Nemo08/ctw2/tools"
)

type Contact struct {
	ID    tools.ID `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email" validate:"required, email"`
	Phone string   `json:"phone"`
}
