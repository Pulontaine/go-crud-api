package main

import (
	"api/internal/db"
	"api/internal/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err = db.Migrate(database)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	handler.RegisterRoutes(router, database)

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
