package handler

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"

	"golang_json_html_convert/internal/model"
)

type jsonDecoder struct{}

// BodyDecoder is a Strategy interface for decoding request bodies to a target struct.
type BodyDecoder interface {
	Accepts(r *http.Request) bool
	Decode(r *http.Request, dst any) error
}

// Handler groups dependencies for HTTP endpoints.
type Handler struct {
	Tmpl         *template.Template
	MaxBodyBytes int64
	Decoders     []BodyDecoder
}

type formDecoder struct{}

func (jsonDecoder) Accepts(r *http.Request) bool {
	ct := r.Header.Get("Content-Type")
	return strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/plain")
}
func (jsonDecoder) Decode(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(dst)
}
func (formDecoder) Accepts(r *http.Request) bool {
	ct := r.Header.Get("Content-Type")
	return strings.HasPrefix(ct, "application/x-www-form-urlencoded") || strings.HasPrefix(ct, "multipart/form-data") || ct == ""
}
func (formDecoder) Decode(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	raw := r.FormValue("json_input")
	if strings.TrimSpace(raw) == "" {
		buf, _ := io.ReadAll(r.Body)
		raw = string(buf)
		if strings.TrimSpace(raw) == "" {
			return io.EOF
		}
	}
	return json.Unmarshal([]byte(raw), dst)
}

// New returns a Handler with default decoders and limits.
func New(tmpl *template.Template) *Handler {
	return &Handler{
		Tmpl:         tmpl,
		MaxBodyBytes: 1 << 20, // 1 MiB
		Decoders:     []BodyDecoder{jsonDecoder{}, formDecoder{}},
	}
}

// Index serves the HTML form for manual input.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, _ = io.WriteString(w, `<!DOCTYPE html>
<html>
  <head><meta charset="UTF-8"><title>Threat data</title></head>
  <body>
    <h1>Threat record</h1>
    <form action="/render" method="POST">
      <label for="json_input">JSON input:</label><br/>
      <textarea rows="12" cols="80" id="json_input" name="json_input"></textarea><br/>
      <input type="submit" value="Submit"/>
    </form>
  </body>
</html>`)
}

// Render decodes the request into a model.Threat and executes the template.
func (h *Handler) Render(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit body size for robustness
	r.Body = http.MaxBytesReader(w, r.Body, h.MaxBodyBytes)
	defer r.Body.Close()

	var threat model.Threat
	var decErr error
	for _, d := range h.Decoders {
		if d.Accepts(r) {
			decErr = d.Decode(r, &threat)
			break
		}
	}
	if decErr != nil {
		http.Error(w, "invalid JSON: "+decErr.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if err := h.Tmpl.Execute(w, threat); err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
