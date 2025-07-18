package pkg

import (
	"deepsearch/utils"
	"fmt"

	g "github.com/serpapi/google-search-results-golang"
)


func Yandex(query string) ([]string, error) {
	Config := utils.LoadConfig("./config/search.ini")

	params := SearchParams{
		Engine: "yandex",
		Query:  query,
	}

	parameter := map[string]string{
		"engine": params.Engine,
		"text":   params.Query, 
	}

	search := g.NewGoogleSearch(parameter, Config.Serpapi)
	results, err := search.GetJSON()
	if err != nil {
		return nil, err
	}

	organicResults, ok := results["organic_results"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected format for organic_results")
	}

	var snippets []string
	for _, item := range organicResults {
		if entry, ok := item.(map[string]interface{}); ok {
			var result OrganicResult
			if snippet, ok := entry["snippet"].(string); ok {
				result.Snippet = snippet
				snippets = append(snippets, result.Snippet)
			}
		}
	}

	return snippets, nil
}
