package pkg

import (
	"deepsearch/utils"
	"fmt"

	g "github.com/serpapi/google-search-results-golang"
)



func Google(query string) ([]string, error) {
	Config := utils.LoadConfig("./config/search.ini")

	params := SearchParams{
		Engine: "google",
		Query:  query,
	}

	parameter := map[string]string{
		"engine": params.Engine,
		"q":      params.Query,
	}

	search := g.NewGoogleSearch(parameter, Config.Serpapi)
	results, err := search.GetJSON()
	if err != nil {
		return nil, err
	}

	organicResults, ok := results["organic_results"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected format")
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
