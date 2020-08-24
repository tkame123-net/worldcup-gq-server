package main

import (
	"log"
	"net/http"
	"os"
	"tkame123-net/worldcup-gq-server/graph"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/infra/mongodb"
	"tkame123-net/worldcup-gq-server/infra/mongodb/competition"
	"tkame123-net/worldcup-gq-server/infra/mongodb/player"
	"tkame123-net/worldcup-gq-server/lib/env"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	//// env
	if err := env.Load(); err != nil {
		log.Fatalf("faild to load env: %+v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repoC := competition.NewRepository(cl)
	repoP := player.NewRepository(cl)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MongoCompetition: repoC,
		MongoPlayer:      repoP,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
