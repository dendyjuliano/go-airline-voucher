package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	seatmap "voucher/index/internal/seetmap"
)

type GenerateRequest struct {
    Name         string `json:"name"`
    ID           string `json:"id"`
    FlightNumber string `json:"flightNumber"`
    Date         string `json:"date"`
    Aircraft     string `json:"aircraft"`
}

type GenerateResponse struct {
    Success bool     `json:"success"`
    Seats   []string `json:"seats,omitempty"`
    Error   string   `json:"error,omitempty"`
}

func GenerateHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req GenerateRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        // Check for existing assignment
        var count int
        err := db.QueryRow(
            "SELECT COUNT(*) FROM assignments WHERE flight_number = ? AND flight_date = ?",
            req.FlightNumber, req.Date,
        ).Scan(&count)
        if err != nil {
            http.Error(w, "DB error", http.StatusInternalServerError)
            return
        }
        if count > 0 {
            json.NewEncoder(w).Encode(GenerateResponse{Success: false, Error: "Assignment already exists"})
            return
        }
        // Generate seats
        seats, err := seatmap.GenerateSeats(req.Aircraft)
        if err != nil {
            json.NewEncoder(w).Encode(GenerateResponse{Success: false, Error: err.Error()})
            return
        }
        // Insert assignment
        _, err = db.Exec(
            "INSERT INTO assignments (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
            req.Name, req.ID, req.FlightNumber, req.Date, req.Aircraft, seats[0], seats[1], seats[2], time.Now().Format(time.RFC3339),
        )
        if err != nil {
            json.NewEncoder(w).Encode(GenerateResponse{Success: false, Error: "DB insert error"})
            return
        }
        json.NewEncoder(w).Encode(GenerateResponse{Success: true, Seats: seats})
    }
}