package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"tkame123-net/worldcup-gq-server/graph"
	"tkame123-net/worldcup-gq-server/graph/generated"
	"tkame123-net/worldcup-gq-server/infra/mongodb"
	"tkame123-net/worldcup-gq-server/infra/mongodb/competition"
	"tkame123-net/worldcup-gq-server/infra/mongodb/match"
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

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	cl := mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	repoC := competition.NewRepository(cl)
	repoP := player.NewRepository(cl)
	repoM := match.NewRepository(cl)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MongoCompetition: repoC,
		MongoPlayer:      repoP,
		MongoMatch:       repoM,
	}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
