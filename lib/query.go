package lib

import (
	"fmt"
	"lok/lib/model"
	"strings"

	"github.com/urfave/cli"
)

// Query query stuff by condition
func Query(c *cli.Context) error {
	key := c.Args().First()
	remark := c.String("remark")
	path := c.String("path")
	db := model.Conn
	defer db.Close()
	stuffs := []model.Stuff{}
	if key != "" {
		db = db.Where("name LIKE ?", strings.Join([]string{"%", key, "%"}, ""))
	}
	if remark != "" {
		db = db.Where("remark = ?", remark)
	}
	if path != "" {
		db = db.Where("path = ?", path)
	}
	db.Find(&stuffs)
	fmt.Printf("共找到 %d 个结果\n", len(stuffs))
	for _, stuff := range stuffs {
		fmt.Printf(" %s, %s, %s\n", stuff.Remark, stuff.Path, stuff.Name)
	}
	return nil
}
