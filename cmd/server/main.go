package main

import (
	"context"
	"log"
	"net/http"

	"github.com/lorsanstand/Aether-go/internal/config"
	"github.com/lorsanstand/Aether-go/internal/database"
	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/lorsanstand/Aether-go/internal/handlers/middlewares"
	"github.com/lorsanstand/Aether-go/internal/handlers/users"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed scrap ENV: %v", err)
	}

	db, err := database.NewPostgresPGX(ctx, cfg)
	if err != nil {
		log.Fatalf("Database failed connection: %v", err)
	}
	defer db.Close()

	//conn, err := database.NewPostgresDB(cfg)
	//if err != nil {
	//	log.Fatalf("Database failed connection: %v", err)
	//}
	//
	//migrt := migrator.MustGetNewMigrator(Aether_go.MigrationsFS, "migrations")
	//
	//if err := migrt.ApplyMigrations(conn); err != nil {
	//	log.Fatalf("Failed migrations: %v", err)
	//}
	//conn.Close()

	queries := gen.New(db)

	userHandler := users.NewUserHandler(queries)

	mux := http.NewServeMux()
	mux.Handle("/users/", userHandler.RegisterRoutes())

	var handler http.Handler = mux
	handler = middlewares.LogMiddleware(middlewares.GetUserIdMiddleware(handler, cfg.SecretKey))

	log.Println("Start server")
	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatalln(err)
	}
}
