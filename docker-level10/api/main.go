package main

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.Error("failed to connect to DB", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			users = append(users, u)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	slog.Info("API server running on :8080")
	http.ListenAndServe(":8080", nil)
}
