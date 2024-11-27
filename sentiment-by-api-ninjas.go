package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func doSentimentByApiNinjasReq(d data) ([]byte, error) {
	baseURL := "https://sentiment-by-api-ninjas.p.rapidapi.com/v1/sentiment"

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL: %v\n", err)
	}

	query := u.Query()
	query.Set("text", d.Text)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-host", "sentiment-by-api-ninjas.p.rapidapi.com")
	req.Header.Set("x-rapidapi-key", TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
