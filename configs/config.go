package configs

import (
	"database/sql"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers"
)

type Config struct {
	DB       *sql.DB
	Auth     *AuthConfig
	Handlers *Handlers
	Env      string
}

type Handlers struct {
	AuthHandler   *handlers.AuthHandler
	EventHandler  *handlers.EventHandler
	VendorHandler *handlers.VendorHandler
}

type AuthConfig struct {
	TTL       int
	SecretKey string
}
