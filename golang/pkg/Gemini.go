package pkg

import (
	"bytes"
	"deepsearch/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Gemini(query string)  (string, error) {
	config := utils.LoadConfig("./config/search.ini")
	apiKey := config.Gemini
    if apiKey == "" {
        return "", fmt.Errorf("Gemini API key not found in config")
    }
	crossdata, err := FetchCrossData(query)
	if err != nil {
		log.Printf("Cross data fetch error: %v", err)
		return "", fmt.Errorf("cross data fetch error: %v", err)
	}

    prompt := fmt.Sprintf("\n%s%s%s%s", config.Prompt, query, crossdata, query)

    url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-lite:generateContent?key=" + apiKey

    payload, _ := json.Marshal(map[string]interface{}{
        "contents": []map[string]interface{}{
            {"parts": []map[string]string{{"text": prompt}}},
        },
    })

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Gemini API error: %s", string(body))
    }

    var gemResp GeminiResponse
    if err := json.Unmarshal(body, &gemResp); err != nil {
        return "", err
    }
    if len(gemResp.Candidates) == 0 || len(gemResp.Candidates[0].Content.Parts) == 0 {
        return "", fmt.Errorf("No summary returned from Gemini")
    }
    return gemResp.Candidates[0].Content.Parts[0].Text, nil
}



func FetchCrossData(query string) ([]string, error)  {
	var combinedResults []string
	var hasError bool

	googleJSON, err := Google(query)
	if err != nil {
		log.Printf("Google sonuçları çekilirken hata: %v", err)
		hasError = true
	} else {
		combinedResults = append(combinedResults, fmt.Sprintf("\n%s", googleJSON))
	}

	yandexJSON, err := Yandex(query)
	if err != nil {
		log.Printf("Yandex sonuçları çekilirken hata: %v", err)
		hasError = true
	} else {
		combinedResults = append(combinedResults, fmt.Sprintf("\n%s", yandexJSON))
	}

	bingJSON, err := Bing(query)
	if err != nil {
		log.Printf("Bing sonuçları çekilirken hata: %v", err)
		hasError = true
	} else {
		combinedResults = append(combinedResults, fmt.Sprintf("\n%s", bingJSON))
	}

	if hasError {
		return combinedResults, fmt.Errorf("bir veya daha fazla arama motoru hata döndürdü, detaylar için logları kontrol edin")
	}

	return combinedResults, nil
}