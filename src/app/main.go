package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	serverTime := time.Now()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Bikin Cors GET Method Only
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		currentTime := time.Now()
		upTime := time.Since(serverTime)

		// TODO: Bikin response buat endpoint /health
		response := map[string]interface{}{
			"nama":      "Gilbran Mahdavikia Raja",
			"nrp":       "5025241134",
			"status":    "UP",
			"timestamp": currentTime,
			"uptime":    upTime.String(),
			"test": "success",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
