package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Config struct {
		HTTPServer
		GTapi
		Logger
	}
	Logger struct {
		Level      string `env:"APP_MODE" envDefault:"dev"`
		Debugging  string `env:"DEBUG" envDefault:"true"`
		OutputFile string `env:"LOG_FILE_PATH" envDefault:"app.log"`
	}
	HTTPServer struct {
		Port string `env:"HTTP_SERVER_PORT" envDefault:"8080"`
		Host string `env:"HTTP_SERVER_HOST" envDefault:"localhost"`
	}
	GTapi struct {
		URL          string `env:"GT_URL" envDefault:"https://groupietrackers.herokuapp.com/api"`
		ArtistsURL   string `env:"GT_ARTISTS_URL" envDefault:"https://groupietrackers.herokuapp.com/api/artists"`
		LocationsURL string `env:"GT_LOCATIONS_URL" envDefault:"https://groupietrackers.herokuapp.com/api/locations"`
		DatesURL     string `env:"GT_DATES_URL" envDefault:"https://groupietrackers.herokuapp.com/api/dates"`
		RelationURL  string `env:"GT_RELATION_URL" envDefault:"https://groupietrackers.herokuapp.com/api/relation"`
	}
)

func New() (*Config, error) {
	var cfg Config

	// Load environment variables from .env file
	err := LoadEnvVars(".env")
	if err != nil {
		return nil, fmt.Errorf("cannot read env: %w", err)
	}
	// TODO: check do default variables settled up in case zero value returned. If no then manage it
	// Access HTTP Server environment variable values
	if value := os.Getenv("HTTP_SERVER_HOST"); value != "" {
		cfg.HTTPServer.Host = value
		log.Printf("HTTP_SERVER_HOST is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot find server host: %w", err)
	}
	if value := os.Getenv("HTTP_SERVER_PORT"); value != "" {
		cfg.HTTPServer.Port = value
		log.Printf("HTTP_SERVER_PORT is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot find server port: %w", err)
	}

	// Access API Groupie Trackers environment variable values
	if value := os.Getenv("GT_URL"); value != "" {
		cfg.GTapi.URL = value
		log.Printf("GT_URL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}
	if value := os.Getenv("GT_ARTISTS_URL"); value != "" {
		cfg.GTapi.ArtistsURL = value
		log.Printf("GT_ARTISTS_URL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}
	if value := os.Getenv("GT_LOCATIONS_URL"); value != "" {
		cfg.GTapi.LocationsURL = value
		log.Printf("GT_LOCATIONS_URL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}
	if value := os.Getenv("GT_DATES_URL"); value != "" {
		cfg.GTapi.DatesURL = value
		log.Printf("GT_DATES_URL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}
	if value := os.Getenv("GT_RELATION_URL"); value != "" {
		cfg.GTapi.RelationURL = value
		log.Printf("GT_RELATION_URL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	// Access logger environment variable values
	if value := os.Getenv("LEVEL"); value != "" {
		log.Printf("LEVEL is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}
	if value := os.Getenv("LOG_FILE_PATH "); value != "" {
		log.Printf("LOG_FILE_PATH  is: %s", value)
	} else {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	return &cfg, nil

}

// LoadEnvVars loads environment variables from a .env file
func LoadEnvVars(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			pair := strings.SplitN(line, "=", 2)
			if len(pair) == 2 {
				os.Setenv(pair[0], pair[1])
			}
		}
	}

	return scanner.Err()
}
