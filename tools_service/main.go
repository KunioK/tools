package main

import (
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Ready"))
}

// TODO: 在此补充业务接口
func toolsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func main() {
	port := getEnv("PORT", "8080")
	log.Printf("Starting tools service on port %s", port)

	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/tools", toolsHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
