package perspective

type SummaryScore struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type Toxicity struct {
	SummaryScore *SummaryScore `json:"summaryScore"`
}

type AttributeScores struct {
	Toxicity *Toxicity `json:"TOXICITY"`
}

type AnalyzeResponse struct {
	AttributeScores   *AttributeScores `json:"attributeScores"`
	Languages         []string         `json:"languages"`
	DetectedLanguages []string         `json:"detectedLanguages"`
}
