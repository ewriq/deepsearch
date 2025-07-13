package pkg

type GeminiRequest struct {
    Model string `json:"model"`
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
