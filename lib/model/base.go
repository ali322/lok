package model

import (
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
)

var Conn *gorm.DB

func init() {
	var err error
	home, _ := homedir.Dir()
	Conn, err = gorm.Open("sqlite3", filepath.Join(home, "lok.db"))
	if err != nil {
		panic("failed to connect database")
	}
}
