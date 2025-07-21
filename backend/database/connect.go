package database

import (
	"log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	// SQLite bağlantısı
	db, err = gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	err = db.Exec(`CREATE VIRTUAL TABLE IF NOT EXISTS search USING fts4(query, content);`).Error
	if err != nil {
		log.Fatalf("FTS tablosu oluşturulamadı: %v", err)
	}
	log.Println("✅ GORM ile Sqlite bağlantısı kuruldu.")
}

func (Search) TableName() string {
	return "search"
}
