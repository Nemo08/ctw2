package contact

import (
	"github.com/Nemo08/ctw2/entity"
	ifr "github.com/Nemo08/ctw2/infrastructure"
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

type ContactUsecase interface {
	Get(id tools.ID) (*entity.Contact, error)
	Search(query string) ([]*entity.Contact, error)
	List() ([]*entity.Contact, error)
	Create(name, email, phone string) (tools.ID, error)
	Update(e *entity.Contact) error
	Delete(id tools.ID) (tools.ID, error)
}

type contactUsecase struct {
	repo   Repository
	logger ifr.Logger
}

//NewContactUsecase for create new usecase
func NewContactUsecase(r Repository, l ifr.Logger) ContactUsecase {
	return contactUsecase{
		repo:   r,
		logger: l,
	}
}

func (cu contactUsecase) Get(id tools.ID) (*entity.Contact, error) {
	cu.logger.Info("get contact by id", id)
	return cu.repo.Get(id)
}

func (cu contactUsecase) Search(query string) ([]*entity.Contact, error) {
	cu.logger.Info("search contact", query)
	return cu.repo.Search(query)
}

func (cu contactUsecase) List() ([]*entity.Contact, error) {
	cu.logger.Info("list contact")
	return cu.repo.List()
}

func (cu contactUsecase) Create(name, email, phone string) (tools.ID, error) {
	cu.logger.Info("create contact", name, email, phone)
	cont := entity.NewContact(name, email, phone)
	return cu.repo.Create(cont)
}

func (cu contactUsecase) Update(e *entity.Contact) error {
	cu.logger.Info("update contact", e.ID)
	return cu.repo.Update(e)
}

func (cu contactUsecase) Delete(id tools.ID) (tools.ID, error) {
	cu.logger.Info("delete contact", id)
	return cu.repo.Delete(id)
}
