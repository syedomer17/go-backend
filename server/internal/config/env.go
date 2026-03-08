package config 

import (
	"os"
	"fmt"
	"strings"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	PORT int 
	JWTSECRET string
	UPSTASH_REDIS_REST_URL string
	UPSTASH_REDIS_REST_TOKEN string
}

func Load() (*Config, error){
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	port, err := strconv.Atoi(strings.TrimSpace(os.Getenv("PORT")))
	if err != nil {
		return nil, fmt.Errorf("Invalid PORT value: %v", err)
	}

	cfg := &Config{
		MongoURI: strings.TrimSpace(os.Getenv("MONGO_URI")),
		PORT: port,
		JWTSECRET: strings.TrimSpace(os.Getenv("JWT_SECRET")),
		UPSTASH_REDIS_REST_URL: strings.TrimSpace(os.Getenv("UPSTASH_REDIS_REST_URL")),
		UPSTASH_REDIS_REST_TOKEN: strings.TrimSpace(os.Getenv("UPSTASH_REDIS_REST_TOKEN")),
	}

	if cfg.MongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is required")
	}
	
	if cfg.JWTSECRET == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	if cfg.UPSTASH_REDIS_REST_URL == "" {
		return nil, fmt.Errorf("UPSTASH_REDIS_REST_URL is required")
	}

	if cfg.UPSTASH_REDIS_REST_TOKEN == "" {
		return nil, fmt.Errorf("UPSTASH_REDIS_REST_TOKEN is required")
	}

	return cfg, nil
}