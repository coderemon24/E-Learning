package config

import (
	"net/http"

	"github.com/coderemon24/e-learning/internal/routes"
)

// register all route file here from routes folder
func RegisterRoutes(mux *http.ServeMux) {
	routes.Api(mux)
}