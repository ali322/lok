package lib

import (
	"fmt"
	"log"

	"lok/lib/model"

	"github.com/urfave/cli"
)

// List stuffs in disk
func List(c *cli.Context) error {
	if c.NArg() > 1 {
		log.Fatal("wrong argument")
	}
	db := model.Conn
	defer db.Close()
	if c.NArg() == 0 {
		rows, err := db.Table("stuffs").Select("remark").Group("remark").Rows()
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var remark string
			if err := rows.Scan(&remark); err != nil {
				log.Fatal(err)
			}
			fmt.Println(remark)
		}
	} else {
		remark := c.Args().First()
		rows, err := db.Table("stuffs").Select("path").Where("remark = ?", remark).Group("path").Rows()
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var path string
			if err := rows.Scan(&path); err != nil {
				log.Fatal(err)
			}
			fmt.Println(path)
		}
	}
	return nil
}
