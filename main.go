package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 200; i++ {
			statusCode := http.StatusOK
			if r.Header.Get("Authorization") == "" {
				statusCode = http.StatusUnauthorized
			}

			response := Response{
				Status:  statusCode,
				Message: http.StatusText(statusCode),
				Data:    nil,
			}

			w.Header().Set("Content-Type", "application/json")

			jsonData, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(statusCode)
			_, _ = w.Write(jsonData)
		}
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
