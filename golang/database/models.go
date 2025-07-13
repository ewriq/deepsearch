package database

import "gorm.io/gorm"

type Search struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Query   string `gorm:"uniqueIndex"`
	Content string `gorm:"type:text"`
}

type SearchResult struct {
	ID             uint
	Content        string
	RelevanceScore float64
}
func (Search) TableName() string {
	return "search"
}