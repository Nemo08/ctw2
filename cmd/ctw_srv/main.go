// main.go
package main

import (
	_ "github.com/Nemo08/ctw2/entity"
	"github.com/Nemo08/ctw2/infrastructure"
	"github.com/Nemo08/ctw2/services/http"
	"github.com/Nemo08/ctw2/services/repository"

	_ "github.com/davecgh/go-spew/spew"
)

func main() {
	//Создаем новый логгер
	logr := infrastructure.NewLogrusLogger()

	//Создаем новое соединение к базе sqlite
	gdb, err := repository.MakeSqliteConnection("sqlite.db", logr)

	//Подключаем репозиторий по соединению
	crepo := repository.NewContactSqliteRepository(gdb)

	//Автомигрируем стуктуру
	err = crepo.AutoMigrate()
	if err != nil {
		logr.Error(err.Error())
	}
	/*
		c := entity.NewContact("Вася Ж", "vasia@google.com", "788963")
		uid, err := crepo.Create(c)
		if err != nil {
			logr.Error(err.Error())
		}
		fmt.Println(uid, c)
		c.Phone = "445566"
		c.Name = "вася Жп"
		crepo.Update(c)
		crepo.Delete(uid)
		cc, _ := crepo.Search("жП")
		spew.Dump(cc)
	*/
	srv := http.NewHttpService("api/v1/")

	srv.Start(7123)
}
