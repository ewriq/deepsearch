package database

import (

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB


func init() {

	var err error
	db, err = gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	err = db.Exec(`CREATE VIRTUAL TABLE IF NOT EXISTS search USING fts4(query, content);`).Error
	if err != nil {
		log.Fatalf("FTS tablosu oluşturulamadı: %v", err)
	}
}


type Search struct {
	RowID   int    `gorm:"column:rowid;primaryKey"`
	Query   string
	Content string
}

func (Search) TableName() string {
	return "search"
}
