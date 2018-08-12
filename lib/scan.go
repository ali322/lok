package lib

import (
	"fmt"
	"jav/model"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func Scan(c *cli.Context) error {
	if c.NArg() != 1 {
		log.Fatal("expect only one arg")
	}
	root := c.Args().First()
	path := c.String("path")
	remark := c.String("remark")
	filepath.Walk(filepath.Join(root, path), func(dest string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		} else {
			if !info.IsDir() {
				fileName := info.Name()
				ext := filepath.Ext(fileName)
				modTime := info.ModTime().Format("2006-01-02 15:04:05")
				if ext == ".avi" || ext == ".wmv" || ext == ".mkv" || ext == ".mp4" {
					fmt.Println(fileName)
					stuff := &model.Stuff{
						Name:      fileName,
						Ext:       ext,
						Path:      path,
						Remark:    remark,
						UpdatedAt: modTime,
					}
					db := model.Conn()
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
