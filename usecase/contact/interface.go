package contact

import (
	"github.com/Nemo08/ctw2/entity"
	"github.com/Nemo08/ctw2/tools"
)

type Repository interface {
	Get(id tools.ID) (*entity.Contact, error)
	Search(query string) ([]*entity.Contact, error)
	List() ([]*entity.Contact, error)

	Create(e *entity.Contact) (tools.ID, error)
	Update(e *entity.Contact) error
	Delete(id tools.ID) (tools.ID, error)
}

type Usecase interface {
	Repository
	Create(name, email, phone string) (tools.ID, error)
}
