package routes

import (
	"net/http"

	"github.com/coderemon24/e-learning/internal/controllers"
)

func Api(app *http.ServeMux) {
	
	app.HandleFunc("/api/v1/test",controllers.Test)
	
}