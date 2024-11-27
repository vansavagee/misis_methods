package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type reqComprehendIt struct {
	Text   string   `json:"text"`
	Labels []string `json:"labels"`
}

func doComprehendItReq(d data) ([]byte, error) {
	url := "https://comprehend-it.p.rapidapi.com/predictions/ml-zero-nli-model"

	reqData := reqComprehendIt{
		Text:   d.Text,
		Labels: d.Labels,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal reqData: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-host", "comprehend-it.p.rapidapi.com")
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
