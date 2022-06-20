package connection

import (
	"errors"
	"log"
	dbConf "news/internal/config/db"
	"news/pkg/shared/enum"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	errorInvalidDbInstance = errors.New("Invalid db instance")
)

var atomicinz uint64
var instanceDb map[string]DbDriver = make(map[string]DbDriver)

var once sync.Once

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
	// Dml(command string, params ...string) error
	// Query(command string, params ...string) (*sql.Rows, error)
}

// NewInstanceDb is used to create a new instance DB
func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Dbname

	// once.Do(func() {
	switch config.Adapter {
	case enum.MySql:
		dbConn, sqlErr := NewMySQLDriver(config)
		if sqlErr != nil {
			err = sqlErr
			log.Fatal("Database connection failed.")
		}
		instanceDb[dbName] = dbConn
	default:
		err = errorInvalidDbInstance
	}
	// })

	return instanceDb[dbName], err
}
