package database


type SearchResult struct {
	ID             uint 
	Content        string 
	RelevanceScore float64
}


type Search struct {
	RowID          int  `gorm:"column:rowid;primaryKey"` 	// satır kimliği
	Query          string // sorgu metni
	Content        string // içerik
	LikeLevel      int     // benzerlik seviyesi (0-100)
	RelevanceScore float64 // benzerlik skoru
	Highlight      string  // vurgulama metni
}
