package translate

import (
	"github.com/atmshang/plog"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

const dsName = "translate.db"

func getDBInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsName), &gorm.Config{})
	if err != nil {
		plog.Panic("GetDBInstance", err)
	}
	return db
}

func autoMigrate(db *gorm.DB, dst interface{}) {
	err := db.AutoMigrate(dst)
	if err != nil {
		plog.Panic("failed to auto migrate ImageInfo")
	}
}

func init() {
	db := getDBInstance()
	autoMigrate(db, &Cache{})
}

// Cache 翻译缓存
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
