package main

import (
	"log"
	"net/http"
	"voucher/index/internal/api"
	"voucher/index/internal/db"
)

func main() {
    database := db.ConnectDB("voucher.db")
    defer database.Close()

    http.HandleFunc("/api/check", api.CheckHandler(database))
    http.HandleFunc("/api/generate", api.GenerateHandler(database))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}