package server

import (
	"net/http"

	"golang_json_html_convert/internal/handler"
	"golang_json_html_convert/internal/render"
)

// Server wires routes and dependencies (acts as a small Facade).
// Implements Facade + Factory method
type Server struct {
	Router *http.ServeMux
}

// New loads the template and builds a ServeMux with routes.
func New(templatePath string) (*Server, error) {
	tmpl, err := render.Load(templatePath)
	if err != nil {
		return nil, err
	}
	h := handler.New(tmpl)

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/render", h.Render)

	return &Server{Router: mux}, nil
}
