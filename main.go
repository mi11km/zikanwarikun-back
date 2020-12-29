package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mi11km/zikanwarikun-back/config"
	"github.com/mi11km/zikanwarikun-back/graph"
	"github.com/mi11km/zikanwarikun-back/graph/generated"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/internal/middleware/auth"
	"github.com/rs/cors"
)

func main() {
	database.Init()
	database.Migrate()
	defer func() {
		db, err := database.Db.DB()
		if err != nil {
			log.Fatalf("action=failed to get *sql.DB, err=%s", err)
		}
		if err := db.Close(); err != nil {
			log.Fatalf("action=failed to close db, err=%s", err)
		}
	}()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Authorization Content-Type"},
		AllowCredentials: true,
		Debug:            config.Cfg.Server.Env == "development",
	}).Handler)
	router.Use(auth.Middleware())

	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := config.Cfg.Server.Port
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
