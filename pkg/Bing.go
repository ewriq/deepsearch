package pkg

import (
	"deepsearch/utils"
	"encoding/json"

	g "github.com/serpapi/google-search-results-golang"
)

func Bing(query string) (string, error) {
	Config := utils.LoadConfig("./config/search.ini")

	parameter := map[string]string{
		"engine": "bing",
		"q":      query,  
		"cc":     "US",  
	}

	search := g.NewGoogleSearch(parameter, Config.Serpapi)

	rawResults, err := search.GetJSON()
	if err != nil {
		return "", err 
	}

	prettyJSON, err := json.MarshalIndent(rawResults, "", "  ")
	if err != nil {
		return "", err 
	}


	return  string(prettyJSON), nil
}
