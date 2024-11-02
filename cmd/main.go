package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DexScen/ApiLinkShortener/internal/repository/psql"
	"github.com/DexScen/ApiLinkShortener/internal/service"
	"github.com/DexScen/ApiLinkShortener/internal/transport/rest"
	"github.com/DexScen/ApiLinkShortener/pkg/database"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty123",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	linksRepo := psql.NewLinks(db)
	linksService := service.NewLinks(linksRepo)
	handler := rest.NewHandler(linksService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("Server started at:", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
