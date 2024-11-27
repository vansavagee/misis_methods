package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const TOKEN = ""

func main() {
	// сервер с одной ручкой
	// хендлер, отправялвет запрос в 4 API, обрабатывает и отправляет ответ беку

	http.HandleFunc("/get_sentiment", h())

	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}

func h() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf(" io.ReadAll(...): %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var q data
		if err := json.Unmarshal(body, &q); err != nil {
			http.Error(w, fmt.Sprintf(" json.Unmarshal(...): %v", err), http.StatusBadRequest)
			return
		}

		resp, err := doReq(q)
		if err != nil {
			http.Error(w, fmt.Sprintf("doreq(...): %v", err), http.StatusBadRequest)
			return
		}

		// Возвращаем успешный ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

type data struct {
	Id     string   `json:"id"`
	Lang   string   `json:"lang"`
	Text   string   `json:"text"`
	Labels []string `json:"labels"`
}

func doReq(d data) ([]byte, error) {
	switch id := d.Id; id {
	case "1":
		return doSentimentAnalysis9Req(d)
	case "2":
		return doSentimentByApiNinjasReq(d)
	case "3":
		return doComprehendItReq(d)
	case "4":
		return doSentimentApi3Req(d)
	}
	return nil, fmt.Errorf("Unknown ID.")
}
