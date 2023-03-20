package main

import (
	"bytes"
	"net/http"
)

func makeRequest(requestType, url, token string, payload []byte) ([]byte, error) {
	client := &http.Client{}

	var request *http.Request
	if payload != nil {
		requestBody := bytes.NewReader(payload)

	}
}
