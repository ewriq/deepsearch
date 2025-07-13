package database

import (
	"fmt"
	"log"
	
	"deepsearch/pkg" 
)

func AddSearchEntry(query, content string) {
	err := db.Create(&Search{Query: query, Content: content}).Error
	if err != nil {
		log.Printf("Arama kaydı eklenirken hata: %v", err)
	}
}

func PerformSearch(term string) ([]SearchResult, error) {
	var results []SearchResult
	sql := `
		SELECT id, content,
		ts_rank(to_tsvector('turkish', content), plainto_tsquery('turkish', ?)) AS relevance_score
		FROM search
		WHERE to_tsvector('turkish', content) @@ plainto_tsquery('turkish', ?)
		ORDER BY relevance_score DESC LIMIT 10;
	`
	err := db.Raw(sql, term, term).Scan(&results).Error
	return results, err
}

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
		log.Printf("ID: %d | Skor: %.2f\n", results[0].ID, results[0].RelevanceScore)
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
		AddSearchEntry(term, geminiSummary)
		log.Println("Kayıt eklendi. Tekrar veritabanında aranıyor...")

		resultsAfterAdd, reSearchErr := PerformSearch(term)
		if reSearchErr != nil {
			log.Printf("Kayıt eklendikten sonra tekrar arama yapılırken hata: %v", reSearchErr)
			return "", fmt.Errorf("kayıt eklendikten sonra tekrar arama yapılırken hata: %w", reSearchErr) 
		}

		if len(resultsAfterAdd) > 0 {
			log.Println("Yeni eklenen kayıt bulundu ve gösteriliyor:")
			newlyAddedContent := resultsAfterAdd[0].Content
			log.Printf("ID: %d | Skor: %.2f\n\n", resultsAfterAdd[0].ID, resultsAfterAdd[0].RelevanceScore)
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