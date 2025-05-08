package api

import "os"

type Config struct {
    Port string
}

func BuildConfigFromEnv() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return Config{
        Port: port,
    }
} 