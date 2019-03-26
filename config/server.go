package config

// Server cofiguration
type Server struct {
	config *Config
}

// New creates Server with config
func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}