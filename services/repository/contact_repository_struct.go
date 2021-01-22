package repository

import (
	"time"

	"github.com/Nemo08/ctw2/entity"
	"github.com/Nemo08/ctw2/tools"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

//Contact Структура контакта в базе
type Contact struct {
	ID        tools.ID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name  string
	Email string
	Phone string
}

// E2R Копирует данные из структуры базы в структуру entity
func E2R(ec entity.Contact) (*Contact, error) {
	c := Contact{}
	err := copier.Copy(&c, &ec)
	return &c, err
}

// R2E Копирует данные из структуры entity в структуру базы
func R2E(c Contact) (*entity.Contact, error) {
	ec := entity.Contact{}
	err := copier.Copy(&ec, &c)
	return &ec, err
}

// ES2RS Копирует массив данных из структуры базы в структуру entity
func ES2RS(ec []entity.Contact) ([]Contact, error) {
	c := []Contact{}
	err := copier.Copy(&c, &ec)
	return c, err
}

// RS2ES Копирует массив данных из структуры entity в структуру базы
func RS2ES(c []Contact) ([]*entity.Contact, error) {
	ec := []*entity.Contact{}
	err := copier.Copy(&ec, &c)
	return ec, err
}
