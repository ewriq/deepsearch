package pkg

import (
	"deepsearch/utils"
	"fmt"

	g "github.com/serpapi/google-search-results-golang"
)

func Bing(query string) ([]string, error) {
	Config := utils.LoadConfig("./config/search.ini")

	parameter := map[string]string{
		"engine": "bing",
		"q":      query,
		"cc":       "US",
		"hl":       "en",
	}

	search := g.NewGoogleSearch(parameter, Config.Serpapi)

	results, err := search.GetJSON()
	if err != nil {
		return nil, err
	}

	organic, ok := results["organic_results"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected format in Bing response")
	}

	var snippets []string
	for _, item := range organic {
		if entry, ok := item.(map[string]interface{}); ok {
			if snippet, ok := entry["snippet"].(string); ok {
				snippets = append(snippets, snippet)
			}
		}
	}

	return snippets, nil
}
