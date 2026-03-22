package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/lmittmann/tint"
	"github.com/lorsanstand/Aether-go/internal/config"
	"github.com/lorsanstand/Aether-go/internal/database"
	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/lorsanstand/Aether-go/internal/handlers/middlewares"
	"github.com/lorsanstand/Aether-go/internal/handlers/users"
	slogctx "github.com/veqryn/slog-context"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed scrap ENV: %v", err)
	}

	logBaseHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      cfg.GetLogLevel(),
		TimeFormat: "15:04:05",
		AddSource:  true,
	})

	logCtxHandler := slogctx.NewHandler(logBaseHandler, nil)

	slog.SetDefault(slog.New(logCtxHandler))

	db, err := database.NewPostgresPGX(ctx, cfg)
	if err != nil {
		slog.Error("Database failed connection", "error", err)
		return
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

	slog.Info("Server started http://localhost:8000")
	err = http.ListenAndServe("0.0.0.0:8000", handler)
	if err != nil {
		slog.Error("Server crashed", "error", err)
	}
}
