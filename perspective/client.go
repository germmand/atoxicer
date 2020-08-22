package perspective

import (
    "log"
    "fmt"
    "encoding/json"
    "bytes"
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
    return session;
}

func (ps *PerspectiveSession) ObtainToxicity(comment string) (map[string]interface{}, error) {
    data := map[string]interface{}{
        "comment": map[string]string{
            "text": comment,
        },
        "languages": []string{"es"},
        "requestedAttributes": map[string]interface{}{
            "TOXICITY": map[string]interface{}{},
        },
    }

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

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    log.Println("Toxicity obtained successfully.")

    return result, nil
}
