package lib

import (
	"fmt"
	"jav/lib/model"
	"log"
	"strings"

	"github.com/urfave/cli"
)

// Query query stuff by condition
func Query(c *cli.Context) error {
	if c.NArg() != 1 {
		log.Fatal("only one argument expected")
	}
	key := c.Args().First()
	remark := c.String("remark")
	db := model.Conn
	defer db.Close()
	stuffs := []model.Stuff{}
	db = db.Where("name LIKE ?", strings.Join([]string{"%", key, "%"}, ""))
	if remark != "" {
		db = db.Where("remark = ?", remark)
	}
	db.Find(&stuffs)
	fmt.Println(len(stuffs))
	for _, stuff := range stuffs {
		fmt.Println(stuff.Name)
	}
	return nil
}
