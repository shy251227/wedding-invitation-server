package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/juhonamnam/wedding-invitation-server/env"
	"github.com/juhonamnam/wedding-invitation-server/httphandler"
	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "/tmp/sql.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	// ▼▼▼▼▼ 주소에서 "/api"를 삭제하여 웹사이트와 통일시켰습니다. ▼▼▼▼▼
	mux.Handle("/guestbook", new(httphandler.GuestbookHandler))
	mux.Handle("/attendance", new(httphandler.AttendanceHandler))
	// ▲▲▲▲▲ 여기가 마지막으로 변경된 부분입니다. ▲▲▲▲▲

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}