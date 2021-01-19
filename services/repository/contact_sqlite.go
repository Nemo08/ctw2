package repository

import (
	"github.com/Nemo08/ctw2/entity"
	"github.com/Nemo08/ctw2/tools"

	_ "fmt"

	"github.com/google/uuid"
	_ "gorm.io/driver/sqlite"

	//"gorm.io/gorm"
	"strings"

	"github.com/jinzhu/gorm"
)

type ContactSqliteRepository struct {
	gormdb *gorm.DB
}

func NewContactSqliteRepository(db *gorm.DB) *ContactSqliteRepository {
	return &ContactSqliteRepository{
		gormdb: db,
	}
}

func (r *ContactSqliteRepository) Get(id tools.ID) (*entity.Contact, error) {
	c := Contact{ID: id}
	var ec *entity.Contact
	tx := r.gormdb.First(&c)
	if tx.Error != nil {
		return ec, tx.Error
	}

	ec, err := R2E(c)
	if err != nil {
		return ec, err
	}
	return ec, nil
}

func (r *ContactSqliteRepository) List() ([]entity.Contact, error) {
	cs := []Contact{}
	var ec []entity.Contact
	tx := r.gormdb.Find(&cs)
	if tx.Error != nil {
		return ec, tx.Error
	}

	ec, err := RS2ES(cs)
	if err != nil {
		return ec, err
	}
	return ec, nil
}

func (r *ContactSqliteRepository) Create(e *entity.Contact) (tools.ID, error) {
	c, err := E2R(*e)
	nullid := uuid.New()
	if err != nil {
		return nullid, err
	}
	tx := r.gormdb.Create(c)
	return e.ID, tx.Error
}

func (r *ContactSqliteRepository) Update(e *entity.Contact) error {
	getc := Contact{ID: e.ID}

	//Ищем полученный контакт в базе по id
	tx := r.gormdb.First(&getc)
	if tx.Error != nil {
		return tx.Error
	}

	//Копируем из полученного все в контакт из базы
	getc.Name = e.Name
	getc.Phone = e.Phone
	getc.Email = e.Email

	// и сохраняем
	tx = r.gormdb.Save(getc)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *ContactSqliteRepository) Delete(id tools.ID) (tools.ID, error) {
	getc := Contact{ID: id}

	//Ищем полученный контакт в базе по id
	tx := r.gormdb.Delete(&getc)
	if tx.Error != nil {
		return id, tx.Error
	}

	return id, nil
}

func (r *ContactSqliteRepository) Search(query string) ([]*entity.Contact, error) {
	var ec []*entity.Contact

	queryL := "%" + strings.ToLower(query) + "%"
	tx := r.gormdb.Where("(utflower(name) LIKE ?) OR (utflower(phone) LIKE ?) OR (utflower(email) LIKE ?)", queryL, queryL, queryL).Find(&ec)

	if tx.Error != nil {
		return ec, tx.Error
	}

	return ec, tx.Error
}

func (r *ContactSqliteRepository) AutoMigrate() error {
	r.gormdb.AutoMigrate(Contact{})
	return nil
}
