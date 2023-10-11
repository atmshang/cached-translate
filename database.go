package translate

import (
	"fmt"
	"github.com/atmshang/plog"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
	"time"
)

const dsName = "translate.db"

// db represents the singleton instance of the database connection.
var (
	db   *gorm.DB
	once sync.Once
)

// connectDB establishes a new connection to the database and runs the necessary migrations.
func connectDB() (*gorm.DB, error) {
	// Establish a new database connection.
	_db, err := gorm.Open(sqlite.Open(dsName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	// Run migrations for the Cache model.
	autoMigrate(_db, &Cache{})
	return _db, nil
}

// getDBInstance returns the singleton instance of the database connection,
// initializing it if it has not been initialized yet.
func getDBInstance() *gorm.DB {
	var err error
	once.Do(func() {
		db, err = connectDB()
		// If an error occurred while connecting to the database, log a message and stop the program.
		if err != nil {
			plog.Panic(err)
		}
		// If the database connection is nil, log a message and stop the program.
		if db == nil {
			plog.Panic("the database connection is nil!")
		}
	})
	return db
}

// autoMigrate runs the necessary migrations for the provided model.
func autoMigrate(db *gorm.DB, dst interface{}) {
	err := db.AutoMigrate(dst)
	// If an error occurred while migrating, log a message and stop the program.
	if err != nil {
		plog.Panic("failed to auto migrate ImageInfo")
	}
}

// The init function is called when the package is initialized.
func init() {
	_ = getDBInstance()
}

// Cache represents a translation cache in the database.
type Cache struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Text       string
	From       string
	To         string
	Translated string
}
