package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	dbConf "news/internal/config/db"
	"news/pkg/adapter/db/model"
	"time"
)

// DriverMySQL object Driver MySQL
type DriverMySQL struct {
	config dbConf.Database
	db     *gorm.DB
}

// NewMySQLDriver new object SQL Driver
func NewMySQLDriver(config dbConf.Database) (DbDriver, error) {
	dbConn, err := connect(config)

	if err != nil {
		panic("failed to connect database")
		//return nil, err
	}
	err = MigrateSchema(dbConn)
	fmt.Println("Errorr", err)
	return &DriverMySQL{
		config: config,
		db:     dbConn,
	}, nil
}

func connect(config dbConf.Database) (*gorm.DB, error) {
	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Dbname
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	var dbConn *gorm.DB
	var err error
	currentWaitTime := 2
	trialCount := 0

	for dbConn == nil && trialCount < 5 {
		trialCount++
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			fmt.Println("unable connecting to DB.")
			if trialCount == 5 {
				return nil, err
			}
			fmt.Println("retrying in", currentWaitTime, "seconds...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
			dbConn = nil
		}
	}
	conn, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(7)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(1 * time.Hour)

	return dbConn, err
}

func (m *DriverMySQL) Db() interface{} {
	return m.db
}

var tables = []interface{}{
	&model.KategoriModel{},
	&model.UserModel{},
	&model.PhotoModel{},
	&model.LoginHistoryModel{},
}

func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
