package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mi11km/zikanwarikun-back/graph"
	"github.com/mi11km/zikanwarikun-back/graph/generated"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
)

// todo configファイルから読み込む
const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.Init()
	database.Migrate()
	defer func() {
		if err := database.Db.Close(); err != nil {
			log.Fatalf("action=close db, err=%s", err)
		}
	}()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
