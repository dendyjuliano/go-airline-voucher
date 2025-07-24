package main

import (
	"log"
	"net/http"
	"voucher/index/internal/api"
	"voucher/index/internal/db"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	database := db.ConnectDB("voucher.db")
	defer database.Close()

	http.Handle("/api/check", withCORS(http.HandlerFunc(api.CheckHandler(database))))
	http.Handle("/api/generate", withCORS(http.HandlerFunc(api.GenerateHandler(database))))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}