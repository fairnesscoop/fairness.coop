package web

import (
	"fmt"
	"net/http"

	"fairness.coop/victor/internal/domain"
	"fairness.coop/victor/web/handlers"
)

type Server struct {
	handler http.Handler
}

func NewServer(ctn domain.Container) *Server {
	r := handlers.NewRouter(ctn)
	return &Server{handler: r}
}

func (s *Server) Listen() error {
	fmt.Println("Listening on http://localhost:1314")
	return http.ListenAndServe(":1314", s.handler)
}
