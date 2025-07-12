package pkg

import (
	"deepsearch/utils"
	"encoding/json" 

	g "github.com/serpapi/google-search-results-golang"
)

func Yandex(query string) (string, error) {
	Config := utils.LoadConfig("./config/search.ini")
	parameter := map[string]string{
		"engine": "yandex",
		"text":   query,
	}

	search := g.NewGoogleSearch(parameter, Config.Serpapi)
	results, err := search.GetJSON()
	if err != nil {
		return "", err
	}

	organicResults := results["organic_results"] 

	prettyJSON, err := json.MarshalIndent(organicResults, "", "  ")
	if err != nil {
		return "", err
	}

	return string(prettyJSON), nil
}