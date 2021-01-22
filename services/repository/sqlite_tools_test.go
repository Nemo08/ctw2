package repository

import (
	"testing"

	"github.com/Nemo08/ctw2/entity"
	"github.com/Nemo08/ctw2/infrastructure"
)

func TestMakeSqliteConnection(t *testing.T) {
	//Создаем новый логгер
	logr := infrastructure.NewLogrusLogger()

	//Создаем новое соединение к базе sqlite
	gdb, err := MakeSqliteConnection("file::memory:?cache=shared", logr)

	//Подключаем репозиторий по соединению
	crepo := NewContactSqliteRepository(gdb)

	//Автомигрируем стуктуру
	err = crepo.AutoMigrate()
	if err != nil {
		t.Error("Error of contact creation ", err.Error())
	}

	c := entity.NewContact("РУССКИЙ", "vasia@google.com", "788963")
	_, err = crepo.Create(c)
	if err != nil {
		t.Error("Error of contact creation ", err.Error())
	}

	cc, err := crepo.Search("русский")
	if err != nil {
		t.Error("Error of contact search ", err.Error())
	}
	if len(cc) == 0 {
		t.Error("Error of contact search - nothing returns")
		return
	}
	if cc[0].ID != c.ID {
		t.Error("Error of contact search - strange returns")
	}
}
