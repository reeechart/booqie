package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/reeechart/booql/book/graphql"
	"github.com/reeechart/booql/book/handlers"
)

type server struct {
	host   string
	port   int
	router chi.Router
}

func NewServer(host string, port int) *server {
	return &server{
		host:   host,
		port:   port,
		router: chi.NewRouter(),
	}
}

func (s *server) setupHandler() {
	graphqlSchema := graphql.GetSchema()

	pingHandler := handlers.NewPingHandler()
	graphqlHandler := handlers.NewGraphQLHandler(graphqlSchema)
	s.router.Route("/ping", func(r chi.Router) {
		r.Get("/", pingHandler.Ping)
	})

	s.router.Route("/graphql", func(r chi.Router) {
		r.Get("/", graphqlHandler.GraphQL)
		r.Post("/", graphqlHandler.GraphQL)
	})
}

func (s *server) Run() {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.setupHandler()
	log.Printf("Server is listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
