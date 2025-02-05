package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"sync"
)

// Env values
type Env struct {
	Server Server
	Log    Log
	Doc    Doc
}

// Server config
type Server struct {
	Host string
	Port string
}

// Log config
type Log struct {
	Enabled bool
	Level   string
}

// Doc - swagger information
type Doc struct {
	Title       string
	Description string
	Enabled     bool
	Version     string
}

var (
	env  *Env
	once sync.Once
)

// GetEnv returns env values
func GetEnv() *Env {

	once.Do(func() {

		viper.AutomaticEnv()
		godotenv.Load("./internal/config/.env")

		env = new(Env)
		env.Server.Port = viper.GetString("PORT")

		env.Log.Enabled = viper.GetBool("LOG_ENABLED")
		env.Log.Level = viper.GetString("LOG_LEVEL")
	})

	return env
}
