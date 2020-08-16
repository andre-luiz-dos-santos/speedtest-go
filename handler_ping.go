package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var reqJSON map[string]interface{}

	// Read JSON request.
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096))
	err = dec.Decode(&reqJSON)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Add client address to JSON.
	reqJSON["addr"] = r.RemoteAddr

	// Send JSON request to ping API.
	b, err := json.Marshal(reqJSON)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// Hide how long it takes to call the ping API.
	go func() {
		resp, err := http.Post(pingURL, "application/json", bytes.NewReader(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "ping API error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			fmt.Fprintf(os.Stderr, "ping API error: %v\n", resp.Status)
			return
		}
	}()
}
