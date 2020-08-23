package perspective

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PerspectiveSession struct {
	Key string
	Url string
}

func New(Key string) *PerspectiveSession {
	session := &PerspectiveSession{
		Key: Key,
		Url: fmt.Sprintf("https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=%s", Key),
	}
	return session
}

func (ps *PerspectiveSession) ObtainToxicity(comment string) (*AnalyzeResponse, error) {
	data := CreateAnalyzeRequest(comment, []string{"es"})

	bytesRepresentation, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	resp, err := http.Post(ps.Url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var analyzeResponse *AnalyzeResponse
	json.Unmarshal(body, &analyzeResponse)

	return analyzeResponse, nil
}
