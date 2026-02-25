package main

import (
	"net/http"

	"github.com/coderemon24/e-learning/internal/config"
)

func main() {
	mux := http.NewServeMux();
	
	config.RegisterRoutes(mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"))
	})
	
	http.ListenAndServe(":8080", mux)
	
}