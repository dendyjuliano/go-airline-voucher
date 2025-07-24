package seatmap

import (
	"fmt"
	"math/rand"
	"time"
)

var layouts = map[string]struct {
    Rows  []int
    Seats []string
}{
    "ATR":         {Rows: rangeInt(1, 18), Seats: []string{"A", "C", "D", "F"}},
    "Airbus 320":  {Rows: rangeInt(1, 32), Seats: []string{"A", "B", "C", "D", "E", "F"}},
    "Boeing 737 Max": {Rows: rangeInt(1, 32), Seats: []string{"A", "B", "C", "D", "E", "F"}},
}

func rangeInt(start, end int) []int {
    arr := make([]int, end-start+1)
    for i := range arr {
        arr[i] = start + i
    }
    return arr
}

func GenerateSeats(aircraft string) ([]string, error) {
    layout, ok := layouts[aircraft]
    if !ok {
        return nil, fmt.Errorf("unknown aircraft type")
    }
    rand.Seed(time.Now().UnixNano())
    seatSet := make(map[string]struct{})
    for len(seatSet) < 3 {
        row := layout.Rows[rand.Intn(len(layout.Rows))]
        seat := layout.Seats[rand.Intn(len(layout.Seats))]
        seatNum := fmt.Sprintf("%d%s", row, seat)
        seatSet[seatNum] = struct{}{}
    }
    seats := make([]string, 0, 3)
    for s := range seatSet {
        seats = append(seats, s)
    }
    return seats, nil
}