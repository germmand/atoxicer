package perspective

type AnalyzeAttributes struct {
	Toxicity map[string]interface{} `json:"TOXICITY"`
}

type Comment struct {
	Text string `json:"text"`
}

type AnalyzeRequest struct {
	Comment             *Comment           `json:"comment"`
	RequestedAttributes *AnalyzeAttributes `json:"requestedAttributes"`
	Languages           []string           `json:"languages"`
}

func CreateAnalyzeRequest(comment string, languages []string) *AnalyzeRequest {
	analyzeRequest := &AnalyzeRequest{
		Comment: &Comment{
			Text: comment,
		},
		Languages: languages,
		RequestedAttributes: &AnalyzeAttributes{
			Toxicity: make(map[string]interface{}),
		},
	}
	return analyzeRequest
}
