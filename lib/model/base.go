package model

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
)

var Conn *gorm.DB

func init() {
	home, _ := homedir.Dir()
	path := filepath.Join(home, "lok.db")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}

	var err error
	Conn, err = gorm.Open("sqlite3", path)
	if err != nil {
		panic("failed to connect database")
	}
}
