package database

import (
	"fmt"
	"log"

	"deepsearch/pkg"
)

// Kayıt ekle
func AddSearchEntry(query, content string) error {
	err := db.Exec("INSERT INTO search (query, content) VALUES (?, ?)", query, content).Error
	if err != nil {
		log.Printf("Arama kaydı eklenirken hata: %v", err)
	}
	return err
}

// Arama yap
func PerformSearch(term string) ([]Search, error) {
	var results []Search
	sql := `
	SELECT rowid, query, content
	FROM search
	WHERE search MATCH ?
	LIMIT 10000000000;
	`
	err := db.Raw(sql, term).Scan(&results).Error
	return results, err
}

// Arama işlemi ve varsa Gemini ile özet alma
func RunSearch(term string) (string, error) {
	log.Printf("Veritabanında '%s' için arama yapılıyor...", term)
	results, err := PerformSearch(term)
	if err != nil {
		log.Printf("Veritabanı araması yapılırken hata: %v", err)
		return "", fmt.Errorf("veritabanı araması yapılırken hata: %w", err)
	}

	if len(results) > 0 {
		log.Println("Veritabanında sonuç bulundu:")
		firstContent := results[0].Content
		log.Printf("RowID: %d\n", results[0].RowID)
		return firstContent, nil
	}

	log.Printf("Veritabanında '%s' için sonuç bulunamadı. Web araması yapılıyor ve ekleniyor...", term)

	geminiSummary, geminiErr := pkg.Gemini(term)
	if geminiErr != nil {
		log.Printf("Gemini'den özet alınırken hata: %v", geminiErr)
		return "", fmt.Errorf("Gemini'den özet alınamadı: %w", geminiErr)
	}

	if geminiSummary != "" {
		log.Println("Gemini'den özet alındı. Veritabanına ekleniyor...")
		if err := AddSearchEntry(term, geminiSummary); err != nil {
			return "", err
		}

		log.Println("Kayıt eklendi. Tekrar veritabanında aranıyor...")

		resultsAfterAdd, reSearchErr := PerformSearch(term)
		if reSearchErr != nil {
			log.Printf("Kayıt eklendikten sonra tekrar arama yapılırken hata: %v", reSearchErr)
			return "", fmt.Errorf("kayıt eklendikten sonra tekrar arama yapılırken hata: %w", reSearchErr)
		}

		if len(resultsAfterAdd) > 0 {
			log.Println("Yeni eklenen kayıt bulundu ve gösteriliyor:")
			newlyAddedContent := resultsAfterAdd[0].Content
			log.Printf("RowID: %d\n\n", resultsAfterAdd[0].RowID)
			return newlyAddedContent, nil
		} else {
			log.Println("Kayıt eklenmesine rağmen hala sonuç bulunamadı. Bir sorun olabilir.")
			return "", fmt.Errorf("kayıt eklenmesine rağmen hala sonuç bulunamadı")
		}
	} else {
		log.Println("Gemini'den boş özet alındığı için veritabanına kayıt eklenmedi.")
		return "", fmt.Errorf("Gemini'den boş özet döndü, kayıt eklenmedi")
	}
}
