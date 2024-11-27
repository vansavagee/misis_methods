package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const TOKEN = "780ef0781dmsh79cdc66ed2ded1fp105f55jsn77c894d1bdac"

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

		var q []data
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

type APIResponse struct {
	Outputs Outputs `json:"outputs"`
}
type Outputs struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
}

func doReq(d []data) ([]byte, error) {
	var aPIResponses = make([]APIResponse, 0, len(d))
	for _, v := range d {
		resp, err := doComprehendItReq(v)
		fmt.Println(resp)
		if err != nil {
			time.Sleep(1 * time.Second)
			resp, err = doComprehendItReq(v)
			if err != nil {
				return nil, err
			}
		}
		aPIResponses = append(aPIResponses, *resp)
	}

	results := make([][]float64, len(aPIResponses))

	for i, v := range aPIResponses {
		t := make([]float64, 3)
		t[0] = float64(v.Outputs.Positive)
		t[1] = float64(v.Outputs.Negative)
		t[2] = float64(v.Outputs.Neutral)

		results[i] = append(results[i], t...)
	}

	return json.Marshal(results)
}
