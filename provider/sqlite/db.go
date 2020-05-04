package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is the DB that will performs all operation
var DB *gorm.DB

func init() {
	var err error
	if DB, err = NewDB("./provider/sqlite/db.sqlite"); err != nil {
		panic(err)
	}
}

// NewDB returns a new DB connection
func NewDB(path string) (*gorm.DB, error) {
	// connect to the example sqlite, create if it doesn't exist
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
