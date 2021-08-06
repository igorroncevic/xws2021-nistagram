package common

import "github.com/rs/cors"

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins, for now
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
