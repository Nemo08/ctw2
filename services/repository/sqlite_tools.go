package repository

import (
	"database/sql"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mattn/go-sqlite3"

	"github.com/Nemo08/ctw2/infrastructure"
)

func utflower(s string) string {
	return strings.ToLower(s)
}

func MakeSqliteConnection(URI string, l infrastructure.Logger) (*gorm.DB, error) {
	sql.Register("sqlite3_custom", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			if err := conn.RegisterFunc("utflower", utflower, true); err != nil {
				return err
			}

			return nil
		},
	})
	db, err := gorm.Open("sqlite3_custom", URI)
	db.LogMode(true)
	db.SetLogger(l)
	return db, err
}
