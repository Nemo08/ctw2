package contact

import (
	"github.com/Nemo08/ctw2/entity"
)

type Repository interface {
	Get(id entity.ID) (*entity.Contact, error)
	Search(query string) ([]*entity.Contact, error)
	List() ([]*entity.Contact, error)

	Create(e *entity.Contact) (entity.ID, error)
	Update(e *entity.Contact) error
	Delete(id entity.ID) (entity.ID, error)
}

type Usecase interface {
	Repository
	Create(name, email, phone string) (entity.ID, error)
}
