package connection

import (
	"fmt"
	"gorm.io/gorm"
	"news/internal/config"
	dbConf "news/internal/config/db"
)

var NewsDB *gorm.DB

func init() {
	var err error
	cfg := config.GetConfig()
	fmt.Println("runnn")

	NewsDB, err = NewsConnection(cfg.Database.News.Mysql)
	if err != nil {
		fmt.Println("Error in connection to database: ", err)
	}

}

func NewsConnection(db dbConf.Database) (*gorm.DB, error) {
	driver, err := NewInstanceDb(db)
	if err != nil {
		return nil, err
	}
	return driver.Db().(*gorm.DB), nil
}
