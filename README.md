# Airline Voucher Seat Assignment API

This project provides a simple REST API for airline crew to assign random voucher seats for flights, ensuring no duplicate assignments per flight/date and supporting different aircraft seat layouts.

## Features

- Input flight and crew details
- Generate 3 unique random seats based on aircraft type
- Prevent duplicate assignments for the same flight/date
- Persist assignments in a SQLite database

## Endpoints

### 1. Check Assignment

**POST** `/api/check`

**Request Body:**

```json
{
  "flightNumber": "GA102",
  "date": "2025-07-12"
}
```

**Response:**

```json
{
  "exists": true
}
```

---

### 2. Generate Assignment

**POST** `/api/generate`

**Request Body:**

```json
{
  "name": "Sarah",
  "id": "98123",
  "flightNumber": "ID102",
  "date": "2025-07-12",
  "aircraft": "Airbus 320"
}
```

**Response:**

```json
{
  "success": true,
  "seats": ["3B", "7C", "14D"]
}
```

---

## Supported Aircraft Types

| Aircraft Type  | Rows | Seats per Row    |
| -------------- | ---- | ---------------- |
| ATR            | 1-18 | A, C, D, F       |
| Airbus 320     | 1-32 | A, B, C, D, E, F |
| Boeing 737 Max | 1-32 | A, B, C, D, E, F |

---

## Getting Started

### 1. Clone the repository

```sh
git clone <your-repo-url>
cd go-airline-voucher
```

### 2. Install dependencies

```sh
go mod tidy
```

### 3. Run the database migration

```sh
sqlite3 voucher.db < internal/db/migration.sql
```

> If you don't have `sqlite3` CLI, install it with `brew install sqlite`.

### 4. Start the server

```sh
go run ./cmd/main.go
```

The server will run on `http://localhost:8080`.

---

## Testing the API

You can use [Postman](https://www.postman.com/) or `curl` to test the endpoints.

**Example:**

```sh
curl -X POST http://localhost:8080/api/check \
  -H "Content-Type: application/json" \
  -d '{"flightNumber":"GA102","date":"2025-07-12"}'
```

---

## Project Structure

```
cmd/
  main.go
internal/
  api/
    check.go
    generate.go
  db/
    db.go
    migration.sql
  seatmap/
    layout.go
go.mod
README.md
```
