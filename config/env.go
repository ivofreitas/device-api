package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"sync"
)

// Env values
type Env struct {
	Server   Server
	Log      Log
	Doc      Doc
	Database Database
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

// Database - Postgres configuration
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
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

		env.Database.Host = viper.GetString("DB_HOST")
		env.Database.Port = viper.GetString("DB_PORT")
		env.Database.User = viper.GetString("DB_USER")
		env.Database.Password = viper.GetString("DB_PASSWORD")
		env.Database.DBName = viper.GetString("DB_NAME")
		env.Database.SSLMode = viper.GetString("DB_SSLMODE")
	})

	return env
}
