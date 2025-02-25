package main

import (
	"net/http"
	"time"

	"github.com/wushengyouya/blog-service/internal/routers"
)

func main() {
	engine := routers.NewRouters()
	// engine.Run()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
