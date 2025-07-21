package database

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"deepsearch/pkg"
)

// Kayıt ekle
func AddSearchEntry(query, content string) error {
	err := db.Exec("INSERT INTO search (query, content) VALUES (?, ?)", query, content).Error
	log.Println("sdsfdsodoksdookkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk k8sssssssssssssssssssssssssssssssssssssssssssssssss")
	if err != nil {
		log.Printf("Arama kaydı eklenirken hata: %v", err)
	}
	return err
}

// Arama yap
func PerformSearch(query string) ([]Search, error) {
	var allResults []Search
	term := strings.ToLower(query)
	likeTerm := "%" + term + "%"

	sql := `
    SELECT rowid, query, content
    FROM search
    WHERE LOWER(query) LIKE ? OR LOWER(content) LIKE ?
    LIMIT 50;
    `

	err := db.Raw(sql, likeTerm, likeTerm).Scan(&allResults).Error
	if err != nil {
		return nil, err
	}

	queryWords := strings.Fields(term)
	finalResults := make([]Search, 0, len(allResults))

	for _, res := range allResults {
		totalText := strings.ToLower(res.Query + " " + res.Content)
		wordsInText := strings.Fields(totalText)
		wordCount := len(wordsInText)
		if wordCount == 0 {
			continue
		}
		
		matchCount := 0
		for _, qw := range queryWords {
			if strings.Contains(totalText, qw) {
				matchCount++
			}
		}

		if matchCount == 0 {
			continue
		}


		totalMatchCount := 0
		for _, qw := range queryWords {
			totalMatchCount += strings.Count(totalText, qw)
		}

		exactQueryMatch := 0
		if strings.Contains(totalText, term) {
			exactQueryMatch = 10
		}

		// Basit alaka puanı hesapla
		relevance := (float64(totalMatchCount) / float64(wordCount)) * float64(matchCount) + float64(exactQueryMatch)
		likeLevel := int(relevance * 100)
		if likeLevel > 100 {
			likeLevel = 100
		}

		highlight := ""
		if totalMatchCount >= 10 {
			highlight = fmt.Sprintf("⚠ Kelime '%s' toplam %d kez geçti", query, totalMatchCount)
		}

		res.LikeLevel = likeLevel
		res.RelevanceScore = relevance
		res.Highlight = highlight

		finalResults = append(finalResults, res)
	}

	sort.SliceStable(finalResults, func(i, j int) bool {
		return finalResults[i].RelevanceScore > finalResults[j].RelevanceScore
	})

	return finalResults, nil
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
