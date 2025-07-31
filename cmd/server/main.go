package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	
	"golang_json_html_convert/internal/server"
)

func main() {
	port := flag.Int("p", 8080, "Port number to listen on")
	tmplPath := flag.String("t", "templates/threat.html.tmpl", "Path to HTML template file")
	flag.Parse()

	srv, err := server.New(*tmplPath)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	handler := srv.Router

	addr := fmt.Sprintf(":%d", *port)
	httpSrv := &http.Server{		
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("golang_json_html_convert listening on http://localhost%s (template: %s)", addr, *tmplPath)
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
