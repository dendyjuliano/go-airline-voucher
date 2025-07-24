package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type CheckRequest struct {
    FlightNumber string `json:"flightNumber"`
    Date         string `json:"date"`
}

type CheckResponse struct {
    Exists bool `json:"exists"`
}

func CheckHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req CheckRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        var count int
        err := db.QueryRow(
            "SELECT COUNT(*) FROM assignments WHERE flight_number = ? AND flight_date = ?",
            req.FlightNumber, req.Date,
        ).Scan(&count)
        if err != nil {
            http.Error(w, "DB error", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(CheckResponse{Exists: count > 0})
    }
}