package config

// Config settings for project
type Config struct {
	Enabled       bool
	Port          string
	StorageParams map[string]string
}

// NewCofig djgklvj
func NewCofig() *Config {
	sp := make(map[string]string)
	sp["db-name"] = "mock-db"
	return &Config{
		Enabled:       true,
		Port:          "8000",
		StorageParams: sp,
	}
}
