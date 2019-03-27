package config

import (
	"encoding/json"
	"net/http"
)

// Server cofiguration
type Server struct {
	config      *Config
	calcStorage StorageController
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/calculators", s.people)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	calcs := s.calcStorage.GetCalculators()
	bytes, _ := json.Marshal(calcs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, calcStorage *StorageController) *Server {
	return &Server{
		config:      config,
		calcStorage: calcStorage,
	}
}
