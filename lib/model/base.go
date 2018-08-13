package model

import (
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Conn *gorm.DB

func init() {
	var err error
	Conn, err = gorm.Open("sqlite3", filepath.Join("data", "lok.db"))
	if err != nil {
		panic("failed to connect database")
	}
}
