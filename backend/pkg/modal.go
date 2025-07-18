package pkg

import "deepsearch/utils"

type GeminiRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

var Config utils.ConfigList

func init() {
    Config = utils.LoadConfig("./config/search.ini")
}

type SearchParams struct {
	Engine string
	Query  string
	CC     string
	HL     string
}

type OrganicResult struct {
	Snippet string
}