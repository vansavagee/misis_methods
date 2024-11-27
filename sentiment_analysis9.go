package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type reqSentimentAnalysis9 struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

func doSentimentAnalysis9Req(d data) ([]byte, error) {
	url := "https://sentiment-analysis9.p.rapidapi.com/sentiment"
	reqData := make([]reqSentimentAnalysis9, 1)
	reqData[0].ID = "1"
	reqData[0].Language = d.Lang
	reqData[0].Text = d.Text

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal reqData: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-host", "sentiment-analysis9.p.rapidapi.com")
	req.Header.Set("x-rapidapi-key", TOKEN)
	req.Header.Set("Accept", "application/json")

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
