package lib

import (
	"fmt"
	"log"
	"lok/lib/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Scan scan stuffs in disk and add correspond record to database
func Scan(c *cli.Context) error {
	if c.NArg() != 1 {
		log.Fatal("only one argument expected")
	}
	root := c.Args().First()
	path := c.String("path")
	remark := c.String("remark")
	exts := strings.Split(c.String("ext"), ",")
	filepath.Walk(filepath.Join(root, path), func(dest string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		} else {
			if !info.IsDir() {
				fileName := info.Name()
				ext := filepath.Ext(fileName)
				modTime := info.ModTime().Format("2006-01-02 15:04:05")
				if contains(exts, ext) {
					fmt.Println(fileName)
					stuff := &model.Stuff{
						Name:      fileName,
						Ext:       ext,
						Path:      path,
						Remark:    remark,
						UpdatedAt: modTime,
					}
					db := model.Conn
					db.AutoMigrate(&model.Stuff{})
					db.Create(stuff)
					defer db.Close()
				}
			}
		}
		return nil
	})

	fmt.Println("finished scan file in", root)
	return nil
}
