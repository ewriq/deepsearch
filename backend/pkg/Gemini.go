package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Gemini(query string) (string, error) {
	apiKey := Config.Gemini
	if apiKey == "" {
		return "", fmt.Errorf("Gemini API key not found in config")
	}
	crossdata, err := FetchCrossData(query)
	if err != nil {
		log.Printf("Cross data fetch error: %v", err)
		return "", fmt.Errorf("cross data fetch error: %v", err)
	}

	prompt := fmt.Sprintf("\n%s%s%s%s", Config.Prompt, query, crossdata, query)

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

func FetchCrossData(query string) ([]string, error) {
	var combinedResults []string

	payload := map[string]string{
		"term": query,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:5000/api/search/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("External API isteği hatası: %v", err)
	} else {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("External API response okuma hatası: %v", err)
		} else {
			combinedResults = append(combinedResults, string(bodyBytes))
			log.Printf("External API response: %s", string(bodyBytes))
		}
	}

	return combinedResults, nil
}
