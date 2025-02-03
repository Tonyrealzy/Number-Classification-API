package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d/math?json", n)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err!= nil {
        return "", err
    }
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err!= nil {
        return "", err
    }
	if fact, ok := result["text"].(string); ok {
		return fact, nil
	}
	return "", fmt.Errorf("could not find a math fact for number %d", n)
}
